# q
如何查看CephFS状态？
# a
```bash
ceph fs status
```

# q
如何将名为nas的CephFS的待机MDS数量设置为0？
# a
```bash
ceph fs set nas standby_count_wanted 0
```

# q
如何停止ceph-mds@ctl15服务？
# a
```bash
systemctl stop ceph-mds@ctl15
```

# q
如何删除名为nas的Ceph文件系统？
# a
```bash
ceph fs rm nas --yes-i-really-mean-it
```

