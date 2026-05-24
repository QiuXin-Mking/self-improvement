# q
OSD_TOO_MANY_REPAIRS 健康警告的触发条件是什么？默认阈值是多少？
# a
当某个 OSD 的 `num_shards_repaired` 超过阈值时触发，默认阈值由配置项 `mon_osd_warn_num_repaired` 控制，值为 10。配置定义如下：
```c
Option("mon_osd_warn_num_repaired", Option::TYPE_UINT, Option::LEVEL_ADVANCED)
.set_default(10)
.add_service("mon")
.set_description("issue OSD_TOO_MANY_REPAIRS health warning if an OSD has more than this many read repairs"),
```

# q
`num_shards_repaired` 计数器在代码中的哪里递增？它代表什么含义？
# a
计数器在 `OSDService::inc_osd_stat_repaired()` 中递增，该函数在 `ReplicatedBackend::handle_push()` 和 `handle_pull_response()` 中被调用，仅在当前 PG 处于 repair 状态时（`pg_is_repair()` 为真）执行。它统计的是该 OSD 上发生的 **读修复（Read Repair）** 次数。关键代码片段：
```c
// src/osd/OSD.cc
void OSDService::inc_osd_stat_repaired() {
  std::lock_guard l(stat_lock);
  osd_stat.num_shards_repaired++;
}
```
```c
// src/osd/ReplicatedBackend.cc
if (get_parent()->pg_is_repair()) {
    pi.stat.num_objects_repaired++;
    get_parent()->inc_osd_stat_repaired();
}
```

# q
Read Repair 与 Scrub Repair 的主要区别是什么？
# a
- **Read Repair（读修复）**：客户端读取时，若从某个副本读取失败（如 EIO 或校验错误），会立即从其他副本读取并修复损坏的副本，属于实时修复。
- **Scrub Repair（清洗修复）**：由定期后台 scrubbing 检查发现数据不一致后触发的修复，属于异步批量修复。

# q
导致 OSD_TOO_MANY_REPAIRS 警告的常见硬件或数据层面根因有哪些？
# a
常见根因包括：
- 磁盘故障或坏块
- 静默数据损坏（bit rot 等）
- 网络问题导致临时读取失败
- 硬件问题（内存、控制器等）

