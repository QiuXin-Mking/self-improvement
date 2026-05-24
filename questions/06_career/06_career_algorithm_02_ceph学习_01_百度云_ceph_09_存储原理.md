# q
Ceph 的 OSD 存储后端主要有哪两种实现方式？各自如何工作？
# a
FileStore 和 BlueStore。FileStore 将 OSD 格式化为文件系统（如 XFS）并挂载到对应目录；BlueStore 借助守护进程机制，绕过传统文件系统直接管理裸设备。

# q
在 Ceph 中，数据从对象到磁盘的映射流程是怎样的？
# a
对象（obj）先映射到 PG（Placement Group），PG 再映射到一组 OSDs，最终以二进制形式存储到磁盘。相关的元数据包括：文件对象与 PG 的属性信息，以及 OSD 对象记录的磁盘数据位置。

# q
Ceph 如何处理大文件的存储？
# a
大文件会被切分成多个对象（切片），然后针对这些切片执行副本操作。

