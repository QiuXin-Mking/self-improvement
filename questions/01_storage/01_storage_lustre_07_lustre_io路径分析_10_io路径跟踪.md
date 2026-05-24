# q
如何在Lustre中开启I/O路径跟踪所需的调试标志？
# a
通过 `lctl set_param debug` 命令设置调试子系统掩码，例如：
```bash
lctl set_param debug="trace iotrace rpctrace nettrace vfstrace"
```
开启后，内核会记录与I/O、RPC、网络、VFS等相关的跟踪信息。可用 `lctl get_param debug` 查看当前设置。

# q
在Lustre中如何使用标记（mark）来定位测试时间区间内的日志？
# a
使用 `lctl mark "标记内容"` 在调试日志中插入带有时间戳的标记，例如：
```bash
lctl mark "0507_1423_mark_start"
# 执行测试操作
lctl mark "0507_1423_mark_stop"
```
日志中会产生类似 `DEBUG MARKER: 0507_1423_mark_start` 的记录，便于提取两个标记之间的所有内核日志。

# q
如何将Lustre内核调试日志导出到文件进行分析？
# a
使用 `lctl debug_kernel` 命令将当前内核调试环形缓冲区的日志内容输出到指定文件：
```bash
lctl debug_kernel /tmp/lustre.log
```
导出后，可以使用编辑器或脚本对日志进行查看和分析。

# q
在开始一次新的I/O路径跟踪前，为什么要执行 `lctl clear` 命令？
# a
`lctl clear` 清空内核调试环形缓冲区中的旧日志，避免与之前的调试信息混淆，确保日志文件中只包含本次测试产生的跟踪数据，使分析更清晰准确。

