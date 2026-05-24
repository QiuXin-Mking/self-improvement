# q
挂载MGT后，`lfs check all`输出的服务状态发生了怎样的变化？
# a
挂载MGT前，OST仅显示后端OSC（如`osc-ffff9a217fe3f000`）；挂载后，每个OST新增与MDT0000的连接（`osc-MDT0000`），同时MDT的元数据客户端（mdc）出现并最终全部变为`active`，标志着文件系统元数据服务恢复正常。

# q
如何使用`tunefs.lustre`擦除旧参数并重新指定Lustre的MGS节点？
# a
使用命令：
```
tunefs.lustre --erase-params --servicenode=192.168.6.175@tcp --mgsnode=192.168.6.175@tcp /dev/rbd1
```
`--erase-params`清除原有配置，`--servicenode`和`--mgsnode`分别指定服务节点和MGS节点的网络标识。

# q
如何取消设备上的MGS角色并指向新的MGS节点？
# a
使用命令：
```
tunefs.lustre --nomgs --mgsnode=new_mgs_nid /dev/mdt-device
```
`--nomgs`移除该设备上的MGS功能，`--mgsnode`设置新的MGS节点NID。

