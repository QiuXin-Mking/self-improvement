# q
在Linux多网卡环境下，为什么所有网卡配置文件中DEFROUTE参数只能有一个设置为yes？
# a
因为系统只能存在一条默认路由。若多个网卡都设置`DEFROUTE=yes`，会导致路由冲突，网络通信异常。应当只将连接外部网络（如路由器）的接口设为`DEFROUTE=yes`，其余接口设为`no`。

# q
如何正确配置多网卡环境下的默认路由？
# a
在网卡配置文件（如`/etc/sysconfig/network-scripts/ifcfg-ethX`）中，设置：
- 管理口（如22网口，连接路由器）：
```
DEFROUTE=yes
```
- 其他所有网卡：
```
DEFROUTE=no
```
确保只有作为出口的接口启用默认路由。

