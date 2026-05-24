# q
ll_ost_io 是什么？
# a
`ll_ost_io` 是 Lustre 文件系统中 OST（对象存储目标）上的 I/O 服务线程，负责处理客户端的读写请求。每个 `ll_ost_io` 进程对应一个内核线程，数量可通过 `/sys/fs/lustre/ost/OSS/ost_io/threads_started` 查看。

# q
如何查看当前 OST 上 ll_ost_io 线程的启动数量？
# a
使用以下命令查看已启动的 ll_ost_io 线程总数：
```
cat /sys/fs/lustre/ost/OSS/ost_io/threads_started
```
示例输出 `202` 表示当前已启动 202 个 I/O 线程。

# q
为什么 `ps aux | grep ll_ost_io | wc -l` 的结果比 `threads_started` 多 1？
# a
因为 `ps aux` 输出包含 grep 进程自身，`grep ll_ost_io` 会匹配到该命令行，导致计数多 1。实际活跃的 ll_ost_io 线程数应以 `/sys/fs/lustre/ost/OSS/ost_io/threads_started` 为准。

