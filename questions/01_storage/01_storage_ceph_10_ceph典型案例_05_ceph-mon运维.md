# q
在Ceph集群中，使用ceph-deploy删除某节点上的Monitor服务应执行什么命令？
# a
```bash
ceph-deploy mon destroy controller4
```
执行前需先进入Ceph配置目录（例如 `/etc/ceph_nautilus`）。

# q
在Ceph集群中，使用ceph-deploy新增一个Monitor节点应执行什么命令？
# a
```bash
ceph-deploy mon add controller4
```
执行前需先进入Ceph配置目录（例如 `/etc/ceph_nautilus`）。

# q
执行ceph-deploy的mon相关命令（如add或destroy）时，对当前工作目录有什么要求？
# a
需要先进入Ceph集群的配置文件目录，例如：
```bash
cd /etc/ceph_nautilus
```
该目录下应包含对应的ceph.conf及keyring等文件。

