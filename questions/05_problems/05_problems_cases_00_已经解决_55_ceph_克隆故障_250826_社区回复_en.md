# q
Ceph OSD出现“clone_size.count(clone)”断言失败崩溃的典型根因是什么？
# a
RBD快照已被删除，但PG内仍残留该快照的克隆对象。恢复回填（recover_backfill）遍历对象调用`add_object_context_to_pg_stat`时，通过`SnapSet::get_clone_bytes`查询clone_size，但因快照已删除导致`clone_size`映射中无对应条目，触发`ceph_assert(clone_size.count(clone))`失败。典型情况：snap id 63的快照已删除，但对象`rbd_data.118b66cf015830.0000000000001418:3f`仍存在。

# q
如何从日志中定位导致崩溃的具体残留对象？
# a
查看崩溃线程的最后日志，关注`add_object_context_to_pg_stat`处理的对象。在崩溃信息前几行可找到类似记录：
```
-548> get_object_context: … rbd_data.118b66cf015830.0000000000001418:3f … snapset: 3f=[]:{}
-544> add_object_context_to_pg_stat 4:48e35a64:::rbd_data.118b66cf015830.0000000000001418:3f
-120> … FAILED ceph_assert(clone_size.count(clone))
```
该对象`rbd_data.118b66cf015830.0000000000001418:3f`即为问题对象，其snapset中无clone_size信息。进一步可运行`rbd snap ls`确认snap id 63的快照已不存在。

# q
解决此类残留快照对象导致OSD崩溃的标准流程是什么？
# a
1. 备份相关RBD数据。
2. 使用`ceph-objectstore-tool`删除残留对象：
   ```bash
   # 列出对象的所有片段（含snapid）
   ceph-objectstore-tool --data-path /var/lib/ceph/osd/ceph-<id> --pgid <pgid> rbd_data.xxx --op list
   # 删除残留的snapid对象（例如:3f）
   ceph-objectstore-tool --data-path /var/lib/ceph/osd/ceph-<id> --pgid <pgid> rbd_data.xxx --op remove
   ```
3. 重启OSD，恢复正常后PG会自动触发backfill完成数据恢复。

