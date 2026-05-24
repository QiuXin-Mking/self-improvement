# q
跳表是什么？它的核心操作平均时间复杂度是多少？常见应用有哪些？
# a
跳表是在有序链表基础上增加多级索引的随机化数据结构，支持快速查找、插入和删除。其搜索、删除、添加的平均时间复杂度均为 O(log n)。常见应用包括 Redis 的 SortedSet 和 LevelDB 的 MemTable。相比平衡树，跳表的实现和维护更简单。

# q
跳表的节点结构（`skiplistNode`）中包含哪些关键字段？
# a
```c
typedef struct skiplistNode {
    void *ele;                // 元素值
    uint64_t score;           // 排序分值
    struct skiplistNode *backward;  // 后退指针
    struct skiplistNode *forward[]; // 柔性数组，前进指针（多级索引）
} skiplistNode_t;
```
每个节点必须维护 `ele` 和 `score`。`backward` 用于反向遍历；`forward` 是数组，长度由节点的层级决定，下标最大值是 `level` 或 `level-1`。

# q
跳表的随机层次生成函数 `sl_random_level` 是如何工作的？层级概率分布是怎样的？
# a
```c
int32_t sl_random_level(void) {
    int32_t level = 1;
    while ((rand()&0xFFFF) < (SKIPLIST_P * 0xFFFF))
        level += 1;
    return (level < SKIPLIST_MAXLEVEL) ? level : SKIPLIST_MAXLEVEL;
}
```
- 默认 `SKIPLIST_MAXLEVEL` 为 32。
- 每次随机数小于阈值（SKIPLIST_P=0.25）时层级+1，所以层数越高概率越低。
- 概率分布：  
  | level | 概率 |
  |:---:|:---:|
  | 1 | 0.25 |
  | 2 | 0.0625 |
  | 3 | 0.015625 |
  | … | 指数递减 |
最高不超过 32 层。

# q
跳表查询函数 `sl_find` 的核心查找逻辑是什么？
# a
```c
x = zsl->header;
for (i = zsl->level-1; i >= 0; i--) {
    while (x->forward[i] && x->forward[i]->score < score)
        x = x->forward[i];
    if (x->forward[i] && score == x->forward[i]->score)
        return x->forward[i];
}
return NULL;
```
从最高层开始，每层向右移动直到节点的 score 小于目标 score，若该层找到相等 score 则直接返回；否则下降一层继续查找，直至最底层都未找到则返回 NULL。

# q
跳表插入操作 `sl_insert` 中，`update` 数组和随机层数 `level` 的作用是什么？如何处理层级提升的情况？
# a
- `update` 数组：在每层记录小于插入 score 的最后一个节点，用于后续调整各层的转发指针。
- 随机生成插入节点的层数 `level`：如果 `level > zsl->level`，则将 `update` 中新增加的层指向表头 `zsl->header`，并更新跳表总层数 `zsl->level = level`。
- 插入节点后，对于 0 到 `level-1` 的每一层，将新节点插入到 `update[i]` 之后，更新转发指针和后退指针。若新节点位于末尾，则更新跳表的 `tail` 指针。

