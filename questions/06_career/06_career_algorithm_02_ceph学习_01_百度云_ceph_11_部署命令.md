# q
如何查看当前主机的网络接口信息？
# a
使用 `ifconfig` 命令查看所有网卡信息。
```bash
ifconfig
```

# q
如何列出所有磁盘分区？
# a
使用 `fdisk -l` 命令列出所有磁盘及分区表信息。
```bash
fdisk -l
```

# q
`blkid` 命令的作用是什么？
# a
`blkid` 用于查看块设备的 UUID、文件系统类型等属性。直接运行即可：
```bash
blkid
```

# q
如何配置 Ceph 集群的 `/etc/hosts` 文件？
# a
编辑 `/etc/hosts` 文件，按以下格式添加节点信息：
```bash
vi /etc/hosts
# 格式：<IP地址> <域名> <主机名> <MON域名>
# 例如：
10.0.0.1 example.com admin mon1
```
保存后可使用 `cat /etc/hosts` 查看。

