# q
在蓝海节点扩容中，如何通过RBD创建新的Lustre MDT和OST块设备？
# a
使用 `rbd create` 命令，指定大小和 `--image-feature layering`。新增的MDT大小为200G，OST大小为1T。示例：
```bash
rbd create --size 200G rbd/lustre_mdt3 --image-feature layering
rbd create --size 200G rbd/lustre_mdt4 --image-feature layering
rbd create --size 1T rbd/lustre_ost06 --image-feature layering
rbd create --size 1T rbd/lustre_ost07 --image-feature layering
rbd create --size 1T rbd/lustre_ost08 --image-feature layering
rbd create --size 1T rbd/lustre_ost09 --image-feature layering
```

# q
格式化新增的Lustre MDT和OST时使用的 `mkfs.lustre` 命令及其关键参数是什么？
# a
使用 `mkfs.lustre` 并指定文件系统名、目标类型、MGS节点列表、索引（--index）和 `--reformat`。例如在节点172上：
```bash
# 格式化 MDT（索引3）
mkfs.lustre --fsname=nas_test --mdt --mgsnode=192.168.5.171@tcp --mgsnode=192.168.5.172@tcp --mgsnode=192.168.5.173@tcp --index=3 --reformat /dev/rbd0

# 格式化 OST（索引6、7）
mkfs.lustre --fsname=nas_test --ost --mgsnode=192.168.5.171@tcp --mgsnode=192.168.5.172@tcp --mgsnode=192.168.5.173@tcp --index=6 --reformat /dev/rbd1
mkfs.lustre --fsname=nas_test --ost --mgsnode=192.168.5.171@tcp --mgsnode=192.168.5.172@tcp --mgsnode=192.168.5.173@tcp --index=7 --reformat /dev/rbd2
```
`--fsname` 指定文件系统名称，`--mdt/--ost` 表示目标类型，`--mgsnode` 列表配置管理服务节点，`--index` 设定目标在Lustre中的唯一编号。

# q
新增的Lustre MDT和OST如何在节点上进行挂载？
# a
在相应的节点上直接使用 `mount -t lustre /dev/rbdX <挂载点>`。例如节点172上将映射后的RBD设备挂载到固定目录：
```bash
mount -t lustre /dev/rbd0 /data/lustre_mdt
mount -t lustre /dev/rbd1 /data/lustre_ost_1
mount -t lustre /dev/rbd2 /data/lustre_ost_2
```
节点179使用相同的挂载点命名，绑定各自映射的 `/dev/rbd0`、`/dev/rbd1`、`/dev/rbd2`。

