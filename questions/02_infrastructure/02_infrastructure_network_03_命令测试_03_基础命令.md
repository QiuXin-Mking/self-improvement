# q
如何重启 NetworkManager 服务并检查其运行状态？
# a
使用 `systemctl` 命令进行重启和状态检查：
```bash
systemctl restart NetworkManager
systemctl status NetworkManager
```

# q
如何查看特定路由表中的路由条目？
# a
使用 `ip route show table <表标识符>` 命令，例如查看表 10 的路由：
```bash
ip route show table 10
```

# q
如何使用 `ip rule` 和 `ip route` 命令为特定源 IP 添加策略路由？
# a
1. 添加策略路由规则，指定源 IP 使用特定的路由表：
   ```bash
   ip rule add from 192.168.234.65 table 10
   ```
2. 向该路由表中添加路由条目，例如将目标网络/主机的流量指向下一跳：
   ```bash
   ip route add 192.168.234.112/32 via 172.22.6.65 table 10
   ```
   通用格式：
   ```bash
   ip route add <目标网络> via <下一跳地址> table <路由表标识符>
   ```

# q
如何查看当前所有的路由策略规则？
# a
使用 `ip rule list` 或 `ip rule` 命令查看所有路由策略规则：
```bash
ip rule list
```

