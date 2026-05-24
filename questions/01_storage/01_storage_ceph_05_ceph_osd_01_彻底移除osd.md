# q
如何彻底从Ceph集群中移除一个OSD？
# a
1. 将OSD标记为out：`ceph osd out osd.X`  
2. 从CRUSH图中移除：`ceph osd crush remove osd.X`  
3. 从OSD map中删除：`ceph osd rm osd.X`  
4. 删除对应的OSD密钥环等残留信息

# q
`ceph osd purge` 命令的作用是什么？
# a
`ceph osd purge osd.X --yes-i-really-mean-it` 是一个组合命令，可以一次性自动完成从CRUSH map移除、从OSD map删除、清理密钥环等彻底移除OSD所需的全部操作，替代多步手动操作。

# q
当OSD存在于CRUSH map但不在OSD map时，应该如何清理？
# a
执行 `ceph osd crush remove osd.X` 将其从CRUSH图中移除，然后运行 `ceph osd rm osd.X` 清理OSD map中的残留记录，并确保删除对应的密钥环。也可以直接使用 `ceph osd purge osd.X --yes-i-really-mean-it` 统一清理。

