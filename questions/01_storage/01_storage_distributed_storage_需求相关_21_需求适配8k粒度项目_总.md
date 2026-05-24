# q
如何理解 ocache 块粒度与 OSD 块粒度的适配问题？换算关系是什么？
# a
当 osd 块大小为 1M，而 ocache 块大小为 64K 时，blk_in_osd 与 blk_in_ocache 不一致。一个 1M 的 osd block 对应 32 个 64K 的 cache block。若粒度改为 8K，则 `blk_in_osd != blk_in_ocache`，需要新增 `block_size` 变量进行计算：
```c
if ( blk_in_osd[i] * block_size % ocache_size != 0 )
    blk_in_ocache = blk_in_osd[i] * block_size / ocache_size + 1;
else
    blk_in_ocache = blk_in_osd[i] * block_size / ocache_size;
```
在 `ocache_query_blks_req_t` 结构中需增加 osd 粒度字段。红黑树 `blk_rbtree` 可能不再通用，需要额外的 `blk_rbnode`。

# q
查询 OSD segment 在 ocache 中是否有脏数据的流程涉及哪些关键函数和消息？
# a
1. `osd_query_dirty_seg_send_req`：以 segment 的 block id 向 ocache 发起查询，该层有 block size 处理。
2. 向 ocache 发送 `OCACHE_BLKS_QUERY` 消息，其中包含 `ocache_query_blks_req_t` 结构，指定查询的 `blk_num` 范围。
3. ocache 收到后调用 `ocache_query_blks_dirty`（或内部 `blk_rbtree_search`）检查 `[0, ocache_query->blk_num]` 区间内是否存在脏块。
4. 结果通过 `reply_msg`（带 `ocache_rsp` 数据空间）返回，最终由 `osd_process_query_dirty_seg_rsp` 或 `osd_process_query_dirty_seg_done` 处理。

# q
ocache 的刷盘逻辑涉及哪些关键函数和标志？
# a
- 刷盘入口：`ocache_flush_start`、`ocache_flush_blk`。
- 刷盘前准备：`ocache_flush_rdirty_alloc_context` 分配上下文，然后通过 `wf_mgt` 刷盘结构体进行。
- 脏数据标志：`dirty_btmp`（bitmap），每一位代表一个 sector（0.5KB）。若 cache 块大小为 64K，dirty_btmp 有 16 个 8 位数字（128bit）。当粒度增大时，dirty_btmp 可能需要扩展或改变每个 bit 代表的扇区大小。
- 取消刷盘：`ocache_cancel_read_ssd_msg` 取消对 SSD 的读请求。
- 满块标志：`BLOCK_FULL_BTMP == (-1)` 表示刷盘中的满块。
- 无效刷盘：若没有脏数据，刷盘直接返回错误码。

# q
`OCACHE_READ_MSG` 命令的读取路径是怎样的？
# a
`OCACHE_READ_MSG` 属于 osd 向 ocache 发出的读命令。处理流程：
1. osd 发送读命令后，ocache 进入 `ocache_process_read`。
2. 先在 cdisk（SSD 缓存盘）中查找，`ocache_read_oio_hit_blk` 判断是否命中。
3. 若命中则直接返回数据；未命中则通过 `osd_persist_read_handle` 向 osd 的 HDD 发起读取（即 `IO_READ_PERSIST_OCACHE`），读回后再返回给 osd。
4. 消息中通过 LBA（逻辑块地址）、`begin_index`（空闲扇区号）、`blk_cache_idx` 等定位数据偏移。

