# q
OST本地故障恢复时，如何直接重新挂载原有的 Lustre 目标？
# a
使用命令 `mount -t lustre /dev/rbd5 /data/lustre_data`，其中 `/dev/rbd5` 为 OST 对应的块设备，`/data/lustre_data` 为挂载点。

# q
跨节点恢复 OST 时，需要先通过哪个命令清除旧配置并指定新的服务节点与 MGS 节点？
# a
需执行 `tunefs.lustre --erase-params --servicenode=192.168.6.172@tcp --mgsnode=192.168.6.174@tcp /dev/rbd0`，然后再挂载对应设备（例如 `mount -t lustre /dev/rbd5 /data/lustre_data`）。

# q
`tunefs.lustre` 命令中的 `--servicenode` 和 `--mgsnode` 参数分别指什么？
# a
`--servicenode` 用于指定服务节点（servicenode）的网络标识，`--mgsnode` 用于指定 msg 节点（即 MGS，管理服务节点）的网络标识，格式为 `IP地址@tcp`。

