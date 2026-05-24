# q
这个脚本的作用是什么？
# a
检查网卡配置文件 `/etc/sysconfig/network-scripts/ifcfg-eth0` 中是否存在 `ONBOOT` 关键字，并输出相应提示。

# q
脚本中 `sed -n '/ONBOOT/p'` 命令的含义是什么？
# a
`sed -n` 禁止默认输出，`'/ONBOOT/p'` 只打印包含 `ONBOOT` 的行；整体上该命令会静默查找并打印匹配行，如果找到返回匹配内容（退出状态为 0），否则无输出（退出状态为非 0）。

# q
在 RHEL/CentOS 系统中，`ONBOOT` 参数的作用是什么？
# a
`ONBOOT` 参数控制网卡是否在系统启动时自动激活，取值为 `yes` 或 `no`。

