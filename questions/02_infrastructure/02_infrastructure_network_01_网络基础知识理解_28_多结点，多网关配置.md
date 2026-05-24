# q
如何在Linux系统中配置多网关路由？
# a
需要编辑 `/etc/iproute2/rt_tables` 添加自定义路由表（如 `252 GATEWAY1`，`253 GATEWAY2`），然后用 `ip route add default via <网关> dev <接口> table <表名>` 为每个网关添加默认路由，最后使用 `ip rule add from <源IP> table <表名>` 设置策略路由，根据源地址选择路由表。

# q
多网关配置中如何定义和使用新的路由表？
# a
在 `/etc/iproute2/rt_tables` 文件中添加数字标识和表名（例如 `252 GATEWAY1`），之后可通过 `ip route add ... table GATEWAY1` 将路由添加到该表，并通过 `ip rule` 引用该表。

# q
`ip rule` 命令在多网关配置中的作用是什么？
# a
用于设置策略路由规则，通过匹配条件（如源IP `from IP1`）指定要使用的路由表（`table GATEWAY1`），决定不同来源的数据包走哪个网关。

# q
多网关路由配置完成后如何验证正确性？
# a
使用 `ip route show` 查看各路由表的路由条目，使用 `ip rule show` 确认策略路由规则是否正确，确保无冲突且每个接口与网关的路由表匹配。

