# q
在树莓派上固定IP地址，需要编辑哪个配置文件？
# a
编辑 `/etc/dhcpcd.conf` 文件，例如使用命令：
```
sudo nano /etc/dhcpcd.conf
```

# q
如何为树莓派的 eth0 接口配置静态 IP 地址？
# a
在 `/etc/dhcpcd.conf` 中添加以下配置：
```
interface eth0
static ip_address=192.168.1.2/24
static routers=192.168.1.1
static domain_name_servers=192.168.1.1
```

# q
树莓派静态 IP 配置中的 `static routers` 和 `static domain_name_servers` 分别代表什么？
# a
- `static routers`：指定默认网关，例如 `192.168.1.1`
- `static domain_name_servers`：指定 DNS 服务器地址，例如 `192.168.1.1`

