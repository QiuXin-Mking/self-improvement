# q
当 SSD 故障导致无法修复时，暂停前端 IO 的完整流程是什么？
# a
流程从 `ocachemgt_process_io_pause` 开始，ocache mgt 构造 `OMT_IO_PAUSE` 命令，通过 `ocache_mgt_send_cmd_T_osd` 发送到每个 osd mgt；osd mgt 收到后在 `osd_process_io_pause` 中将 `osd->io_pause` 设置为 `true`，并回复响应；ocache mgt 在 `ocachemgt_process_io_mgt_rsp` 中统计所有 osd 的完成计数（原子变量 `ctx->osd_cnt`），当计数归零后释放上下文并返回最终结果。

# q
在修复过程中，如何判断元数据是否脏（dirty）？
# a
ocache mgt 向 osd mgt 发送 `OMT_REPAIR_SEG_META` 请求，osd mgt 在 `osd_process_repair_segs_meta_dirty` 中通过 `ocache_make_dirty_meta_blk_request` 向 ocache io 发送 `OCACHE_META_BLKS_QUERY` 询问指定 sid 的脏状态；ocache io 在 `ocache_query_meta_blks_dirty` 中查询后通过 `osd_process_query_dirty_meta_seg_rsp` 返回结果；osd mgt 根据结果标记该 segment 是否脏，并继续处理下一个 sid（`sid+1`），直到全部查询完毕在 `osd_process_repair_segs_meta_done` 中完成。

# q
ocache 相关服务组件 ocachemgr、ocachemgt 和 ocache_iosvc 各自承担什么角色？
# a
- `ocachemgr`：负责监控 SSD 状态，发现掉线后将消息交由 ocachemgt 处理。
- `ocachemgt`：统筹管理线程，管理 ocache 的 zdev，一个节点只有一个 ocachemgt；负责协调修复流程（如暂停 IO、开始修复、切换 mdlog 等）。
- `ocache_iosvc`：每个 osd 对应一个 ocache_iosvc 实例（12 个 osd 有 12 个 ocache_iosvc），负责处理具体的缓存 IO 查询和修复操作。

# q
用户数据脏时，如何上报并触发修复流程？
# a
osd mgt 通过 `osd_process_ocache_req` 收到 `OMT_REPAIR_SEG` 后，在 `osd_process_repair_segs` 中调用 `segmgr_ocache_repair_segs_async`，该函数发送 `SEGMGR_MSG_OSD_PART_REPAIRE` 消息给 segment manager，附带需要局部修复的 segment 列表；go 侧接收 `segMsgOsdPartRepair` 类型并执行 `segOsdPartRepairExec` 进行修复。

# q
ocache mgt 向 osd mgt 发送命令时，主要用到哪些命令类型及函数调用方式？
# a
主要命令类型及调用方式：
- `OMT_REPAIR_SEG`：修复用户数据，调用 `ocache_mgt_send_cmd_T_osd(..., OMT_REPAIR_SEG, ..., ocachemgt_repair_segs_rsp)`
- `OMT_IO_PAUSE`：暂停 IO，调用 `ocache_mgt_send_cmd_T_osd(..., OMT_IO_PAUSE, ..., ocachemgt_process_io_mgt_rsp)`
- `OMT_IO_RESUME`：恢复 IO，回调用同上
- `OMT_OCACHE_REPAIRING`：通知 osd 正在修复，调用 `ocache_mgt_send_cmd_T_osd(..., OMT_OCACHE_REPAIRING, NULL, NULL)`
- `OMT_REPAIR_META_DIRTY`：在元数据或用户数据脏时上报修复并重建位图（注释：`OMT_REPAIR_META_DIRTY, /* 在segment元数据或用户数据脏，上报修复，重建位图*/`）

