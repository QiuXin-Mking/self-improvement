# q
如何使用lctl命令查看Lustre MGS上特定目标的配置日志？
# a
使用命令 `lctl --device MGS llog_print <target>`，例如：
```
lctl --device MGS llog_print nas-MDT0000 | grep OST0000
```
输出会显示配置日志条目，每行包含索引（index）、事件类型（event）、设备名（device）、UUID等信息。

# q
如何使用lctl命令删除MGS配置日志中的某条记录？
# a
使用命令 `lctl --device MGS llog_cancel <target> <index>`，其中 `<target>` 为目标名称，`<index>` 为日志条目索引号。例如：
```
lctl --device MGS llog_cancel nas-MDT0000 12
lctl --device MGS llog_cancel nas-MDT0000 13
lctl --device MGS llog_cancel nas-MDT0000 14
```
这些命令会从 `nas-MDT0000` 的配置日志中删除索引为 12、13、14 的记录。

# q
Lustre配置日志中 attach、setup、add_osc 事件分别表示什么操作？
# a
- `attach`：关联一个 OSC 设备类型，为指定 OST 在 MDT 上创建对应的 OSC 实例。
- `setup`：设置该 OSC 设备的 UUID 和节点 LNet 地址（如 `172.31.0.26@tcp`）。
- `add_osc`：将该 OST 的 OSC 添加到 MDT 的 LOV（Logical Object Volume）中，完成 OST 在 MDT 上的挂载配置。

