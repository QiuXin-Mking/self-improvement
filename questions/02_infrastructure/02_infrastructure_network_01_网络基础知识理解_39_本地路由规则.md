# q
什么是本地路由规则？
# a
本地路由规则（Local Route）写明 IP 和网口的连接关系。

# q
如何查看本地路由表？
# a
使用命令：
```sh
ip route show table local
```

# q
`ip route show table local` 的输出中通常包含哪些类型的路由条目？
# a
至少包含 `broadcast` 和 `local` 两种类型的路由，每条记录指明了 IP 地址、绑定的网络接口、协议和源地址，例如：
```sh
broadcast 172.22.0.0 dev eno1 proto kernel scope link src 172.22.251.105
local 172.22.251.105 dev eno1 proto kernel scope host src 172.22.251.105
```

