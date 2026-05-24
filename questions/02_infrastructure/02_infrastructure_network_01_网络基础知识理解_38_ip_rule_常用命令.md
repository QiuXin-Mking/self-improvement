# q
如何查看系统中当前定义的路由规则？
# a
使用 `ip rule` 命令。示例输出如下：
```bash
0:      from all lookup local
32766:  from all lookup main
32767:  from all lookup default
```

# q
`ip rule` 输出中每行前的数字（0、32766、32767）代表什么含义？
# a
数字表示路由规则的优先级，数字越小优先级越高。系统在选择路由时，会按优先级从高到低匹配规则。

