# q
在Ceph集群中重建SATA OSD前，需要设置哪些标志来防止自动数据迁移和集群将OSD踢出？
# a
需要设置三个集群标志：`nobackfill`、`norecover` 和 `noout`。命令为：
```bash
ceph osd set nobackfill
ceph osd set norecover
ceph osd set noout
```
这些标志阻止集群在手动操作 OSD 时进行数据回填、恢复和将停止的 OSD 标记为 out。

# q
从Ceph集群中安全删除一个Bluestore OSD并清理磁盘的完整步骤是什么？
# a
标准删除及清理流程（假设 OSD ID 为变量 `$id`）：
1. 停止 OSD 服务：`systemctl stop ceph-osd@$id`
2. 标记 down 并删除：`ceph osd down osd.$id; sleep 1; ceph osd rm osd.$id`
3. 删除认证：`ceph auth del osd.$id`
4. 从 CRUSH map 移除：`ceph osd crush remove osd.$id`
5. 卸载挂载点：`umount /var/lib/ceph/osd/ceph-$id`
6. 清除分区（如适用）：`parted /dev/device -s rm 1` 和 `parted /dev/device -s rm 2`
7. 清理磁盘头：`dd if=/dev/zero of=/dev/device bs=10M count=50`

# q
在Ceph集群中，如何重新准备并激活一个Bluestore OSD（支持独立WAL设备）？
# a
使用 `ceph-disk` 工具完成重建：
- 若仅使用主数据盘：
  ```bash
  ceph-disk prepare --zap --bluestore /dev/sdc
  ceph-disk activate /dev/sdc1
  ```
- 若需指定独立的 WAL 设备（通过 partuuid 引用）：
  ```bash
  ceph-disk prepare --zap --bluestore --block.wal /dev/disk/by-partuuid/$wal_partition_uuid_link /dev/sdc
  chown -R ceph:ceph /dev/disk/by-partuuid/$wal_partition_uuid_link
  ceph-disk activate /dev/sdc1
  ```
`--zap` 选项会清除磁盘上的原有分区和数据。

# q
OSD重建完成后，如何将其重新加入CRUSH map并调整权重？
# a
使用两条命令完成CRUSH重新加入和权重调整：
1. 调整OSD权重（通常先设为 1.0）：`ceph osd reweight osd.$id 1.0`
2. 将OSD添加到指定主机桶下，并设定 CRUSH 权重（例如1.45，根据磁盘容量调整）：
   `ceph osd crush add osd.$id 1.45 host=storage23`

