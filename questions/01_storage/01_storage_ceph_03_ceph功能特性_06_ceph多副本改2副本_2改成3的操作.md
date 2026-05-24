# q
如何使用Ceph命令将所有存储池的副本数批量设置为2？
# a
```bash
for pool in $(ceph osd pool ls); do
  ceph osd pool set "$pool" size 2
done
```

# q
如何从Ceph集群中彻底移除一个OSD（例如osd.3）？
# a
先用 `ceph osd crush remove osd.3` 从CRUSH图中移除，再执行 `ceph osd rm 3` 删除OSD认证和数据引用。

