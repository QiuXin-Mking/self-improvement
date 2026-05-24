# q
在Linux中查看路由表的推荐命令是什么？
# a
推荐使用 `ip route` 命令，它已取代旧的 `route` 命令。查看路由表：
```bash
ip route
```
如果需要数字格式的地址，可以使用旧命令 `route -n`，但该命令已被废弃。

# q
如何理解 `ip route` 输出中的默认路由条目？
# a
默认路由以 `default` 开头，用于处理所有未被其他明确路由匹配的流量。例如：
```
default via 172.22.0.1 dev eno1 proto static metric 100
```
该条目表示所有默认流量通过网关 `172.22.0.1` 经接口 `eno1` 发出，该路由是静态配置的，度量值为 100。

# q
路由表中的 metric 值有什么作用？
# a
metric（度量值）代表路由的优先级，数值越低优先级越高。当存在多条到达同一目的地的路由时，系统会选用 metric 值最小的路由进行转发。

# q
`traceroute` 命令的功能是什么？
# a
`traceroute` 用于跟踪数据包从源到目的地所经过的路径，逐跳显示路由信息。例如：
```bash
traceroute example.com
```
它会输出每一跳的 IP 地址，帮助诊断网络连通性。该命令通常需要 `sudo` 权限执行。

