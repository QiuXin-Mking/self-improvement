# q
如何从 ceph -s 输出中判断集群存在 PG unfound 问题？
# a
从 `ceph -s` 输出中看到以下关键信息可确认 unfound 问题：
- health: HEALTH\_ERR
- objects unfound 计数：例如 `1/24588320 objects unfound (0.000%)`
- 可能数据损坏警告：`Possible data damage: 1 pg backfill_unfound`
- PG 状态中存在 `active+backfill_unfound+undersized+degraded+remapped`
- 通常伴随 osd down 和大量 degraded/undersized PG

# q
ceph 集群出现 backfill_unfound 状态的 PG 时，典型现象有哪些？
# a
典型现象包括：
- 集群健康状态为 HEALTH\_ERR
- `ceph -s` 显示 unfound objects（如 `1/24588320 objects unfound`）
- 出现 `Possible data damage: 1 pg backfill_unfound` 告警
- 至少一个 PG 处于 `backfill_unfound` 状态，可能结合 `undersized`、`degraded`、`remapped` 等状态
- 对应 OSD 可能处于 down 状态，恢复 I/O 可能受限

