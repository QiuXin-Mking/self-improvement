# q
使用 `ceph-deploy` 部署 Ceph 集群的核心步骤是什么
# a
核心步骤包括：
1. `ceph-deploy new <mon节点>` 生成初始配置和密钥
2. `ceph-deploy install <节点>` 在所有节点上安装 Ceph 软件包
3. `ceph-deploy mon create-initial` 创建初始 Monitor
4. `ceph-deploy admin <节点>` 将管理密钥和配置推送到各节点
5. `ceph-deploy mgr create <节点>` 创建 Manager 守护进程

常用参数：`--public-network`、`--cluster-network` 指定网络，`--overwrite-conf` 强制覆盖配置。

# q
如何彻底清理通过 `ceph-deploy` 部署的 Ceph 集群
# a
依次执行以下三条命令，以完全清除集群：
```
ceph-deploy purge ceph1 ceph2 ceph3
ceph-deploy purgedata ceph1 ceph2 ceph3
ceph-deploy forgetkeys
```
`purge` 删除软件包和配置，`purgedata` 清空集群数据，`forgetkeys` 清理本地密钥缓存。

# q
在 Ceph 中如何通过 LVM 创建一个包含 DB 和 WAL 分区的 Bluestore OSD
# a
先通过 LVM 将磁盘划分为三个逻辑卷（WAL、DB、Block），再用 `ceph-volume` 组合创建 OSD：
```
pvcreate /dev/sdb
vgcreate ceph-vg1 /dev/sdb
lvcreate -n ceph-wal -L 2G ceph-vg1
lvcreate -n ceph-db  -L 8G ceph-vg1
lvcreate -n ceph-block -l 100%FREE ceph-vg1
ceph-volume lvm create --bluestore \
  --data /dev/ceph-vg1/ceph-block \
  --block.db /dev/ceph-vg1/ceph-db \
  --block.wal /dev/ceph-vg1/ceph-wal
```
确保 `/var/lib/ceph/bootstrap-osd/ceph.keyring` 存在且权限为 600。

# q
使用 `ceph-deploy` 时因 pip 版本过旧导致安装旧包，如何解决
# a
卸载现有 `ceph-deploy` 后重新安装：
```
pip uninstall ceph-deploy -y
pip install ceph-deploy
```
建议先升级 pip 版本，确保获取最新包，以避免兼容性问题。

