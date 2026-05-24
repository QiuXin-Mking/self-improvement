# q
rbd rm vms/rbd.xxxx 发现删不掉应该怎么办
# a
/etc/ceph/ceph.conf debug_rbd 设置成20
然后再触发

# q
rados -p vms stats rbd.xxxx 发现卡住，怎么查询呢。
# a
debug_rados = 20
再触发

