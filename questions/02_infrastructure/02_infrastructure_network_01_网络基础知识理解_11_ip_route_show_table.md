# q
查看系统所有路由表的核心要点是什么
# a
使用命令 `ip route show table all` 可以显示系统上所有路由表的信息。

# q
local 路由表中有哪些关键特征
# a
`ip route show table local` 输出的条目由内核自动生成，每条记录代表一个接口的本地地址或广播地址：  
- 类型分为 `local`（本地接口 IP）和 `broadcast`（网络广播地址）  
- `scope host` 表示仅本机可见的本地地址，`scope link` 表示链路本地范围的广播地址  
- 协议均为 `kernel`，源地址（`src`）为对应接口的 IP，例如：
```
broadcast 192.168.0.0 dev bond192 proto kernel scope link src 192.168.80.80
local 192.168.80.80 dev bond192 proto kernel scope host src 192.168.80.80
broadcast 192.168.255.255 dev bond192 proto kernel scope link src 192.168.80.80
```

# q
lo 回环接口有哪些关键特征
# a
- 类型为 `LOOPBACK`，状态 `UP` 且 `LOWER_UP`  
- `mtu 65536`，队列调度为 `noqueue`，状态显示 `UNKNOWN`  
- IPv4 地址为 `127.0.0.1/8`，`scope host`，有效生存时间永久  
- IPv6 地址为 `::1/128`，`scope host`，有效生存时间永久
```
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
    inet6 ::1/128 scope host
       valid_lft forever preferred_lft forever
```

# q
main 路由表中默认路由和直连路由的关键特征是什么
# a
`ip route show table main` 显示主要路由表的条目：  
- 默认路由：`default via <网关> dev <接口> proto static metric <值>`（例如 `default via 172.22.0.1 dev eno1 proto static metric 100`）  
- 直连路由：网络前缀通过具体接口，`proto kernel scope link src <接口IP> metric <值>`，metric 反映了接口优先级，如：
```
172.22.0.0/16 dev eno1 proto kernel scope link src 172.22.80.80 metric 100
172.250.0.0/16 dev enp175s0f3 proto kernel scope link src 172.250.80.80 metric 102
172.251.0.0/16 dev enp175s0f1 proto kernel scope link src 172.251.80.80 metric 101
192.168.0.0/16 dev bond192 proto kernel scope link src 192.168.80.80 metric 300
```

