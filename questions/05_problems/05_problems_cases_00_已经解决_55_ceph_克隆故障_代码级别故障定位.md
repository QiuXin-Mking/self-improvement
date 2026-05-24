# q
如何从Ceph OSD日志中定位PG回填的起始对象和回填区间？
# a
在OSD日志中搜索 `recover_backfill` 函数的输出，会打印 `last_backfill_started` 和 `my backfill interval`。例如：
```
recover_backfill (1) bft=10,16 last_backfill_started 4:48e33b35:::rbd_data.00528c12bdeadf.000000000000149c:head
...
my backfill interval BackfillInfo(4:48e33bca:::rbd_data.6deca8a58b558e.000000000000c315:head-4:48ea44bd:::rbd_data.b20f02ee127a37.0000000000000320:head 511 objects {...})
```
`last_backfill_started` 指示回填的起始对象，`BackfillInfo` 给出当前回填的区间与对象数量。

# q
当Ceph PG处于 `backfilling` 且同时带有 `MUST_REPAIR`、`MUST_DEEP_SCRUB` 等标志时，典型根因是什么？
# a
日志中出现 `active+undersized+degraded+remapped+backfilling+forced_backfill MUST_REPAIR MUST_DEEP_SCRUB MUST_SCRUB` 表明该PG存在数据不一致或损坏。通常是由于OSD故障、异常重启或快照/克隆操作导致对象版本不一致，系统强制要求执行 `repair` 和 `deep_scrub` 来修复。

# q
在 `PrimaryLogPG::recover_backfill` 代码中，`new_backfill` 为真时如何重置回填状态？
# a
当 `new_backfill` 为真时，代码会执行：
```
backfill_info.reset(last_backfill_started);
backfills_in_flight.clear();
pending_backfill_updates.clear();
```
即从 `peer_info.last_backfill` 初始化 `last_backfill_started`，清空进行中的回填请求和待更新的回填信息，相当于重新开始回填。

# q
如何通过日志确定回填目标OSD及它们所需的回填对象范围？
# a
在 `recover_backfill` 函数中会遍历回填目标集合并为每个peer输出：
```
peer osd.10 info 4.112(...) interval 4:48e33b35:::rbd_data... - 4:48ea4a94:::rbd_data... 512 objects
peer osd.16 info 4.112(...) interval 4:48e33b35:::rbd_data... - 4:48ea5580:::rbd_data... 512 objects
```
其中 `interval` 后的起止对象名及 `objects` 数量即为该目标OSD需要回填的对象范围。

