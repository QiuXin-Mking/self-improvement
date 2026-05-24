# q
将Ceph池的副本数从2改为3（或从3改为2）时，`ceph -s` 输出中通常会出现哪些集群健康状态和PG状态变化？
# a
- 集群健康变为 `HEALTH_WARN`，提示 `Degraded data redundancy`，例如 `14845664/51265914 objects degraded (28.958%)`。
- 大量PG处于 `active+undersized+degraded+remapped+backfill_wait` 或 `active+undersized+degraded+remapped+backfilling` 等异常状态。
- 出现 `remapped pgs`（例如 2435 remapped pgs）。
- 可能出现 `slow ops`，阻塞时间较长（如 blocked for 1536 sec）。
- 恢复流量：`recovery: 112 MiB/s, 29 objects/s`。

# q
在副本数变更引发的slow ops中，通过 `ceph health detail` 查看操作事件，经常会出现哪个关键状态？该状态的含义是什么？
# a
关键事件状态为 `waiting for clean to repair`。它表示当前客户端IO操作被阻塞，正在等待对应的PG完成恢复（如backfill）并变为 `clean` 状态后才能处理该IO。

