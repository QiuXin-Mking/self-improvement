# q
Ceph RBD 中如何为一个镜像创建快照？
# a
使用 `rbd snap create <pool>/<image>@<snap-name>` 命令，例如：
```
rbd snap create rbd/cyy_nas@snapshot_cyy_nas
```

# q
如何删除 Ceph RBD 镜像的指定快照？
# a
使用 `rbd snap rm <pool>/<image>@<snap-name>` 命令，例如：
```
rbd snap rm rbd/cyy_nas@snapshot_cyy_nas
```

# q
在测试快照操作时，常用哪些 Ceph 命令查看集群和存储状态？
# a
常用命令包括：
- `ceph df`：查看集群存储使用情况
- `df -h`：查看客户端挂载点使用情况
- `ceph -s`：查看集群整体健康状态
- `ceph pg dump | grep snap`：过滤查看与快照相关的 PG 信息

