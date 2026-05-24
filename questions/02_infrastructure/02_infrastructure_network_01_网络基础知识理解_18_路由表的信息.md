# q
如何查看Linux主机的路由表信息？
# a
使用命令 `cat /proc/net/route` 可以查看内核当前的路由表，每一行代表一个路由项。也可以使用 `ip route` 命令进行查询。

# q
Linux系统下 `/proc/net/route` 文件中的路由表是如何存储的？
# a
该文件的内容是内核运行时动态生成的，并非静态存储在磁盘上，它反映了当前内核网络子系统的路由状态。

# q
如何在系统重启后保留路由配置？
# a
需要将路由设置写入持久化配置，例如使用 `ip route` 命令配合脚本，或在 `/etc/sysconfig/network-scripts/`（取决于发行版）下的接口配置文件中定义静态路由。NetworkManager 或 systemd-networkd 等网络管理工具也会使用自己的配置文件来管理路由。

