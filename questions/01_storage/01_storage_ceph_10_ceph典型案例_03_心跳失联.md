# q
Ceph OSD 心跳失联通常是什么原因导致的？
# a
当 OSD 自身操作超时（如 osd_op_tp 卡在 do_request->do_op->do_pg_op 中），该 OSD 就不会向其他节点回复心跳包，其他节点因此判定其失联并上报无心跳。常见触发原因包括：
- 特定的时间点（如 1-6 点）在执行 scrub 任务
- 客户端（如 cinder 服务）正在执行 rados ls 等操作
- osd_op_tp 线程数设置过小，无法及时处理积压的操作

# q
如何缓解由 scrub 和线程数不足导致的 Ceph OSD 心跳超时？
# a
可调整以下配置参数：
```ini
[osd]
osd_op_thread_timeout = 30
filestore_op_thread_timeout = 120
osd_scrub_sleep = 0.5
```
`osd_scrub_sleep` 让 scrub 操作间歇睡眠，降低对请求线程的占用；增大超时值可为操作提供更多完成时间，避免过早判定超时。

