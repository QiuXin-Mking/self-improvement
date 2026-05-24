# q
OCACHE_READ_MSG 在 ocache 读写流程中的作用是什么？
# a
`OCACHE_READ_MSG` 是 OSD 向 ocache 发起的读命令。处理流程为：首先在 cdisk 中查找，若命中则直接返回；若未命中，则进一步向 OSD 中的 HDD 发起读取。

# q
ocache 支持哪些核心命令类型？请列出并简述。
# a
```c
OCACHE_READ = 1,        /* 读命令：osd向ocache发起，先找cdisk，没有则找osd */
OCACHE_READ_LIFT,       /* 读提升命令：osd向ocache查找，找到后返回 */
OCACHE_WRITE_BACK,      /* 写回命令 */
OCACHE_WRITE_THROUGH,   /* 透写命令：直接写到osd，不经过ocache */
OCACHE_WRITE_AROUND,    /* 写穿命令：ocache进入只读状态，同时写到osd和ocache */
OCACHE_FLUSH,           /* 刷盘命令：cdisk内容写到内存，再写到osd */
OCACHE_SPEED_READ,      /* 读命令（暂时未用） */
OCACHE_SPEED_WRITE,     /* 写命令（暂时未用） */
OCACHE_UNMAP,           /* unmap命令：回收删除的vol，释放ocache内存 */
OCACHE_READ_AHEAD,      /* 预读命令：内部发起，把osd内容预取到cdisk */
```

# q
ocache 中 OSD 块与 cache 块的粒度映射关系是怎样的？
# a
1 个 1MB 的 block（在 OSD 中）对应 32 个 64KB 的 cache（在 ocache 中）。ocache 内部红黑树维护的粒度是 64KB，命中位图也是以 64KB 为粒度进行管理。目前 32 个 cache 的 ID 分布既非完全离散也非完全连续，需作为规格确定。

# q
ocache 处理读命中时，涉及哪些关键函数和步骤？
# a
- 通过 `ocache_read_oio_hit_blk` 判断命中的块。
- 使用 `ocache_rcmd_read_hitdata` 读取命中的数据位图（64KB 粒度）。
- 若需要提升读，可通过 `ocache_readlift_oio_hit_blk` 和 `ocache_readlift_cmd_hit_blk` 处理。
- 整个流程中会使用 `blk_rbtree_search` 在红黑树中查找、`blk_rbtree_remove` 移除节点。

