# q
如何从Ceph集群的输出中快速定位存在unfound对象的PG？
# a
使用命令 `ceph pg dump | grep unfound` 可以过滤出状态中包含 `unfound` 的PG。例如：
```
4.9f  ... active+recovery_unfound+undersized+degraded+remapped ...
4.126 ... active+recovery_unfound+undersized+degraded+remapped ...
```
重点关注状态字段中的 `recovery_unfound` 标记，`pg dump` 会列出PG ID及其当前状态。

# q
解决Ceph PG unfound对象的标准操作流程是什么？
# a
1. 确认受损PG的ID：`ceph pg dump | grep unfound`
2. 根据需要暂时停止恢复以避免数据抖动：
   `ceph osd set nobackfill`
   `ceph osd set norecover`
3. 将unfound对象标记为丢失：
   `ceph pg <pgid> mark_unfound_lost revert`（尝试回滚到旧版本）
   或 `ceph pg <pgid> mark_unfound_lost delete`（彻底删除）
4. 完成标记后解除恢复限制：
   `ceph osd unset nobackfill`
   `ceph osd unset norecover`
5. 监控集群恢复：`watch -n 1 "ceph -s"`

# q
当 `mark_unfound_lost revert` 执行后PG仍未恢复，如何使用 `ceph-objectstore-tool` 手动提取或删除残留对象？
# a
- 先列出指定对象的所有版本，获取精确的JSON标识：
  ```bash
  ceph-objectstore-tool --data-path /var/lib/ceph/osd/ceph-<osd-id> --pgid <pgid> <object-name> --op list
  ```
  输出会提示多个版本，需要指定具体snapid。
- 根据输出构造JSON描述（如 `["4.112",{"oid":"rbd_data.118b66cf015830.0000000000001418","key":"","snapid":63,"hash":643483410,"max":0,"pool":4,"namespace":"","max":0}]`），然后用 `get-bytes` 提取数据备份，或用 `remove` 删除：
  ```bash
  ceph-objectstore-tool --data-path /var/lib/ceph/osd/ceph-0 --pgid 4.112 \
    '["4.112",{"oid":"rbd_data.118b66cf015830.0000000000001418","key":"","snapid":63,...}]' remove
  ```
- 注意：操作前建议先用 `get-bytes` 导出数据，避免丢失。

# q
在排查Ceph unfound问题时，哪些OSD级别的动态调试命令可用？
# a
可以动态调整OSD日志等级，减少日志干扰：
```bash
ceph tell osd.0 injectargs '--debug_osd 0'
```
将 `osd.0` 替换为目标OSD编号，`0` 表示关闭额外调试输出。

