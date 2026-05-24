# q
`ifdown` 和 `ifup` 操作默认会修改系统的本地路由表（`table local`）吗？
# a
默认情况下，`ifdown` 和 `ifup` 关闭和重启网络接口不会直接修改本地路由表（`table local`），因为该表主要包含本地主机地址（如 loopback）和广播地址，与接口的启用/禁用无关。

# q
在什么情况下 `ifdown` 和 `ifup` 操作可能会影响到路由表（包括 `table local`）？
# a
如果网络接口的启动/关闭配置中执行了自定义脚本，并且这些脚本包含了对路由表的修改（例如添加额外的路由规则），那么 `ifdown` 和 `ifup` 就可能间接影响路由表。

# q
如何验证 `ifdown` 和 `ifup` 操作是否改变了 `table local` 的内容？
# a
通过以下步骤对比操作前后的输出：
1. 执行 `ip route show table local` 查看当前状态。
2. 使用 `ifdown` 关闭网络接口，再用 `ifup` 重启。
3. 再次执行 `ip route show table local` 并对比两次输出是否有变化。

