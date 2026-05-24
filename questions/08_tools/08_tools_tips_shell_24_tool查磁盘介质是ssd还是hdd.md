# q
用什么命令可以查看系统磁盘是固态硬盘(SSD)还是机械硬盘(HDD)？
# a
使用 `lsblk -d -o name,rota` 命令。其中 `-d` 表示不显示分区，只显示磁盘设备，`-o name,rota` 只输出名称和旋转属性。

# q
在 `lsblk -d -o name,rota` 输出中，rota 值为 0 和 1 分别代表什么类型的磁盘？
# a
rota 值为 **0** 表示固态硬盘（SSD），rota 值为 **1** 表示机械硬盘（HDD）。

# q
执行 `lsblk -d -o name,rota` 时需要注意什么？
# a
该命令可能需要足够的权限，建议使用 `sudo` 前缀，例如 `sudo lsblk -d -o name,rota`。

