# q
部署Ceph前如何关闭防火墙和SELinux？
# a
```bash
systemctl disable --now firewalld
setenforce 0
sed -i 's/^SELINUX=.*/SELINUX=disabled/' /etc/selinux/config
```

# q
如何为Ceph集群配置时间同步？
# a
```bash
yum install -y chrony
systemctl enable --now chronyd
systemctl status --now chronyd
```

# q
如何查询当前系统已安装的Ceph相关软件包？
# a
```bash
rpm -qa | grep ceph
```

