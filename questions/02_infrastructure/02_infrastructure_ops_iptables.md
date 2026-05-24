# q
如何使用 iptables 允许特定 IP 段访问某个端口并阻止其他所有 IP 的访问？
# a
先执行允许规则，再执行拒绝规则，顺序不可颠倒：
```bash
sudo iptables -A INPUT -p tcp --dport <端口号> -s <IP地址/子网掩码> -j ACCEPT
sudo iptables -A INPUT -p tcp --dport <端口号> -j DROP
```
将 `<端口号>` 替换为实际端口，`<IP地址/子网掩码>` 替换为允许的 IP 地址段。

# q
如何将当前的 iptables 规则持久化保存到文件？
# a
使用 `iptables-save` 命令将规则保存到文件，例如：
```bash
sudo iptables-save > /etc/iptables/rules.v4
```

