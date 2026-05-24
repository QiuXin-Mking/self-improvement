# q
在Linux上安装multipath并启用服务的完整命令是什么？
# a
```sh
yum install -y device-mapper-multipath
systemctl enable multipathd
systemctl start multipathd
systemctl status multipathd
```
还必须创建 `/etc/multipath.conf` 配置文件，否则服务无法启动。

# q
multipathd服务启动失败，日志显示“ConditionPathExists=/etc/multipath.conf was not met”是什么原因？
# a
multipathd服务要求 `/etc/multipath.conf` 文件必须存在，否则服务不会启动。即使文件为空，也需要创建该文件。

# q
有 `/etc/multipath.conf` 配置和没有配置时，`lsblk` 中多路径设备的名称有何区别？
# a
无配置文件时，多路径设备名使用WWID长格式（如 `35000cca0545b9890`）；有配置文件并设置了 `user_friendly_names yes` 和 `multipaths { ... alias ... }` 时，设备名会变成用户友好的别名（如 `mpathb`）。

