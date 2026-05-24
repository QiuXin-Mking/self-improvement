# q
在Linux系统中，使用哪个命令可以查看服务器的网关信息？
# a
使用 `ip route show` 命令查看路由表，其中包含默认网关信息；若要查看特定接口（如 eth0）的网关，可使用 `ip route show dev eth0`。

# q
如何使用 `route` 命令查看路由表及默认网关？
# a
执行 `route -n` 命令（需安装 net-tools 软件包），它会以数字形式显示路由表，包含默认网关。

# q
在 `ip route show` 输出中，如何识别默认网关？
# a
找到以 `default via` 开头的行，或目标地址为 `0.0.0.0/0` 的行，该行中 `via` 后面的 IP 地址即为默认网关。例如 `default via 172.22.0.1 proto static metric 108` 表示网关是 `172.22.0.1`。

