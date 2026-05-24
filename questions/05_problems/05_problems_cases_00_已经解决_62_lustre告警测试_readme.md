# q
如何在Ceph RBD上部署Lustre的MGS（管理服务）并格式化，标准命令是什么？
# a
使用 `mkfs.lustre` 指定文件系统名称为 `nas`，服务类型为 `--mgs`，并配置 MGS 节点地址（双节点冗余）：
```bash
mkfs.lustre --fsname=nas --mgs --mgsnode=192.168.5.219@tcp --mgsnode=192.168.5.220@tcp --reformat /dev/rbd5
```
其中 `/dev/rbd5` 是预先创建的 RBD 块设备映射到本地后的设备路径。

# q
如何使用 `mkfs.lustre` 在 RBD 设备上创建 MDT（元数据目标）并分配索引？
# a
根据节点规划执行带 `--mdt` 和 `--index` 参数的格式化命令，例如在节点 10.5.65.130 上创建索引为 0 的 MDT：
```bash
mkfs.lustre --fsname=nas --mdt --mgsnode=192.168.5.219@tcp --mgsnode=192.168.5.220@tcp --reformat --index=0 /dev/rbd0
```
其他节点类似，只需修改 `--index` 值（如 1、2）。

# q
如何使用 `mkfs.lustre` 在 RBD 设备上创建 OST（对象存储目标）并分配索引，其命令格式是怎样的？
# a
OST 格式化命令需指定 `--ost` 和唯一的 `--index`，并指向对应的 RBD 设备路径。例如在 10.5.65.130 节点上创建索引 0~3 的 OST：
```bash
mkfs.lustre --fsname=nas --ost --mgsnode=192.168.5.219@tcp --mgsnode=192.168.5.220@tcp --reformat --index=0 /dev/rbd1
mkfs.lustre --fsname=nas --ost --mgsnode=192.168.5.219@tcp --mgsnode=192.168.5.220@tcp --reformat --index=1 /dev/rbd2
...
```
所有节点上的 OST 索引必须全局不重复，范围覆盖 0 到 11。

# q
完成 Ceph RBD 上的 Lustre 格式化后，如何挂载该文件系统？
# a
首先创建挂载点目录，然后使用 `mount` 命令指定 Lustre 类型和 MGS 节点地址及文件系统名：
```bash
mkdir -p /mnt/lustre
mount -t lustre 192.168.5.219@tcp:192.168.5.220@tcp:/nas /mnt/lustre
```
该命令可在任意安装了 Lustre 客户端的节点上执行。

