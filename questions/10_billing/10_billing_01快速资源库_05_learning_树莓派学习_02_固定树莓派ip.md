# q
在树莓派中，用于配置静态 IP 地址的配置文件是什么？
# a
配置文件为 `/etc/dhcpcd.conf`。

# q
在 `/etc/dhcpcd.conf` 中，如何为有线网卡（eth0）设置静态 IP 地址？
# a
添加以下配置：
```
interface eth0
static ip_address=192.168.1.2/24
static routers=192.168.1.1
static domain_name_servers=192.168.1.1
```
分别指定 IP 地址（含子网掩码）、网关和 DNS 服务器。

