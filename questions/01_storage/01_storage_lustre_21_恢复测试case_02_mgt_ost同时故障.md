# q
Lustre 客户端在执行 fio 写操作时大量报错 "Cannot send after transport endpoint shutdown"，典型的根因是什么？
# a
根因是 Lustre 的 MGT（管理目标）和 OST（对象存储目标）同时被卸载（umount），导致客户端与服务端之间的 LNet 传输连接断开，写请求无法送达，从而触发传输端点关闭错误。

# q
从 fio 输出中如何定位因 MGT 和 OST 同时故障导致的 IO 失败？
# a
在 fio 输出中可以看到大量类似以下的错误日志，交替出现 "Cannot send after transport endpoint shutdown" 和 "Input/output error"：
```
fio: io_u error on file /mnt/lustre/test1/Fio.102.0: Cannot send after transport endpoint shutdown: write offset=0, buflen=1048576
fio: io_u error on file /mnt/lustre/test1/Fio.130.0: Input/output error: write offset=0, buflen=1048576
```
这些错误表明客户端无法与后端存储服务正常通信。

# q
解决 MGT 和 OST 同时故障导致客户端 IO 失败的标准恢复流程是什么？
# a
1. 重新挂载（mount）被卸载的 MGT 和 OST 设备。
2. 在客户端卸载（umount）并重新挂载（mount）Lustre 文件系统，以重建与服务端的连接。
3. 重新执行 IO 测试验证恢复情况。

