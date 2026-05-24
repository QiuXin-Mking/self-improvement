# q
回忆下inode操作集合有哪些？ext3
# a
操作         作用                                                对应 ext3 函数
create      在目录中创建常规文件                                 ext3_create
lookup      根据文件名查找 inode                                 ext3_lookup
link        创建硬链接（增加 inode 链接计数）                   ext3_link
unlink      删除硬链接（减少链接计数，可能触发删除）           ext3_unlink
symlink     创建符号链接                                         ext3_symlink
mkdir       创建目录                                             ext3_mkdir
rmdir       删除空目录                                           ext3_rmdir
mknod       创建设备文件、FIFO、socket 等特殊文件               ext3_mknod
rename      重命名文件或目录                                     ext3_rename
setattr     设置 inode 属性（chmod/chown/truncate 等）          ext3_setattr
getattr     获取 inode 属性（供 stat 调用）                     ext3_getattr
permission  检查访问权限                                         ext3_permission

# q
debugfs 应该下载哪个包
# a
apt install e2fsprogs

