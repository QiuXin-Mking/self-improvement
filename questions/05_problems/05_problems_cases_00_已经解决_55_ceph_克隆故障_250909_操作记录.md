# q
如何从Ceph OSD日志中定位与克隆相关的故障对象？
# a
查看OSD日志，搜索 `add_object_context_to_pg_stat` 事件。该事件会输出涉及的对象，如：
```
zcat ceph-osd.23.log-20250907.gz | grep "add_object_context_to_pg_stat"
-2> 2025-09-06T22:37:42.103+0800 7fa3b6962700 10 osd.23 pg_epoch: 67350 pg[4.120(...)] add_object_context_to_pg_stat 4:04d9491d:::rbd_data.18570ff66275d5.000000000000a4da:d25
```
从日志中可提取出对象名（rbd_data.18570ff66275d5.000000000000a4da）和快照ID（d25，16进制为3365）。

# q
如何找出包含特定rbd_data前缀的RBD镜像？
# a
通过遍历存储池中的镜像并检查 `rbd info` 输出实现：
```bash
arr=($(rbd ls -p vms))
for img in "${arr[@]}"; do
    if rbd info vms/"$img" | grep -q 18570ff66275d5; then
        echo "$img"
    fi
done
```
本例中找到镜像 `6c5c11e7-108b-4a96-8977-74023816824e_disk`。之后可进一步检查快照：
```bash
rbd snap ls vms/6c5c11e7-108b-4a96-8977-74023816824e_disk
```

# q
解决Ceph OSD因克隆对象导致PG卡住的标准操作流程是什么？
# a
1. 停止故障OSD：
   ```bash
   systemctl stop ceph-osd@23
   ```
2. 列出PG中对象的快照：
   ```bash
   ceph-objectstore-tool --data-path /var/lib/ceph/osd/ceph-23 --pgid 4.120 rbd_data.18570ff66275d5.000000000000a4da --op list
   ```
   输出中会出现 `snapid:3365` 和 `snapid:-2`（head对象）。
3. （可选）导出快照数据：
   ```bash
   ceph-objectstore-tool --data-path /var/lib/ceph/osd/ceph-23 --pgid 4.120 '["4.120",{"oid":"rbd_data.18570ff66275d5.000000000000a4da","key":"","snapid":3365,"hash":3096615712,"max":0,"pool":4,"namespace":"","max":0}]' get-bytes rbd_data.18570ff66275d5.000000000000a4da_snapid_3365_bin
   ```
4. 删除异常快照对象：
   ```bash
   ceph-objectstore-tool --data-path /var/lib/ceph/osd/ceph-23 --pgid 4.120 '["4.120",{"oid":"rbd_data.18570ff66275d5.000000000000a4da","key":"","snapid":3365,"hash":3096615712,"max":0,"pool":4,"namespace":"","max":0}]' remove
   ```
5. 启动OSD并观察恢复：
   ```bash
   systemctl start ceph-osd@23
   ```
6. 还原debug级别（若之前调高过）：
   ```bash
   ceph tell osd.23 injectargs '--debug_osd 0/5'
   ```

