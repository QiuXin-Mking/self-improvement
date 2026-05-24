# q
ip rule 命令的基本作用是什么？
# a
`ip rule` 是 Linux 中用于管理路由策略规则的命令。它允许管理员定义数据包匹配条件（SELECTOR）以及匹配后使用的路由表（ACTION），从而影响数据包的路由决策，常用于高级网络配置和定制。

# q
如何添加一条路由规则，使来自 192.168.1.0/24 网段的数据包使用路由表 100？
# a
```bash
ip rule add from 192.168.1.0/24 table 100
```

# q
如何查看当前系统上的所有路由策略规则？
# a
```bash
ip rule list
```

# q
如何清空所有的路由策略规则？
# a
使用 `flush` 子命令清空所有路由规则：
```bash
ip rule flush
```
（注意：该操作会删除所有规则，请谨慎执行。）

# q
如何获取 ip rule 的详细使用说明？
# a
查看 ip-rule 的手册页：
```bash
man ip-rule
```

