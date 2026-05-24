# q
NFS的基本组成部分有哪些？
# a
NFS由服务器（提供共享资源）、客户端（访问共享资源）和导出（Export，服务器将目录设置为可被其他机器访问）组成。

# q
在Ubuntu服务器上安装NFS服务的命令是什么？
# a
```bash
sudo apt-get install nfs-kernel-server
```

# q
如何在/etc/exports文件中配置一个允许所有客户端读写访问的共享目录？
# a
```bash
/srv/nfs/share  *(rw,sync,no_subtree_check)
```
- `/srv/nfs/share`：共享目录路径
- `*`：允许所有客户端
- `rw`：读写权限
- `sync`：数据同步写入磁盘
- `no_subtree_check`：关闭子树检查以提高性能

# q
客户端如何挂载NFS共享目录？
# a
创建挂载点并执行mount命令：
```bash
sudo mkdir -p /mnt/nfs_share
sudo mount 192.168.1.100:/srv/nfs/share /mnt/nfs_share
```
其中`192.168.1.100`为NFS服务器IP地址。

# q
如何在客户端实现NFS共享目录的自动挂载？
# a
在客户端的`/etc/fstab`文件中添加如下行：
```bash
192.168.1.100:/srv/nfs/share /mnt/nfs_share nfs defaults 0 0
```

