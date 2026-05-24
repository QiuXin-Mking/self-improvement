# q
如何创建一个用于 RBD 的 Ceph 存储池？
# a
使用命令 `ceph osd pool create <pool_name> <pg_num>` 创建，例如：
```
ceph osd pool create mypool 128
```
其中 `128` 是 PG 数量。

# q
如何创建一个 RBD 镜像？
# a
先创建存储池，然后使用 `rbd create <pool>/<image> --size <size_MB>` 创建，例如：
```
ceph osd pool create rbd 10
rbd create rbd/cyy_nas --size 20
```
也可以指定大小单位为 GiB：`--size 204800` 表示 200 GiB。

# q
如何将 RBD 镜像映射到本地块设备并使用？
# a
1. 映射：`rbd map <pool>/<image>`，例如 `rbd map mypool/myimage`，会输出设备名如 `/dev/rbd0`。
2. 格式化：`mkfs.ext4 /dev/rbd0`
3. 挂载：`mkdir /mnt/myrbd && mount /dev/rbd0 /mnt/myrbd`
之后就可以像本地磁盘一样读写文件。

# q
如何为 RBD 镜像创建、列出和删除快照？
# a
- 创建快照：`rbd snap create <pool>/<image>@<snap_name>`
  例如：`rbd snap create rbd/cyy_nas@snapshot_cyy_nas`
- 列出快照：`rbd snap ls <pool>/<image>`，例如：`rbd snap ls rbd/cyy_nas`
- 删除快照：`rbd snap rm <pool>/<image>@<snap_name>`
  例如：`rbd snap rm rbd/cyy_nas@cyy_nas_snap`

