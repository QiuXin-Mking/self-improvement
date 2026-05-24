# q
Ceph PG在执行repair后仍处于inconsistent+failed_repair状态，典型根因是什么？
# a
根因通常是PG内RBD对象的snap克隆数据在部分OSD上缺失或损坏，导致repair无法修复。例如日志中报错：
```
repair 4.d 4:b6c47a39:::rbd_data.7b70996d74511d.0000000000000bb5:head : expected clone 4:b6c47a39:::rbd_data.7b70996d74511d.0000000000000bb5:16 1 missing
```
这表明对象`rbd_data.7b70996d74511d.0000000000000bb5`的快照`16`缺失，使PG 4.d一直处于inconsistent状态。

# q
如何从日志定位Ceph PG repair失败的具体缺失对象？
# a
查看集群日志或使用`ceph pg <pgid> query`，搜索包含`repair ... expected clone ... missing`的错误条目。例如：
```
2025-06-09T17:40:00.693+0800 7eff6cfb6700 -1 log_channel(cluster) log [ERR] : repair 4.d 4:b6c47a39:::rbd_data.7b70996d74511d.0000000000000bb5:head : expected clone 4:b6c47a39:::rbd_data.7b70996d74511d.0000000000000bb5:16 1 missing
```
该日志直接指出缺失的clone对象名，即可锁定问题对象。

# q
解决Ceph PG因快照对象缺失导致inconsistent的标准流程是什么？
# a
1. 定位问题对象和PG：通过`ceph pg <pgid> query`或日志找到缺失对象名及所属PG。
2. 确认对象所在的OSD：使用`ceph osd map <pool> <object>`查看acting OSD（注意head与snap可能在不同PG，需分别查询，如`rbd_data.7b70996d74511d.0000000000000bb5`在PG 4.d，其snap `:16`在PG 4.6）。
3. 停止目标OSD（必须OSD DOWN）：`systemctl stop ceph-osd@<osd_id>`。
4. 使用`ceph-objectstore-tool`直接操作OSD数据目录，删除不一致的对象及其snap：
   ```bash
   ceph-objectstore-tool --data-path /var/lib/ceph/osd/ceph-<osd_id> --op remove-all --oid 'rbd_data.7b70996d74511d.0000000000000bb5'
   ```
   案例中`removeall`会同时移除head和snap对象，例如输出：
   ```
   remove clone #4:b6c47a39:::rbd_data.7b70996d74511d.0000000000000bb5:16#
   remove #4:b6c47a39:::rbd_data.7b70996d74511d.0000000000000bb5:head#
   ```
5. 启动OSD：`systemctl start ceph-osd@<osd_id>`。
6. 执行deep-scrub：`ceph pg <pgid> deep-scrub`，验证PG恢复为`active+clean`。
（注意：该操作会丢失对象数据，仅在确认对象可重建或已备份时使用。）

# q
当`rados`或`ceph-objectstore-tool get-bytes`提示“No object id ... found”时，可能的原因是什么？
# a
常见原因是指定的对象名不正确，例如尝试获取snap对象`rbd_data.7b70996d74511d.0000000000000bb5:16_head`，而实际OSD上该snap对象可能缺少`_head`后缀或根本不存在。案例中执行：
```
ceph-objectstore-tool --data-path /var/lib/ceph/osd/ceph-3 'rbd_data.7b70996d74511d.0000000000000bb5:16_head' get-bytes /tmp/backup_object.data
```
报错：
```
No object id 'rbd_data.7b70996d74511d.0000000000000bb5:16_head' found or invalid JSON specified
```
此时应通过`rados -p <pool> ls | grep <object_prefix>`确认实际存在的对象名（如仅有`rbd_data.7b70996d74511d.0000000000000bb5`），再操作正确的对象。

