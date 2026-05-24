# q
在Lustre多冗余MGS环境中，umount主MGS后，其他节点如何恢复MGC连接？
# a
将MGT设备在备用节点上映射并挂载。
恢复成功的特征日志为：
```
Lustre: MGC192.168.6.172@tcp: Connection restored to  (at 192.168.6.174@tcp)
```
（表示指向原主IP的MGC连接已转移到备用节点恢复）

# q
多冗余MGS故障切换后，原主MGS节点（172）为何会报错“Configuration from log failed from MGS -5”并无法挂载客户端？
# a
因为原主节点上MGS服务已卸载，客户端挂载时仍尝试连接本地之前的MGS IP获取配置数据，导致通信失败。
关键日志：
```
LustreError: MGC192.168.6.172@tcp: Confguration from log nas_test-client failed from MGS -5.
Communication error between node & MGS, a bad configuration, or other errors.
Lustre: Unmounted nas_test-client
LustreError: llite: Unable to mount <unknown>: rc = -5
```

# q
如何从日志定位Lustre客户端因MGS故障导致的挂载失败与恢复过程？
# a
**恢复成功**的日志特征：
```
Lustre: MGC<原IP>@tcp: Connection restored to (at <备用IP>@tcp)
```
**挂载失败/配置获取失败**的日志特征（dmesg/syslog）：
```
LustreError: MGC<IP>@tcp: Confguration from log <fsname>-client failed from MGS -5.
Communication error between node & MGS, a bad configuration, or other errors.
```
随后会出现客户端卸载 `Unmounted` 以及挂载返回码 `rc = -5`。

# q
多冗余MGS的Lustre文件系统，mkfs.lustre命令应如何配置MGS节点列表？
# a
**MGS设备**格式化时使用 `--servicenode` 列出所有冗余节点：
```bash
mkfs.lustre --fsname=nas_test --mgs --servicenode=192.168.6.172@tcp --servicenode=192.168.6.174@tcp --servicenode=192.168.6.175@tcp --reformat /dev/rbd5
```
**MDT/OST设备**格式化时使用 `--mgsnode` 指定相同节点：
```bash
mkfs.lustre --fsname=nas_test --mdt --mgsnode=192.168.6.172@tcp --mgsnode=192.168.6.174@tcp --mgsnode=192.168.6.175@tcp --index=0 --reformat /dev/rbd0
```

