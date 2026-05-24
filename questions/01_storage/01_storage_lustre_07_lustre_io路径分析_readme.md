# q
如何启用Lustre全部内核调试并将日志输出到文件？
# a
使用以下命令：
```
lctl set_param debug=-1
lctl debug_kernel /tmp/lustre.log
```
`debug=-1` 打开所有调试子系统和级别，`debug_kernel` 将内核日志保存到指定文件。

# q
如何使用ftrace追踪所有Lustre内核模块函数调用？
# a
执行以下命令：
```
echo function > /sys/kernel/debug/tracing/current_tracer
echo ':mod:lustre' > /sys/kernel/debug/tracing/set_ftrace_filter
echo 1 > /sys/kernel/debug/tracing/tracing_on
```
设置 `function` 追踪器，过滤 `lustre` 模块，然后开启追踪。

# q
Lustre调试日志中的XID和TID分别代表什么含义？
# a
- XID：eXchange ID，关联特定RPC或操作的唯一标识号。
- TID：Transaction ID / Thread ID / Tracepoint ID，具体含义取决于日志实现，通常用于标识事务、线程或追踪点。

