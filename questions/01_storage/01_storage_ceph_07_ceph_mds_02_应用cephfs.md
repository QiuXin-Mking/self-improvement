# q
创建CephFS文件系统前需要准备哪些存储池？
# a
需要分别创建元数据池和数据池，示例命令：
```bash
ceph osd pool create cephfs_metadata 32
ceph osd pool create cephfs_data 128
```

# q
如何基于已有的存储池新建一个CephFS文件系统？
# a
使用 `ceph fs new` 命令，指定文件系统名称、元数据池和数据池：
```bash
ceph fs new cephfs cephfs_metadata cephfs_data
```
验证创建结果：
```bash
ceph mds stat
```

# q
CephFS中的子卷组和子卷如何创建与管理？
# a
创建子卷组：
```bash
ceph fs subvolumegroup create cephfs mygroup
```
查询子卷组：
```bash
ceph fs subvolumegroup ls cephfs
```
在组内创建子卷：
```bash
ceph fs subvolume create cephfs test_subvol2 --size 10G --group_name mygroup
```
查询某组下所有子卷：
```bash
ceph fs subvolume ls cephfs --group_name mygroup
```

