# q
在Ceph部署中，创建OSD之前通常需要先切换到哪个目录，为什么？
# a
切换到 `/etc/ceph` 目录，因为该目录存放着集群配置文件（ceph.conf）和密钥环，`ceph-deploy` 默认会从此目录读取配置。

# q
如何使用 `ceph-deploy` 命令为节点 `node01` 创建OSD，并将 `/dev/sdc` 磁盘作为OSD数据设备？
# a
```bash
ceph-deploy osd create --data /dev/sdc node01
```
`--data` 参数指定本地磁盘设备作为OSD的存储后端，`node01` 是目标节点的主机名。

