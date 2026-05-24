# q
Lustre文件系统的数据总容量如何计算？
# a
Lustre文件系统是**所有文件（OST）的容量总和**，即所有对象存储目标（OST）提供的空间之和。

# q
OSS和OST上如何实现负载均衡？
# a
Lustre在OSS（对象存储服务器）和OST（对象存储目标）上内置了负载均衡机制，确保数据、元数据和请求被均匀分布，避免热点。

# q
查看Lustre文件系统中inode使用情况的命令是什么？
# a
使用 `lfs df -i` 命令可以查看每个OST和MDT的inode总数、已用及可用数量。

# q
resize2fs命令在Lustre场景中用于什么？
# a
`resize2fs` 用于调整底层逻辑卷上的本地文件系统（如ldiskfs）大小，通常在扩容MDT或OST的块设备后执行，使文件系统识别新增的空间。

