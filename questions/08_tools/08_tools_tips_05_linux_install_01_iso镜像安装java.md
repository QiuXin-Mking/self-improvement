# q
如何使用 mount 命令将 CentOS 7 的 ISO 镜像挂载到 /mnt/cdrom 目录？
# a
```bash
mount -o loop -t iso9660 /root/CentOS-7-x86_64-DVD-2003.iso /mnt/cdrom
```

# q
在 CentOS-Media.repo 配置中，baseurl 需要添加什么路径才能从 /mnt 挂载点安装软件？
# a
baseurl 中必须包含 `file:///mnt/` 这一行。

# q
挂载 ISO 镜像后，刷新 yum 缓存的完整命令是什么？
# a
```bash
yum clean all
yum makecache
```

# q
如何实现开机时自动将光盘挂载到 /mnt？
# a
在 `/etc/fstab` 中添加以下行：
```bash
/dev/cdrom /mnt iso9660 defaults 0 0
```

