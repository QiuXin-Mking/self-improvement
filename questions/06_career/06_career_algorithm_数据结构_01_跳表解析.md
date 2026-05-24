# q
跳表结构体 `skiplist_t` 包含哪些成员？
# a
```c
typedef struct skiplist {
    struct skiplistNode *header, *tail;
    uint64_t length;
    int32_t level;
} skiplist_t;
```
包含指向头节点和尾节点的指针、跳表中节点的数量 `length`、当前最大层数 `level`。

# q
`sl_create_node` 函数如何计算节点的内存大小？
# a
```c
skiplistNode_t *zn = malloc(sizeof(*zn) + level * sizeof(skiplistNode_t *));
```
分配空间包含节点本身大小加上 `level` 个前向指针数组的大小，即柔性数组 `forward[]` 所需的额外内存。

# q
跳表创建函数 `sl_create` 初始化了哪些关键字段？
# a
- `level` 初始化为 1
- `length` 初始化为 0
- 创建最大层数 (`SKIPLIST_MAXLEVEL`) 的头节点，所有 `forward[j]` 置为 `NULL`
- `header->backward` 和 `tail` 初始化为 `NULL`

# q
`sl_random_level` 函数如何生成随机层数？上限是多少？
# a
```c
int32_t level = 1;
while ((rand()&0xFFFF) < (SKIPLIST_P * 0xFFFF))
    level += 1;
return (level < SKIPLIST_MAXLEVEL) ? level : SKIPLIST_MAXLEVEL;
```
通过随机数按概率 `SKIPLIST_P` 递增层数，最终返回不超过 `SKIPLIST_MAXLEVEL` 的值。

