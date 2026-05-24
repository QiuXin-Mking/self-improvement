# q
OCache离线流程的触发条件是什么？
# a
当 `adapter_process_rsp` 和 `AMT_WATCH_SSD` 监控到 SSD 状态变化后，由 `ocachemgr_process_ssd_changed` 触发离线修复流程。

# q
OCache离线修复如何暂停和恢复数据刷新？
# a
在 `ocachemgt_process_repair_begin` 之后通过 `OCACHE_FLUSH_PAUSE` 暂停数据刷新，并执行 `ocache_repair_stop_flush_rsp` 和 `ocachemgt_repair_mdlog_shift_rsp`，待异步操作全部完成后进入修复准备阶段（`ocachemgt_repair_prepare_done`）。

# q
脏段查询修复的主要循环流程是什么？
# a
由 `osd_process_query_dirty_seg` 发起，通过 `OCACHE_BLKS_QUERY` 请求查询 `ocache` 中的脏块，`ocache_query_blks_dirty` 处理并返回响应，循环直至所有 segment 遍历完毕，最终调用 `osd_process_query_dirty_seg_done` 结束查询，随后通过 `segmgr_ocache_repair_segs_async` 和 `segMsgOcacheFaulty` 等执行具体修复任务。

