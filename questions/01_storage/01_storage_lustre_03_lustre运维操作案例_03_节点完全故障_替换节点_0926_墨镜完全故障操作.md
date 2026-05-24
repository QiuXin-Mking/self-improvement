# q
在Lustre节点完全故障时，如何使用新节点恢复MDT和OST服务？
# a
1. 使用 `rbd map` 将RBD镜像映射为本地块设备（如 /dev/rbd0）。
2. 创建对应的挂载目录（如 /data/mj_mdt01）。
3. 使用 `tunefs.lustre --erase-params --servicenode=新节点IP@tcp --mgsnode=MGS节点IP@tcp` 清除原有配置并写入新的服务节点和管理节点参数。
4. 使用 `mount -t lustre /dev/rbdX /data/目标目录` 挂载Lustre文件系统。

# q
`tunefs.lustre --erase-params` 命令在节点替换中的作用是什么？
# a
它用于清除Lustre目标（MDT/OST）上原有的所有配置参数（如旧的网络地址、MGS节点等），然后重新设置 `--servicenode` 为新节点的NID以及 `--mgsnode` 为MGS的NID，使目标可以在新节点上注册并提供服务。

# q
`rbd map rbd/mj_mdt01` 命令执行后，设备映射路径如何确定？
# a
映射后内核会分配一个未使用的 `/dev/rbd` 设备号，通常按顺序分配为 `/dev/rbd0`、`/dev/rbd1` 等。可以通过 `rbd showmapped` 查看映射关系。

