# q
将 `blk_in_osd` 改为 `blk->blk_in_osd` 时，需要评估对哪些变量的影响？
# a
需要评估对 `blk_cmdnum`、`blk_state`、`TQ_state` 的影响。

# q
红黑树操作中，如何通过节点地址获取包含该节点的 `blk_info_t` 结构体地址？
# a
使用以下方法：
1. `blk_rbnode` 是 `blk_info_t` 的成员，二者节点类型一致。
2. 将节点地址强制转换为 `blk_info_t *` 类型。
3. 减去该成员在结构体中的偏移量。
4. 偏移量通过 `offsetof(type, member)` 计算。

# q
本次 block id 修改涉及哪些红黑树核心操作函数？
# a
涉及 `blk_rbtree_search`、`blk_rbtree_remove`、`blk_rbtree_search_next` 等搜索和删除函数。

# q
在 `ocache_query_blks_dirty` 中，为什么 `blk_in_osd` 可以不改为 64 位？
# a
因为 `64*10241/(int)` 的结果不超过 32 位，所以 `blk_in_osd` 不需要改为 64 位。

