# q
如何为 CentOS 节点的管理网卡配置静态 IP 地址？
# a
在 `/etc/sysconfig/network-scripts` 目录下编辑对应网卡配置文件（如 `ifcfg-ens192`），设置以下参数：
```
TYPE=Ethernet
BOOTPROTO=static
DEFROUTE=yes
IPADDR=172.22.101.111
GATEWAY=172.22.0.1
PREFIX=16
NAME=ens192
DEVICE=ens192
ONBOOT=yes
```

# q
如何为后端存储网卡配置静态 IP（不设置默认路由）？
# a
编辑后端网卡配置文件（如 `ifcfg-ens256`），关键是将 `DEFROUTE` 设为 `no`，避免干扰管理流量路由：
```
TYPE=Ethernet
BOOTPROTO=static
DEFROUTE=no
IPADDR=192.168.234.111
PREFIX=16
NAME=ens256
DEVICE=ens256
ONBOOT=yes
```

# q
multipathd 服务启动失败的常见原因是什么？如何解决？
# a
multipathd 启动失败并提示 `ConditionPathExists=/etc/multipath.conf was not met`，原因是缺少配置文件。只需创建 `/etc/multipath.conf`（即使为空文件）即可解决：
```
vi /etc/multipath.conf
:wq
```
然后重启服务即可正常启动。

# q
multipath.conf 配置文件包含哪些关键段？
# a
配置文件包含四个核心段：
- `blacklist { }`：排除不需要多路径的设备（通过 WWID 指定）。
- `defaults { }`：全局默认参数，如 `user_friendly_names yes`、`polling_interval`、`no_path_retry` 等。
- `devices { }`：针对特定存储厂商的设备参数，如 `vendor "MacroSAN"`、`path_grouping_policy`、`prio alua` 等。
- `multipaths { }`：为多路径设备定义易记的别名，如将 WWID `3600b342...` 映射为 `mpatha`。

# q
配置 multipath 后，磁盘设备名会发生什么变化？
# a
- 未配置或使用默认配置时，多路径设备会显示为基于 WWID 的冗长名称，例如 `35000cca06e0a15bc` 及其分区 `35000cca06e0a15bc1`、`35000cca06e0a15bc2`。
- 配置 `multipaths { alias ... }` 并使用 `user_friendly_names yes` 后，设备会显示为简洁的别名如 `mpathe`，分区为 `mpathe1`、`mpathe2`，便于识别和管理。

