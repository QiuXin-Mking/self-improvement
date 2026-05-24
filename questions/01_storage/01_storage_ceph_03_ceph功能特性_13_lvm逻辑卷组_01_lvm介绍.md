# q
什么是逻辑卷管理器（LVM）？
# a
LVM（Logical Volume Manager）是 Linux 上用于管理磁盘空间的工具，能在线动态调整磁盘分区大小、创建快照，并灵活地管理存储设备，而无需中断系统运行。

# q
LVM 包含哪些关键组件，各自的作用是什么？
# a
- **物理卷（PV）**：实际的磁盘分区或整个硬盘，是存储的基础单元。
- **卷组（VG）**：由一个或多个物理卷组成的存储池，为逻辑卷提供空间。
- **逻辑卷（LV）**：从卷组中分配出来的、类似传统分区的存储单元，可动态调整大小。

# q
如何用命令创建一个 10GB 的逻辑卷？
# a
1. 创建物理卷：`sudo pvcreate /dev/sdX1`
2. 创建卷组：`sudo vgcreate my_volume_group /dev/sdX1`
3. 创建逻辑卷：`sudo lvcreate -L 10G -n my_logical_volume my_volume_group`

# q
扩展逻辑卷的命令步骤是什么？
# a
1. 扩展逻辑卷：`sudo lvextend -L +5G /dev/my_volume_group/my_logical_volume`
2. 调整文件系统大小（如 ext4）：`sudo resize2fs /dev/my_volume_group/my_logical_volume`

# q
减小逻辑卷时，必须先执行哪一步？完整的减小步骤是什么？
# a
必须先缩小文件系统。步骤：
1. 缩小文件系统：`sudo resize2fs /dev/my_volume_group/my_logical_volume 5G`
2. 缩减逻辑卷：`sudo lvreduce -L 5G /dev/my_volume_group/my_logical_volume`

