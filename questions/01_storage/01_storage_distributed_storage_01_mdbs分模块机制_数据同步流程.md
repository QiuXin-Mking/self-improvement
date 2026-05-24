# q
数据同步流程中的 segSyncData 和 segSyncSendMsg 分别负责什么？
# a
- `segSyncData`：启动一个 goroutine 执行 `segDataSync`，负责实际的数据同步操作。
- `segSyncSendMsg`：发送 `SMT_DATA_SYNC` 类型的消息，触发 `sio_process_cmd` -> `sio_process_seg_sync`，从而进入缓存盘修复或 OSD 修复流程。

# q
缓存盘修复（`SEG_SYNC_CDISK_REPAIR`）与 OSD 全量修复（`SEG_SYNC_OSD`）的区别是什么？
# a
- **触发条件**：缓存盘修复针对 OSD 在线（segment 状态值 2）时的部分元数据丢失；OSD 修复对应换盘修复，OSD 处于 repairing 状态（segment 状态值 1）。
- **执行流**：
  - 缓存盘修复：`sio_process_seg_sync` → 设置 `sync_task->cdisk_repair` → 发送 `OMT_QUERY_SEG_REPAIR` 查脏，查到则只修复脏 segment。
  - OSD 修复：`sio_process_seg_sync` → `sio_seg_sync_begin` → `SEG_RF` → 通过 `IO_SPEED_READ` 读、`IO_SPEED_WRITE` 写，全量修复 segment。
- **状态码**：`SEG_SYNC_CDISK_REPAIR` 对应缓存盘修复任务类型；`SEG_SYNC_OSD` 对应 OSD 全修复任务类型。

# q
ocache 故障导致元数据丢失时，为什么需要执行 OMT_IO_PAUSE 并释放对应的 segment_map？
# a
- **停 IO 原因**：设置 `osd->io_pause = true` 后，写命令在 `osd_process_write` 中返回 `ERROR_OSD_IO_PAUSE`，阻止新写 IO 下发，避免元数据丢失的 segment 继续产生脏数据或冲突。
- **释放 segment_map 原因**：元数据在 ocache 上且未刷盘的 segment，其元数据已丢失，对应的 `segment_map` 必须被 free 掉，否则修复时无法正确重建。释放后会触发查脏流程（`OMT_QUERY_SEG_REPAIR`），若查不到元数据，则会将任务类型从 `SEG_SYNC_CDISK_REPAIR` 切换为 `SEG_SYNC_OSD`，实现全 segment 修复（修复粒度放大）。
- **命令序列**：
  1. `OMT_IO_PAUSE` → `osd->io_pause = true`
  2. `OCACHE_FLUSH_PAUSE` 停止刷盘
  3. 遍历 segment，识别元数据脏且不在内存中的 segment 并释放其 `segment_map`
  4. 后续通过查脏命令 `OMT_QUERY_SEG_REPAIR` 决定转为 `SEG_SYNC_OSD` 修复

# q
缓存盘修复完成后，在何时以及如何重建 alloc bitmap？
# a
- **时机**：在 OSD 数据状态设置为有效位之前（执行 `OMT_DATASTATE_VALID` 前），遍历 OSD 上所有元数据非脏的 segment 链表，重建分配位图（alloc bitmap）。
- **方法**：根据这些 segment 的 block 映射重建位图，如果某个 block 在新位图中不存在但在原位图中存在，则直接 put（释放）该 block。
- **原因**：修复过程中可能产生新的 block 分配，且原脏数据已被清除，必须基于当前有效 segment 重新生成到位图，确保空间分配一致性。同时，重建期间通过 `osd->io_pause` 阻止新写 IO，避免竞争。

