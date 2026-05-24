# q
lsblk工具的核心作用是什么？
# a
lsblk（list block devices）用于以树状结构列出所有可用的块设备信息，包括磁盘、分区及其挂载点，方便快速查看存储设备的层次关系和挂载情况。

# q
如何使用lsblk查看特定磁盘的序列号(SN)？
# a
使用 `-o` 选项指定输出列，例如：
```shell
lsblk -o SERIAL /dev/sdb
```
该命令会输出 `/dev/sdb` 磁盘的序列号。可以替换 `/dev/sdb` 为其他磁盘设备路径，如 `/dev/sdc`。

