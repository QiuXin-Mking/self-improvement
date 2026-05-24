# q
如何查询Ceph OSD当前的恢复相关配置参数？
# a
使用以下命令分别查询各项参数：
```bash
ceph config get osd osd_max_backfills
ceph config get osd osd_recovery_max_active
ceph config get osd osd_recovery_op_priority
ceph config get osd osd_client_op_priority
```
示例输出显示 `osd_max_backfills` 为 2000，`osd_recovery_max_active` 为 2000，`osd_recovery_op_priority` 为 10，`osd_client_op_priority` 为 63。

# q
`osd_recovery_op_priority` 和 `osd_client_op_priority` 的核心区别是什么？
# a
这两个参数控制 I/O 操作的优先级，数值越大优先级越高。`osd_recovery_op_priority` 是恢复操作的优先级，默认值为 10；`osd_client_op_priority` 是客户端业务操作的优先级，默认值为 63。客户端操作优先级远高于恢复操作，从而确保正常业务 I/O 不会被恢复流量挤占。

# q
如何使用 `ceph config set` 持久化修改 OSD 恢复参数？
# a
使用 `ceph config set osd <参数名> <值>` 语法进行持久化设置，修改会写入配置数据库并长期生效。示例：
```bash
ceph config set osd osd_max_backfills 20000
ceph config set osd osd_recovery_max_active 20000
ceph config set osd osd_recovery_op_priority 10
ceph config set osd osd_client_op_priority 63
```

# q
`ceph tell osd.* injectargs` 命令的作用是什么？与 `ceph config set` 有何不同？
# a
该命令用于在运行时动态向所有 OSD 注入参数，即时生效且不重启服务。示例：
```bash
ceph tell osd.* injectargs "--osd_recovery_sleep 0.0001"
ceph tell osd.* injectargs '--osd_recovery_max_active 2'
```
与 `ceph config set` 的区别在于：`injectargs` 是临时性的，重启后失效；`ceph config set` 是持久化配置，会记录到配置库中，重启后依然生效。

# q
`osd_max_backfills` 和 `osd_recovery_max_active` 参数分别控制什么？
# a
`osd_max_backfills` 控制每个 OSD 允许同时进行的回填（backfill）操作的最大数量；`osd_recovery_max_active` 控制每个 OSD 同时活跃的恢复操作的最大数量。两者都用于限制恢复流量，避免对正常业务造成冲击。查询命令分别为 `ceph config get osd osd_max_backfills` 和 `ceph config get osd osd_recovery_max_active`。

