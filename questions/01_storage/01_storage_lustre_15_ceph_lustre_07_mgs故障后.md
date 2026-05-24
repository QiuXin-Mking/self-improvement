# q
Lustre MGS 故障后，是否需要重新挂载 MDS 和 OST？
# a
根据测试记录，MGS 故障后需要再次挂载 MDS 和 OST 才能恢复正常。具体操作包括先卸载 MGT（管理目标），然后卸载并重新挂载受影响的 OST。

# q
在 MGS 故障恢复过程中，重新挂载一个 OST 的典型命令是什么？
# a
先卸载旧的挂载点，再使用 lustre 类型重新挂载。示例：
```
umount /data/lustre_ost_4
mount -t lustre /dev/rbd4 /data/lustre_ost_4
```
其中 `/dev/rbd4` 是 OST 对应的后端块设备，挂载点为 `/data/lustre_ost_4`。

# q
该测试环境中，MGS 服务由哪个节点提供，对应的挂载点和 RBD 镜像是什么？
# a
MGS 运行在节点 `10.176.102.172`（主机名 `stg_ssd_3-102-172`）上，MGT 挂载点为 `/data/lustre_mgt`，对应的 RBD 镜像为 `lustre_mgt`，位于 pool `lustre_pool` 中。该节点同时运行着一个 MDT（`lustre_mdt0`）和多个 OST。

