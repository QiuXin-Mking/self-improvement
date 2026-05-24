# q
lfsck_speed_limit 参数的作用是什么？默认值 0 表示什么？
# a
`lfsck_speed_limit` 用于控制 Lustre 文件系统检查（lfsck）的速率限制。设置为 `0` 表示无速率限制，设置为非 `0` 值则表示每秒允许扫描的最大对象数量。该参数存在于 OST 和 MDT 节点的 `/sys/fs/lustre/` 对应目录下，MGS 节点上不存在此参数。

# q
在 Linux 中，fsck 是什么？通常在什么条件下使用？
# a
`fsck`（file system check）是 Linux 中用于检查和修复本地文件系统一致性的工具。通常在单用户模式、修复模式下，且对应块设备处于卸载（umount）状态时使用。

# q
lsblk 命令输出中的 RM 和 RO 字段分别表示什么含义？
# a
- `RM`：指示设备是否可移除，`1` 表示可移除，`0` 表示不可移除。
- `RO`：指示设备是否只读，`1` 表示只读，`0` 表示可读写。

# q
lfsck_namespace 是什么？如何定位其路径？
# a
`lfsck_namespace` 是 Lustre 文件系统中与命名空间检查（lfsck）相关的调试接口。可以通过以下命令找到其路径：
```bash
find / -name 'lfsck_namespace'
```
通常位于 `/sys/kernel/debug/lustre/mdd/<MDT名称>/lfsck_namespace`。

