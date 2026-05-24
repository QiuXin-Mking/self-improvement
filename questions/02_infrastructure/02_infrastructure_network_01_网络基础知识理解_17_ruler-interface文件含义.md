# q
/etc/sysconfig/network-scripts/rule-eno1 文件的作用是什么
# a
该文件用于配置与网络接口 eno1 关联的策略路由规则，允许根据源 IP 等条件将流量路由到自定义路由表，而非仅使用默认路由表。

# q
rule-eno1 文件中的策略路由规则采用什么格式
# a
格式为 `from <源IP范围> table <路由表名>`，例如：
```
from 192.168.1.0/24 table custom_table
```
表示来自 192.168.1.0/24 的流量将使用名为 custom_table 的路由表。

# q
如何查看 rule-eno1 文件的内容
# a
使用 `cat` 或其他文本查看命令：
```bash
cat /etc/sysconfig/network-scripts/rule-eno1
```

