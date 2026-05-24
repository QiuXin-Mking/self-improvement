# q
如何在 Linux 中通过配置文件为网卡添加默认路由？
# a
编辑网卡对应的路由配置文件（如 `/etc/sysconfig/network-scripts/route-eth0`），在其中添加 `default via <网关IP>`，保存后重启网络服务（`sudo systemctl restart network`）。

# q
网卡路由配置文件的命名规则是什么？
# a
文件名为 `route-网卡名称`，例如网卡为 `eth0` 时，配置文件路径为 `/etc/sysconfig/network-scripts/route-eth0`。

# q
默认路由在路由配置文件中的典型格式是什么？
# a
使用 `default via <网关IP>` 格式，例如：
```
default via 192.168.1.1
```

