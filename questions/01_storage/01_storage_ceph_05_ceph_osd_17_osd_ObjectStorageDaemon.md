# q
如何安全地停止并销毁一个Ceph OSD（例如 osd.2）？
# a
先用 `systemctl stop ceph-osd@2` 停止服务，然后执行 `ceph osd destroy osd.2 --yes-i-really-mean-it` 从集群中移除并销毁该OSD。

# q
删除OSD后，需要执行哪些步骤来彻底清理磁盘（例如 /dev/vdb）？
# a
1. `umount /var/lib/ceph/osd/ceph-2` 卸载挂载点。
2. `wipefs --all --force /dev/vdb` 擦除文件系统签名。
3. 若有 device mapper 残留则执行 `dmsetup remove`（按需指定设备名）。
4. `ceph-volume lvm zap /dev/vdb` 清除 LVM 元数据。
5. `dd if=/dev/zero of=/dev/vdb bs=5M count=100` 覆盖磁盘起始区域，确保完全清空。

# q
在 /etc/ceph/ceph.conf 中设置 `osd_crush_update_on_start = true` 和 `osd_class_update_on_start = true` 的作用是什么？
# a
允许 OSD 在启动时自动将其自身添加到 CRUSH map 并更新其设备类别，避免手动执行 `ceph osd crush add/set` 等操作，使新OSD能自动加入CRUSH拓扑。

# q
如何使用 `ceph-volume` 创建一个ID为2且使用独立WAL和DB设备的OSD？
# a
```bash
ceph-volume lvm create --osd-id 2 --data /dev/sdb --block.wal /dev/nvme1n1p1 --block.db /dev/nvme0n1p1
```
该命令在主数据设备 `/dev/sdb` 上创建OSD，并将 WAL 写入 `/dev/nvme1n1p1`，DB 写入 `/dev/nvme0n1p1`，实现高速日志分离。

