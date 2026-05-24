# q
MGS停止服务一天后再次挂载MGT时，mount失败并报“Cannot send after transport endpoint shutdown”的典型根因是什么？
# a
当MGS长时间不在线，客户端和服务器上已有的MGC连接到MGS的内核传输端点会被关闭（transport endpoint shutdown）。重新挂载MGT时，内核试图通过MGC向MGS注册目标（MDT/OST），但因MGS之前断开导致底层连接状态异常，返回错误并挂载失败。日志中常见错误码-108（可能是ENOTCONN）或-110（ETIMEDOUT）。

# q
如何从日志中确认MGS不可达导致“cannot register this server with the MGS”错误？
# a
在挂载MGT或MDT时的内核日志中，查找如下关键行：
```
LustreError: 15f-b: nas_test-MDT0032: cannot register this server with the MGS: rc = -108. Is the MGS running?
```
如果还出现以下日志，则进一步证实MGS连接已丢失：
```
LustreError: 137-5: nas_test-MDT0000_UUID: not available for connect from 192.168.6.179@tcp (no target)
```
或者MGC被驱逐的日志：
```
Lustre: Evicted from MGS (at MGC192.168.6.172@tcp_0) after server handle changed from ... to ...
```

# q
解决“MGS不在服务后MGT无法挂载”的标准恢复流程是什么？
# a
1. 确保MGS服务已运行：如果是独立MGT设备，需重新挂载MGT。例如：
   ```
   mount -t lustre /dev/rbd0 /data/lustre_mgt
   ```
2. 在所有使用该文件系统的客户端上卸载Lustre：
   ```
   ansible lustre -m shell -a 'umount /mnt/lustre'
   ```
3. 如果挂载MGT时仍失败，先卸载所有可能残留的MGT/MDT/OST，再重新挂载MGT；必要时重新加载Lustre内核模块或重启节点以清理残留的MGC连接。
4. 检查nodemap等配置是否正常（可选）：
   ```
   lctl nodemap_info all
   lctl get_param nodemap.BirdResearchSite.idmap
   ```

# q
在重新挂载MGS时，日志中出现“Evicted from MGS... after server handle changed”是什么含义？
# a
该日志表示MGC客户端检测到MGS服务端的句柄（server handle）发生了变化（通常是MGS重启或重新挂载所致），原有连接被强制驱逐。这是MGS恢复后的正常行为，MGC会随后自动重建连接，后续的目标注册和挂载操作可以正常完成。例如：
```
Lustre: Evicted from MGS (at MGC192.168.6.172@tcp_0) after server handle changed from 0x35b4f0e21d2a103a to 0x35b4f0e21d2a2dde
```

