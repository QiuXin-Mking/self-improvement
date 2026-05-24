# q
如何查看Linux系统中的路由策略数据库（路由规则）？
# a
使用 `ip rule show` 命令。示例输出：
```bash
[root@openeuler-repo network-scripts]# ip rule show
0:      from all lookup local
32765:  from all lookup 10
32766:  from all lookup main
32767:  from all lookup default
```

# q
如何查看自定义路由表 `table 10` 中的具体路由条目？
# a
使用 `ip route show table 10` 命令。示例输出：
```bash
[root@openeuler-repo network-scripts]# ip route show table 10
192.168.234.112 via 172.22.102.96 dev eno1
```

# q
在节点2上，如何查看主路由表当前内容？
# a
使用 `ip route show` 命令（默认显示 main 表）。示例输出：
```bash
[root@openeuler-repo network-scripts]# ip route show
default via 172.22.0.1 dev eno1 proto static metric 101
172.22.0.0/16 dev eno1 proto kernel scope link src 172.22.6.65 metric 101
```

