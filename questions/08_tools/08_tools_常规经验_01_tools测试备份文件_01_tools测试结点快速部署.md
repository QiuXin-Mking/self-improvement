# q
启动 multipathd 服务时遇到 `ConditionPathExists=/etc/multipath.conf was not met` 错误应如何解决？
# a
multipathd 服务要求 `/etc/multipath.conf` 文件必须存在。执行 `vi /etc/multipath.conf` 创建文件（即使为空内容），保存退出后即可正常启动服务。

# q
在 `/etc/multipath.conf` 中如何为多路径设备设置友好的别名？请给出一个 multipaths 段落的配置示例。
# a
在 `multipaths { }` 块中添加 `multipath { wwid ... alias 别名 }`，例如：
```
multipaths {
        multipath {
           wwid    3600b342095668d2d9132d970ad0000d9
           alias    mpatha
        }
        multipath {
            wwid 3600b3426fb82c96d983ddfc59d0000d8
            alias   mpathb
        }
}
```

# q
在 RHEL/CentOS 中为一个网卡配置静态 IP、网关和子网前缀需要修改哪个文件？请给出一个管理口配置的示例。
# a
修改 `/etc/sysconfig/network-scripts/ifcfg-<接口名>`（如 `ifcfg-ens192`），示例：
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
如何安装 iSCSI initiator 并设置其开机自动启动？
# a
安装：
```bash
yum install -y iscsi-initiator-utils
# 或
yum install open-iscsi
```
设置开机自启：
```bash
chkconfig iscsi on
chkconfig iscsid on
```
修改配置文件 `vi /etc/iscsi/iscsid.conf` 后重启服务：
```bash
service iscsi restart
```

