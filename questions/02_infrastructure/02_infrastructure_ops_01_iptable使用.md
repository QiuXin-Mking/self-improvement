# q
如何使用iptables允许特定IP段访问特定端口并拒绝其他所有访问？
# a
先添加允许规则：`iptables -A INPUT -p tcp --dport <端口> -s <源IP段> -j ACCEPT`，再添加拒绝规则：`iptables -A INPUT -p tcp --dport <端口> -j DROP`。必须严格按此顺序，因为iptables按规则顺序处理数据包。

# q
如何查看当前iptables的所有规则？
# a
常用查看命令：
- `iptables -L`：列出规则
- `iptables -L -v -n`：以详细模式列出，不进行DNS反向解析
- `iptables -t filter -L -v -n`：指定filter表查看
- `iptables -S`：以命令格式显示当前规则

# q
如何保存iptables规则使其在系统重启后仍然生效？
# a
使用 `iptables-save > /etc/sysconfig/iptables` 将当前规则保存到配置文件（适用于Red Hat/CentOS）。恢复规则可使用 `iptables-restore < /etc/sysconfig/iptables`。不同发行版配置文件路径可能不同。

