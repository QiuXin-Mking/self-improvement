# q
修改Ceph pool的PG数量后导致PG down的典型根因是什么
# a
在业务流量存在时扩大PG数量会触发大量PG分裂与数据迁移，引发backfilling和recovery风暴。OSD因处理大量IO和内部操作，线程超时（heartbeat_map显示线程timed out after 30），导致OSD心跳延迟甚至超时，被集群标记为down，继而造成PG down。关键日志示例：
```
heartbeat_map is_healthy 'OSD::osd_op_tp thread 0x...' had timed out after 30
```

# q
如何从Ceph日志定位因PG扩容导致的OSD down问题
# a
检查`/var/log/ceph/`下对应OSD的日志：
1. 搜索`heartbeat_map is_healthy`，查找类似`'OSD::osd_op_tp thread ... had timed out after 30`的记录。
2. 搜索`slow request`，如：`cluster [WRN] slow request osd_op(client.38544.0:276863 ...) initiated ... currently delayed`。
3. 执行`ceph -s`，观察输出中的slow ops告警，例如：`46 slow ops, oldest one blocked for 41 sec, daemons [osd.4,osd.8] have slow ops.`。
4. 查看ping交互日志（`handle_osd_ping`）确认心跳延迟是否导致down。

# q
解决Ceph因改PG数量导致PG down的标准处理流程是什么
# a
1. 立即暂停或减少业务写入流量。
2. 临时调整OSD恢复参数以控制恢复速度（需权衡，增大参数可能加重负载）：
   ```sh
   ceph config set osd osd_max_backfills 5
   ceph config set osd osd_recovery_max_active 5
   ceph config set osd osd_recovery_op_priority 3
   ```
3. 重启受影响的OSD或整个集群服务：
   ```sh
   systemctl restart ceph-osd.target
   systemctl restart ceph-mon.target
   ```
4. 持续监控集群状态：`watch -n 1 "ceph -s"`，直到PG恢复正常且无slow ops。
5. 必要时可删除并重建pool（注意数据丢失风险）：
   ```sh
   ceph osd pool rm qiuxin qiuxin --yes-i-really-really-mean-it
   ceph osd pool create qiuxin 100
   ```
6. 集群稳定后，分步调整PG数量（先设较小值再逐步增大），避免一次性大幅度变更。

# q
在Ceph日志中看到`heartbeat_map is_healthy`线程超时意味着什么
# a
这是OSD内部心跳监控发现op线程池中线程超过30秒未响应，表明OSD进程可能因过重负载（如大量backfill/recovery）而卡顿。持续超时会导致集群将该OSD标记为down。日志示例：
```
2025-09-26T16:08:39.401+0800 7fc91d726700  1 heartbeat_map is_healthy 'OSD::osd_op_tp thread 0x7fc8dbe25700' had timed out after 30
2025-09-26T16:08:39.401+0800 7fc91d726700  1 heartbeat_map is_healthy 'OSD::osd_op_tp thread 0x7fc8dfe2d700' had timed out after 30
```

# q
在Ceph集群存在写入负载时修改PG数量，会直接观察到哪些现象
# a
- `ceph -s`显示多个OSD有slow ops阻塞，如`46 slow ops, oldest one blocked for 41 sec, daemons [osd.4,osd.8] have slow ops.`
- OSD日志出现大量`slow request`延迟警告。
- 心跳超时：`heartbeat_map`记录线程timed out。
- OSD间心跳异常，部分OSD被标记为down，导致PG down。
- 集群性能急剧下降，读写延迟飙升。

