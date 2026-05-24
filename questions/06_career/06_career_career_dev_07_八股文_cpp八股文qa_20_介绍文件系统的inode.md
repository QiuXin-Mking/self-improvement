# q
介绍下文件系统的inode
# a
inode包括权限、属主、大小、时间戳、数据块指针 无文件名

访问路径本质是：​路径解析 → 通过目录项找到 inode 号 → 读 inode → 按块指针定位数据。

# q
元数据包括哪些
# a
文件类型/权限（i_mode）
属主 UID/GID
大小 i_size
时间戳 atime/mtime/ctime
链接数（i_nlink）

# q
inode 常见问题
# a
1. inode耗尽
2. 频繁更新小文件导致 a/m/c time写放大， mount -o noatime relatime减低压力。

