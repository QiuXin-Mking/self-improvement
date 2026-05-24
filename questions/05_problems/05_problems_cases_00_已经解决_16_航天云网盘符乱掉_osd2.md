# q
如何通过Docker检查Ceph OSD容器的设备挂载配置？
# a
执行命令 `docker inspect ceph_osd_sdd | grep OSD`，查看环境变量中的 `OSD_DEVICE`、`OSD_DEVICE_BLOCK_DB`、`OSD_DEVICE_BLOCK_WAL`、`OSD_DEVICE_BLOCK_CACHE` 等配置。

# q
OSD容器 `ceph_osd_sdd` 启动后立即退出（ExitCode 1），典型根因是什么？
# a
容器环境变量配置的设备与实际盘符不匹配。例如 `OSD_DEVICE=/dev/bcache2`，但实际 OSD 块设备可能映射到了 `bcache3`，导致 OSD 无法正确访问底层存储。

# q
如何定位 Ceph OSD 盘符错乱时块设备映射关系错误？
# a
1. 使用 `lsblk` 查看物理盘、bcache 设备和 LVM 逻辑卷的树状关系，例如 `sdc -> sdc1 -> bcache3 -> ceph--...-osd-block--...`。
2. 检查 OSD 目录下的符号链接：`ls -l /var/lib/ceph/osd/ceph-<id>/`，观察 `block`、`block.db`、`block.wal` 链接到的设备路径。
3. 通过 `readlink -f` 或查看 `/dev/disk/by-partuuid/` 确认符号链接最终指向的分区（如 `nvme0n1p3`、`nvme0n1p7`），并与 Docker 环境变量进行比对。

