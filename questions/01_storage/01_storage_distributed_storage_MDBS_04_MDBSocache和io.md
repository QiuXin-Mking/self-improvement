# q
ocache_io_t 结构体包含哪些关键组件？
# a
```c
typedef struct ocache_io {
    struct blist *io_bl;               /* ocache io的blist */
    struct list_head ocache_cmd_list;  /* ocache命令队列     */
    msg_t pri_req;                     /* 读写命令对应的msg,该msg是ocache数据流入口收到的 */
    atomic_t cmd_num;                  /* 未回调的ocache命令个数     */
    int32_t io_result;                 /* ocache io处理的返回值 */
} ocache_io_t;
```
包含：IO数据的 `blist` 链表、子命令链表头、关联的原始消息 `pri_req`、原子计数的未回调命令数 `cmd_num`、处理结果 `io_result`。

# q
ocache_alloc_io 函数的作用是什么？
# a
完成消息到 `ocache_io` 结构的转化。它分配并清零 `ocache_io_t` 结构，初始化字段：
- 初始化 `ocache_cmd_list` 链表头；
- 保存 `pri_req` 并调用 `msg_get` 增加消息引用；
- 设置原子变量 `cmd_num` 为 0；
- 如果消息包含 blist，则获取 `io_bl` 并增加引用；
- 递增全局 `g_ocache_ctrl->oio_num` 计数。

# q
ocache_split_cmd 函数的核心拆分原理是什么？
# a
根据原始 IO 的起始 LBA 和扇区数，以块页大小（`g_ocache_ctrl->blk_page_size`）为单位分割：
1. 计算首块偏移 `lba_offset = get_block_head_offset(begin)`；
2. 拆分出的子命令个数 `cmd_num = (lba_offset + size + blk_page_size - 1) / blk_page_size`；
3. 循环创建 `ocache_cmd`，设置每个子命令的起始 LBA、长度（取块尾对齐后剩余大小）和命令类型（读/写回/UNMAP）；
4. 对 `io_bl` 进行处理：若单个 blist 长度与子命令长度相等则直接引用，否则调用 `blist_split` 拆分链表，子命令获得对应数据 blist；
5. 每个子命令加入 `ocache_cmd_list`，并增加 `cmd_num` 计数。失败时回滚，释放已分配的命令。

# q
io_svc 初始化生命周期的关键配置有哪些？
# a
- `io_svc->ocache_data_size` 设为 5GB（`5 * 1024 * 1024 * 1024l`）；
- 分配 `blk_alloc` 内存，大小为 `ocache_data_size / BITMAP_PER_GROUP_SIZE`；
- 初始化四个链表：`blk_fullist`、`blk_lrulist`、`cmd_waitlist`、`wait_cache_list`；
- 初始化第一个 zone 的分配位图 `cache_alloc_bitmap`（长度 `OCACHE_ALLOC_BITMAP_LEN`），清零位图，并将 `free_num` 设为 `BLK_NUM_PER_ZONE`。

