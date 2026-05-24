# q
使用 iptables 允许 192.168.1.0/24 网段访问本机 TCP 1089 端口的命令是什么？
# a
```bash
sudo iptables -A INPUT -p tcp -s 192.168.1.0/24 --dport 1089 -j ACCEPT
```

# q
在设置 iptables 时，如何确保只有允许的网段能访问 1089 端口并拒绝其他所有来源？
# a
在添加 ACCEPT 规则之后，再添加一条 DROP 规则：
```bash
sudo iptables -A INPUT -p tcp --dport 1089 -j DROP
```
注意 DROP 规则必须放在 ACCEPT 规则之后。

# q
在 Debian/Ubuntu 系统上如何持久化保存 iptables 规则？
# a
推荐安装 `iptables-persistent`，然后使用以下命令保存：
```bash
sudo apt-get install iptables-persistent
sudo netfilter-persistent save
```
或直接将规则写入文件：
```bash
sudo sh -c "iptables-save > /etc/iptables/rules.v4"
```

# q
在 CentOS/RHEL 7 及以上版本中，若坚持使用 iptables（而非 firewalld），如何安装并持久化保存规则？
# a
先安装 `iptables-services`，启用并启动服务，然后保存规则：
```bash
sudo yum install iptables-services
sudo systemctl enable iptables
sudo systemctl start iptables
sudo service iptables save
```
注意直接使用 `iptables-save`/`restore` 可能会被 firewalld 覆盖。

