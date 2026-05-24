# q
如何检查Ceph OSD中不健康的工作线程数？
# a
使用命令 `ceph daemon osd.<id> perf dump cct`，查看返回结果中 `"cct"` 部分的 `"unhealthy workers"` 字段值。例如：
```json
"cct": {
    "total_workers": 21,
    "unhealthy workers": 0
}
```

# q
OSD日志出现 heartbeat timeout 时，推荐使用什么方法分析卡住的OSD操作线程？
# a
立即对 OSD 进程执行 `pstack <osd_pid>`，可以查看当时卡住的 osd op 线程正在处理的任务，有助于定位问题原因。

# q
为了捕获心跳超时后 OSD 的详细日志，应将日志级别调整到什么值？有什么局限性？
# a
需要将 `debug_osd` 调整为 `20`。但注意该设置只能记录后续再次发生卡住时的信息，当前正在发生的卡顿无法通过此方式回溯。

