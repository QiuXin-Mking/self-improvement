# q
lctl工具的核心作用是什么？
# a
lctl是Lustre的多功能管理命令，覆盖网络配置、设备操作、调试控制、池管理、屏障、快照、节点映射（nodemap）、变更日志、持久客户端缓存、LFSCK、LLOG等子系统。

# q
lctl --list-commands 输出的主要功能类别有哪些？
# a
分类包括：metacommands、control、network config、obd device selection/operations、debugging control、Pools、Barrier、Snapshot、Nodemap、Changelogs、Persistent Client Cache、LFSCK、LLOG，以及已废弃(obsolete)和测试(testing)命令。

# q
如何查看Lustre文件系统上的所有池列表？
# a
执行 `lctl pool_list <文件系统名称>`，例如 `lctl pool_list lustre1`。

# q
如何列出当前Lustre系统中的所有nodemap？
# a
使用命令 `lctl nodemap_info list`，输出会显示存在的nodemap名称（如 `nodemap.default`）。

# q
nodemap信息在proc文件系统中的对应路径是什么？
# a
位于 `/proc/fs/lustre/nodemap` 目录。

