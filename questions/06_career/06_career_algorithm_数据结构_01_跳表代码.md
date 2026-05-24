# q
跳表节点的内存分配是如何计算的？
# a
分配大小为 `sizeof(*zn) + level * sizeof(skiplistNode_t *)` 的内存，其中 `zn` 是 `skiplistNode_t *`，`level` 是节点层数，额外空间用于存储 `level` 个后继指针。

# q
跳表创建时如何初始化头节点及其前向指针？
# a
调用 `sl_create_node(SKIPLIST_MAXLEVEL, 0, NULL)` 创建头节点，然后将 `zsl->header` 的所有前向指针 `forward[0]` 至 `forward[SKIPLIST_MAXLEVEL-1]` 全部置为 `NULL`。

# q
跳表中随机层数 `sl_random_level` 的生成算法是什么？
# a
从 `level = 1` 开始，在 `(rand() & 0xFFFF) < (SKIPLIST_P * 0xFFFF)` 时 `level` 自增，最后返回 `level` 与 `SKIPLIST_MAXLEVEL` 的较小值，其中 `SKIPLIST_P` 为概率常数（如 0.25）。

# q
在 `sl_insert` 中，如果插入一个与现有节点分数（score）相同的元素会怎样处理？
# a
如果遍历时发现 `x->forward[i]` 存在且 `score == x->forward[i]->score`，会将 `*need_assign` 设为 `false`，并直接返回该已存在的节点，不再执行插入。

# q
删除节点后，`sl_delete_node` 如何调整跳表的最大层数？
# a
通过 `while (zsl->level > 1 && zsl->header->forward[zsl->level-1] == NULL)` 循环，将 `zsl->level` 递减，直到最高层有实际节点或只剩一层，以避免保留空层。

