# q
在冗余MGS环境中，如何重新配置OST的服务节点和MGS节点信息？
# a
使用`tunefs.lustre`命令清除现有配置并指定新的网络参数：
```bash
tunefs.lustre --erase-params --servicenode=192.168.6.174@tcp --mgsnode=192.168.6.174@tcp --mgsnode=192.168.6.172@tcp --mgsnode=192.168.6.175@tcp /dev/rbd6
```
该命令通常在将OST迁移到新节点或调整MGS地址时执行。

# q
如何从客户端定位特定OST（如OST0003）的sysfs接口并检查其连接状态？
# a
使用`find`命令查找OST相关路径：
```bash
find /sys/fs/lustre -name "*OST0003*"
```
连接状态信息位于`/sys/fs/lustre/osc/<fsname>-OST0003-osc-<client_uuid>/`，其中：
- `active`文件：指示该OSC是否处于活跃状态
- `ost_conn_uuid`：显示连接的服务节点UUID
- `ping`：可查看通信延迟
示例：`/sys/fs/lustre/osc/nas_test-OST0003-osc-ffffa1e0b6df5800/`

# q
在lustre客户端上，如何确认某个OST对应的OSC设备是否已正确挂载并连接到服务节点？
# a
通过检查`/sys/fs/lustre/osc/`下对应OST的符号链接与实际OSP目录。例如：
```bash
ll /sys/fs/lustre/osc/nas_test-OST0003-osc-MDT0000
```
若输出显示该链接指向`../osp/nas_test-OST0003-osc-MDT0000`，则表示该OSC设备已关联至对应的OSP（Object Storage Proxy）。进一步查看该目录下的`active`、`ost_conn_uuid`等属性文件可以验证连接状态和服务UUID。

