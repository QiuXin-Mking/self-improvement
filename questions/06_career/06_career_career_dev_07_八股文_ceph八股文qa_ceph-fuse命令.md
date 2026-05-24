# q
ceph-fuse 挂载文件系统的命令是什么？
# a
```bash
ceph-fuse -m 10.1.1.2:6789 /mnt/lustre/ -n client.manila -k /etc/ceph/ceph.client.manila.keyring
```

