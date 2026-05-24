# q
如何查看当前系统的策略路由规则表？
# a
```bash
ip rule list
```

# q
如何添加一条基于源IP的策略路由规则，使来自特定IP的流量使用指定路由表？
# a
```bash
ip rule add from 172.240.134.236 table 100
```

# q
如何删除已添加的策略路由规则？
# a
```bash
ip rule del from 172.240.134.236 table 100
```

# q
如何为策略路由规则指定自定义优先级？
# a
```bash
ip rule add priority 100 from 192.168.1.0/24 table 100
```

# q
当为同一源IP添加多个路由表规则时，优先级如何确定？
# a
后添加的规则优先级更高（数字更小），会排在前面优先匹配。例如第一次添加 `ip rule add from 172.240.134.236 table 100` 后，规则序号为32765；第二次添加 `ip rule add from 172.240.134.236 table 101` 后，规则序号变为32764，表示优先级更高。

