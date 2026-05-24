# q
Ceph OSD 在 recover_backfill 阶段崩溃，触发 `SnapSet::get_clone_bytes` 中 `clone_size.count(clone)` 断言失败的典型根因是什么？
# a
当一个 RBD 快照已经被删除，但该快照对应的数据分片对象（clone object）仍然残留在 OSD 上时，`recover_backfill` 会遍历所有对象并调用 `add_object_context_to_pg_stat`。在 `SnapSet::get_clone_bytes(clone)` 中检查 `clone_size` 是否包含该 clone，如果缺失就会触发断言失败。  
本例中残留对象为 `rbd_data.118b66cf015830.0000000000001418:3f`，通过 `snap ls` 确认其对应的 snap id 63 的快照卷已被删除。

# q
如何从 Ceph OSD 日志中定位此类 `ceph_assert` 失败的原因？
# a
将 `debug_osd` 调整为 30 或更高，复现故障后查看崩溃线程日志。在崩溃前会看到 `add_object_context_to_pg_stat` 的调用记录，例如：
```sh
-560> ... get_object_context: obc NOT found in cache: 4:48e35a64:::rbd_data.118b66cf015830.0000000000001418:3f
-548> ... get_object_context: 0x55dc1b33fb00 ... oi: ... ssc: 0x55dc05bb98c0 snapset: 3f=[]:{}
-544> ... add_object_context_to_pg_stat 4:48e35a64:::rbd_data.118b66cf015830.0000000000001418:3f
-120> ... FAILED ceph_assert(clone_size.count(clone))
```
其中 `snapset: 3f=[]:{}` 表示该对象是 clone，但对应的快照信息已缺失。Crash 信息中的 backtrace 也会显示 `PrimaryLogPG::recover_backfill` -> `add_object_context_to_pg_stat` -> `SnapSet::get_clone_bytes` 的调用链。

# q
对于因残留 clone 对象导致 OSD 崩溃的标准处理流程是什么？
# a
1. **确认快照已删除**：`rbd snap ls <pool>/<image>` 确认快照确实不存在。  
2. **使用 `ceph-objectstore-tool` 删除残留对象**：
   - 列出对象的所有分片：  
     ```sh
     ceph-objectstore-tool --data-path /var/lib/ceph/osd/ceph-<id> --type bluestore <pgid> list-objects | grep rbd_data.118b66cf015830.0000000000001418
     ```
   - 删除问题分片：  
     ```sh
     ceph-objectstore-tool --data-path /var/lib/ceph/osd/ceph-<id> --type bluestore <pgid> remove-object rbd_data.118b66cf015830.0000000000001418:3f
     ```
3. **恢复 OSD**：设置 `nobackfill` 为 unset 并启动该 OSD，观察 recovery 不再崩溃。  
⚠️ 操作前建议对该 RBD 数据做好备份，防止误删。

