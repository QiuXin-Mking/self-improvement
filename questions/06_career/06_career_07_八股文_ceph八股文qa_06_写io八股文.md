# q
写 IO 有哪些类型？
# a
1 覆盖写
2 对齐写
3 非对齐写
4 追加写

# q
Linux 中的文件系统架构分为哪些层次？
# a
用户态：应用程序
内核态：VFS 页面缓存 文件系统 IO 驱动
硬件：机械盘 SSD

# q
POSIX 中用于删除文件的语义是什么？以覆盖写为例
# a
truncate
这时候需要先把文件 truncate 为 0，表示把文件原来所有的数据删掉，接着再写入。

# q
VFS 的四个基本元素是什么？
# a
inode， dentry ， superblock  file

# q
dentry 对象在 VFS 中的作用是什么？
# a
dentry 对象则是用于记录文件的结构关系的，也就是不同目录的上下级层级结构关系和树状关系等

