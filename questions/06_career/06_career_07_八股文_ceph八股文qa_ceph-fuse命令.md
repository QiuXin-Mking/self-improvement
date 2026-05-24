# q
如何挂载Ceph文件系统？
# a
```bash
ceph-fuse -m 10.1.1.2:6789 /mnt/lustre/ -n client.manila -k /etc/ceph/ceph.client.manila.keyring
```

