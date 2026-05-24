# q
西南八区Ceph集群中，哪些与PG相关的配置参数可能导致PG状态异常？
# a
关键参数包括：
- `osd_pool_default_pg_autoscale_mode = off`：关闭PG自动伸缩，手动修改PG数量后无法自动调整，易引发PG分布不均或down。
- `osd_min_pg_log_entries = 3000` 和 `osd_target_pg_log_entries_per_osd = 600000`：控制PG日志保留量，过低可能影响恢复时PG状态的正确性。
- `osd_map_message_max = 10`：限制OSD map消息大小，可能延迟PG状态变更的传播，导致部分PG长期处于异常状态。

# q
西南八区的Ceph配置中，PG自动伸缩是如何设置的？这如何与“修改PG数量导致PG down”的问题关联？
# a
配置中明确设置了 `osd_pool_default_pg_autoscale_mode = off`，即全局禁用PG自动伸缩。当运维人员手动调整PG数量时，集群不会主动重新平衡或调整其他池的PG，若修改不当（例如未正确计算或未分步操作），容易导致PG分布不均匀、出现down或incomplete状态的PG，最终引发PG down问题。

