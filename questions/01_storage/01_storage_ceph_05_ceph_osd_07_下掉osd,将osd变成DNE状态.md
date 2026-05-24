# q
在Ceph中将OSD彻底下掉并变为DNE状态需要执行哪些步骤？
# a
依次执行以下命令（以osd.$id为例）：
```bash
id=20; ceph osd down osd.$id; sleep 1; ceph osd out osd.$id; sleep 1; ceph osd rm osd.$id; sleep 1; ceph auth del osd.$id; umount /var/lib/ceph/osd/ceph-$id
```
步骤顺序为：mark down → mark out → remove from crush map → delete auth → umount OSD数据目录。

# q
ceph osd tree中显示OSD状态为DNE意味着什么？
# a
DNE（Does Not Exist）表示该OSD曾经存在于集群中，但已被彻底移除（已执行`ceph osd rm`、删除认证并卸载目录），crush map中不再将其视为活跃或存在的OSD。

