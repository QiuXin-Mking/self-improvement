# q
在Ubuntu/Debian系统上，安装NFS服务器和客户端分别需要执行什么命令？
# a
服务器端安装：
```bash
sudo apt-get install nfs-kernel-server
```
客户端安装：
```bash
sudo apt-get install nfs-common
```

# q
NFS服务器配置文件 `/etc/exports` 中的一条典型共享目录设置 ` /srv/nfs/share *(rw,sync,no_subtree_check)` 各字段含义是什么？
# a
- `/srv/nfs/share`：共享目录的路径。
- `*`：允许访问的客户端，`*` 表示所有客户端。
- `rw`：读写权限。
- `sync`：数据同步写入磁盘。
- `no_subtree_check`：关闭子树检查以提高性能。

# q
客户端如何手动挂载NFS共享目录？假设服务器IP为 `192.168.1.100`，共享路径为 `/srv/nfs/share`，客户端挂载点为 `/mnt/nfs_share`。
# a
```bash
sudo mkdir -p /mnt/nfs_share
sudo mount 192.168.1.100:/srv/nfs/share /mnt/nfs_share
```

# q
如何实现NFS客户端开机自动挂载共享目录？请在 `/etc/fstab` 中添加相应条目。
# a
在 `/etc/fstab` 文件中添加如下行：
```
192.168.1.100:/srv/nfs/share /mnt/nfs_share nfs defaults 0 0
```

