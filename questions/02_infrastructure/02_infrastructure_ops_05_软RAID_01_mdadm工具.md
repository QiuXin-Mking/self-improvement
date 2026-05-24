# q
mdadm 是什么工具？
# a
mdadm 是用于构建、管理和监控 Linux 软 RAID 设备的命令行工具，支持创建、装配、调整、监控等多种模式。

# q
RAID 0 的核心原理是什么？
# a
RAID 0 的核心机制是**条带化**，数据被分割成小块（条带）并均匀分布到阵列中的所有磁盘上，实现并行读写，从而显著提高性能。

# q
RAID 0 的存储容量如何计算？
# a
RAID 0 的存储容量是所有成员磁盘容量的总和，不存在任何容量损失。

# q
使用 mdadm 创建 RAID 0 的基本命令是什么？
# a
```bash
sudo mdadm --create --verbose /dev/md0 --level=0 --raid-devices=2 /dev/sdb /dev/sdc
```
该命令创建名为 `/dev/md0` 的 RAID 0 阵列，使用两块磁盘，并输出详细过程信息。

