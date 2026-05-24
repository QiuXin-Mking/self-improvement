# q
如何查看当前系统所有 iptables 规则？
# a
使用以下命令，输出包含详细信息和端口号：
```bash
sudo iptables -L -v -n
```

# q
如何允许特定 IP 地址（如 192.168.1.100）访问服务器？
# a
在 `INPUT` 链尾部追加一条允许规则：
```bash
sudo iptables -A INPUT -s 192.168.1.100 -j ACCEPT
```
- `-A INPUT`：追加到 INPUT 链
- `-s 192.168.1.100`：匹配源 IP
- `-j ACCEPT`：接受匹配的流量

# q
如何拒绝特定 IP 地址（如 192.168.1.200）访问服务器？
# a
在 `INPUT` 链尾部追加一条丢弃规则：
```bash
sudo iptables -A INPUT -s 192.168.1.200 -j DROP
```
- `-A INPUT`：追加到 INPUT 链
- `-s 192.168.1.200`：匹配源 IP
- `-j DROP`：丢弃匹配的流量

# q
如何允许特定 IP 访问服务器的特定 TCP 端口（如 80）？
# a
追加一条带协议和端口匹配的允许规则：
```bash
sudo iptables -A INPUT -p tcp -s 192.168.1.100 --dport 80 -j ACCEPT
```
- `-p tcp`：指定协议为 TCP
- `-s 192.168.1.100`：源 IP
- `--dport 80`：目标端口 80
- `-j ACCEPT`：接受匹配的流量

# q
如何在 CentOS/RHEL 上保存 iptables 规则，使其在重启后仍然生效？
# a
执行以下命令保存当前规则：
```bash
sudo service iptables save
```

