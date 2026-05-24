# q
如何从Lustre dmesg日志定位网络超时问题？
# a
搜索日志中的 `ptlrpc_expire_one_request` 关键字。该日志表明RPC请求因等待慢回复而超时，是连接中断或雪崩的典型前兆。关键日志片段：
```
[Wed Nov  5 17:59:06 2025] Lustre: 70872:0:(client.c:2295:ptlrpc_expire_one_request()) @@@ Request sent has timed out for slow reply: [sent 1762337280/real 1762337280]  req@00000000f178e0a7 x1828787560993280/t0(0) o41->st_nas-MDT0002-osp-MDT0000@192.168.5.173@tcp:24/4 lens 224/368 e 0 to 1 dl 1762337287 ref 1 fl Rpc:XQr/0/ffffffff rc 0/-1 job:'u0.g0'
```
超时后通常伴随连接丢失和恢复日志，如 `Connection to ... was lost` 和后续的 `Connection restored`。

# q
Lustre日志中出现大量"deleting orphan objects"表明服务器端发生了什么？
# a
表明服务器端检测到连接中断并触发了恢复（recovery）机制，正在清理因客户端异常断开而产生的孤儿对象。这通常发生在连接恢复后，相关日志如：
```
[Wed Nov  5 18:15:26 2025] Lustre: st_nas-OST0008: deleting orphan objects from 0x540000401:1107896 to 0x540000401:1107937
```
此现象常见于节点重启、网络故障或雪崩后的恢复过程。

# q
Lustre服务启动时报`can't mount /dev/rbd1: -5`错误，常见原因是什么？
# a
错误码 `-5` 是 `EIO`（输入/输出错误），通常表明底层块设备（如Ceph RBD）无法正常访问。可能是后端存储集群不可达、RBD映射异常或权限问题。关键日志：
```
[Wed Nov  5 18:40:20 2025] LustreError: 14930:0:(osd_handler.c:8111:osd_mount()) st_nas-OST0010-osd: can't mount /dev/rbd1: -5
[Wed Nov  5 18:40:20 2025] LustreError: 14930:0:(obd_mount_server.c:1993:server_fill_super()) Unable to start osd on /dev/rbd1: -5
```
解决需检查后端存储连接状态和块设备可用性。

# q
Lustre Imperative Recovery的日志"recovery window shrunk from 300-900 down to 150-900"含义是什么？
# a
Imperative Recovery（命令式恢复）被启用后，恢复窗口从默认的300-900秒缩小到150-900秒，加速了客户端重连的超时判定。这意味着服务端会更快地要求客户端在指定时间内重连，否则清理其资源。日志示例：
```
[Wed Nov  5 18:40:14 2025] Lustre: st_nas-MDT0002: Imperative Recovery enabled, recovery window shrunk from 300-900 down to 150-900
[Wed Nov  5 18:40:16 2025] Lustre: st_nas-MDT0002: Will be in recovery for at least 2:30, or until 13 clients reconnect
```

