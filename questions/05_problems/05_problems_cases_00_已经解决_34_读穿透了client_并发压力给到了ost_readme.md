# q
在Lustre环境诊断读穿透client并发压力问题时，如何检查客户端缓存使用量？
# a
使用命令 `cat /sys/kernel/debug/lustre/llite/*/max_cached_mb | grep used_mb` 或 `lctl get_param llite.*.max_cached_mb` 查看 used_mb 字段，该值表示已用缓存量（MB）。若 used_mb 远小于 max_cached_mb，可能表明缓存未命中，导致读穿透。

# q
如何获取Lustre客户端的预读(read ahead)使用量？
# a
执行命令 `cat /sys/kernel/debug/lustre/llite/*/max_cached_mb | grep used_read_ahead_mb`，返回的数值即为已使用的预读缓存量（MB）。

# q
Lustre客户端 max_cached_mb 输出中包含哪些关键指标？
# a
输出示例包括 users、max_cached_mb、used_mb、unused_mb、reclaim_count、max_read_ahead_mb、used_read_ahead_mb。关键指标：used_mb（已用缓存）、unused_mb（空闲缓存）、reclaim_count（回收次数）、used_read_ahead_mb（已用预读量），可用于诊断缓存行为和读穿透问题。

