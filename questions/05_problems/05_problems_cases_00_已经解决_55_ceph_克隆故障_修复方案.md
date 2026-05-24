# q
如何从OSD日志中定位Ceph PG克隆故障涉及的具体数据对象？
# a
在故障OSD的系统日志中搜索PG ID，例如：
```bash
cat ceph-osd.16.log | grep 4.112
```
日志中会出现类似 `lb 4:48e33b35:::rbd_data.00528c12bdeadf.000000000000149c:head` 的记录，从中即可确定故障数据分片对象名（如 `rbd_data.00528c12bdeadf.000000000000149c`）。

# q
使用`ceph-objectstore-tool`导出OSD中指定对象数据的标准命令是什么？
# a
```bash
ceph-objectstore-tool --data-path /var/lib/ceph/osd/ceph-<OSD_ID> \
  --op get-bytes --oid '<完整对象名>' --file <输出文件>
```
示例（导出rbd数据对象）：
```bash
ceph-objectstore-tool --data-path /var/lib/ceph/osd/ceph-3 \
  --op get-bytes \
  --oid 'rbd_data.7b70996d74511d.0000000000000bb5:16_head' \
  --file /tmp/backup_object.data
```

# q
当PG主副本数据尚未同步到新加入的OSD时，如何快速恢复PG？
# a
确认数据缺失的OSD（如osd.0）上没有对应PG的数据后，直接将其out出集群，让PG重新选择活动副本：
```bash
ceph osd out 0
ceph osd unset nobackfill
ceph osd unset norecover
```
之后PG会利用主副本（如osd.16）上已有的完整数据进行恢复。若需停止自动恢复操作，可使用：
```bash
ceph osd set nobackfill
ceph osd set norecover
```

