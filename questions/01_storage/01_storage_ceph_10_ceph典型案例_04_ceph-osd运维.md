# q
在重建OSD之前为什么要执行`ceph osd set nobackfill`和`ceph osd set norecover`？
# a
为了防止在重建过程中发生数据回填和恢复，避免对集群性能和重建操作造成干扰，保证操作环境稳定。

# q
删除一个Ceph OSD的完整步骤包含哪些操作？
# a
1. 将OSD标记为down并out：`ceph osd down osd.$id`、`ceph osd out osd.$id`
2. 从CRUSH map中移除：`ceph osd rm osd.$id`
3. 删除认证密钥：`ceph auth del osd.$id`
4. 卸载对应的数据目录：`umount /var/lib/ceph/osd/ceph-$id`

# q
如何彻底清理OSD上的bcache设备及其残留数据？
# a
1. 删除LVM：`lvremove -y $lv_path`，`vgremove -y $vg_name`；若失败可用`dmsetup remove`或`dmsetup remove_all`
2. 删除bcache：`bcachectl destroyosdbcache <osd-id> --yes-i-really-mean-it`
3. 清理缓存残留：停止并擦除cache和data设备，必要时强制擦除头部数据，并重新加载bcache模块
```bash
bcachectl stop $cset_uuid
bcachectl stop $data_device
bcachectl wipe $data_device
bcachectl stop $cache_device
bcachectl wipe $cache_device
# 如果仍有残留，强制擦除缓存分区头部数据并重载模块
dd if=/dev/zero of=<cache_device> bs=1M count=50
rmmod bcache && partprobe && modprobe bcache
```

# q
重建bcache缓存映射时，`make-bcache`命令的正确参数是什么？
# a
```bash
dd if=/dev/zero of=$cache_device bs=1M count=10
dd if=/dev/zero of=$data_device bs=1M count=10
make-bcache --writeback -B $data_device -C $cache_device
partprobe
```
参数说明：`--writeback`启用写回缓存模式，`-B`指定后端数据设备，`-C`指定缓存设备。

# q
使用`ceph-volume lvm create`创建OSD时，如何指定独立的WAL和DB设备？
# a
使用`--block.wal`和`--block.db`参数指定分区，例如：
```bash
ceph-volume lvm create --osd-id 6 --data /dev/sdb --block.wal /dev/nvme1n1p1 --block.db /dev/nvme0n1p1
```
创建前确保在ceph.conf中启用了相关选项：`osd_crush_update_on_start = true`、`osd_class_update_on_start = true`。

