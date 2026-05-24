# q
SSD故障时ocachemgt的处理流程包括哪些关键步骤和通知？
# a
1. 通知所有OSD的ocache停止刷盘（`OCACHE_FLUSH_PAUSE`）
2. 通知所有OSD ocache上报需要修复的segment列表给修复中心
3. 若处理成功，通知osdM将SSD磁盘状态标记为removed，可再次创建SSD
4. 若处理失败，通知osdM标记为待修复，该节点所有OSD需要全盘修复

# q
如何查询某个segment在ocache中是否有脏数据？涉及哪些关键函数和消息类型？
# a
流程：
- `osd_process_query_dirty_seg` 获取segment所有block id，向ocache发送 `OCACHE_BLKS_QUERY` 消息
- `osd_query_dirty_seg_send_req` 构造 `ocache_query_blks_req_t` 请求，其中需包含 `block_size` 和 `blk_num`
- ocache端通过 `blk_rbtree_search` 检查对应block id是否在脏块红黑树中
- 响应通过 `OCACHE_BLKS_QUERY` 返回脏数据标记，结果由 `osd_process_query_dirty_seg_rsp` 处理

# q
当OSD block size与ocache cache size不一致时，如何将OSD块索引转换为ocache块索引？
# a
公式：  
若 `blk_in_osd[i] * block_size % ocache_size != 0`  
 `blk_in_ocache = blk_in_osd[i] * block_size / ocache_size + 1`  
否则  
 `blk_in_ocache = blk_in_osd[i] * block_size / ocache_size`  

脏位图粒度建议保持一致（如1 bit : 0.5KB），如果cache size变化需要相应扩展 `dirty_btmp` 长度。

# q
SSD故障修复的整体流程中，暂停和恢复ocache IO的关键命令是什么？
# a
暂停阶段：  
- 通知ocache io服务暂停刷盘：`OCACHE_FLUSH_PAUSE`  
- 暂停所有ocache的IO：`OMT_IO_PAUSE`  
- 通知mdlog切换存储位置到OSD：`MDLOG_SHIFT_OSD`  
恢复阶段：  
- 恢复ocache IO：`OMT_IO_RESUME`

