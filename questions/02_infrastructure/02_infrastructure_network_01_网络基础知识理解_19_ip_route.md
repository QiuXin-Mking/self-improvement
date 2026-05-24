# q
如何理解路由表中的default路由？
# a
default路由（默认路由）用于处理目标地址不在本地路由表中的所有流量。在示例`default via 172.22.0.1 dev eno1 proto static metric 100`中，下一跳地址为172.22.0.1，出口接口为eno1，该路由是静态路由（proto static），metric值为100。

# q
路由表中的metric值的作用是什么？
# a
metric用于决定路由的优先级，较小的metric值优先级更高。例如示例中默认路由metric为100，不同直接连接路由的metric分别为100、101、102和300，系统会优先选择metric值较小的路由来发送数据包。

# q
ip route输出中的scope link是什么意思？
# a
scope link表示该路由是本地链路路由，意味着目标子网直接连接到指定的网络接口，无需通过下一跳路由器。例如`172.22.0.0/16 dev eno1 proto kernel scope link src 172.22.80.80`表示172.22.0.0/16子网直接连接在eno1接口上，源地址为172.22.80.80。

# q
路由条目中的proto kernel代表什么？
# a
proto kernel表示该路由是由内核自动生成的，通常对应于本地接口配置的IP地址所在的子网。例如`172.22.0.0/16 dev eno1 proto kernel scope link src 172.22.80.80`就是内核为接口eno1生成的路由。

# q
ip route输出中linkdown的含义是什么？
# a
linkdown表示对应的网络接口处于链路断开（down）状态。例如示例中docker0接口的路由标记了linkdown，说明该接口当前未正常工作。

