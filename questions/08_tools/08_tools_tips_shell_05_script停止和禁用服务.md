# q
如何使用 sed 命令将网口配置文件中的 ONBOOT 参数修改为 yes？
# a
```bash
sed -i 's/^ONBOOT=.*/ONBOOT=yes/' /etc/sysconfig/network-scripts/ifcfg-eth0
```

# q
如何使用 systemctl 停止并禁用 NetworkManager 服务？
# a
```bash
systemctl stop NetworkManager
systemctl disable NetworkManager
```

# q
在 RHEL/CentOS 中，如何使用 systemctl 重启网络服务？
# a
```bash
systemctl restart network
```

# q
脚本中将 BOOTPROTO 设置为 none 的含义是什么？
# a
表示禁用动态 IP 配置协议（如 DHCP），通常用于手动配置静态 IP 的场景，`BOOTPROTO=none` 与 `BOOTPROTO=static` 在很多系统中效果相同。

