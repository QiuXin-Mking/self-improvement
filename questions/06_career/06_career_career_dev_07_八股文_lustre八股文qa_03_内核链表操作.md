# q
list_splice 和 list_splice_tail 的主要区别是什么？
# a
- list_splice：将 src 链表的所有节点插入到 dst 链表**头部之后**（原 dst 第一个节点之前），插入后顺序为 HEAD ⇄ [src内容] ⇄ [dst原内容] ⇄ TAIL。
- list_splice_tail：将 src 链表的所有节点插入到 dst 链表**尾部**（原 dst 最后一个节点之后），插入后顺序为 HEAD ⇄ [dst原内容] ⇄ [src内容] ⇄ TAIL。

# q
list_splice_tail_init 与 list_splice_tail 在操作 src 链表上有什么关键不同？
# a
list_splice_tail_init 在将 src 内容拼接到 dst 尾部之后，**会将 src 重新初始化为空链表**（HEAD ⇄ HEAD），而 list_splice_tail 不会重置 src，src 仍指向已搬空的链表结构。

# q
这些链表拼接接口通常应用于哪些场景？
# a
- 链表的批量合并（如任务调度、缓存迁移）
- 高效批量插入数据，管理链表组织结构
- 按需调整链表头部或尾部内容（如优先队列）
- 重用链表头节点，避免频繁创建和销毁，提升效率

