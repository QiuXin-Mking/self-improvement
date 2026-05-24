# q
list_splice(&src, &dst) 的作用是什么？
# a
list_splice(&src, &dst) 将 src 链表的全部节点从头部插入到 dst 链表头部后的位置（原始 dst 第一个元素之前）。插入后顺序：HEAD ⇄ [src内容] ⇄ [dst原内容] ⇄ TAIL。src 只是内容被拼接，不会被清空或重置。

# q
list_splice_tail(&src, &dst) 和 list_splice 有什么区别？
# a
list_splice_tail(&src, &dst) 把 src 链表的全部节点从尾部插入到 dst 链表末尾（原始 dst 最后一个节点之后）。插入后顺序：HEAD ⇄ [dst原内容] ⇄ [src内容] ⇄ TAIL。

# q
list_splice_tail_init(&src, &dst) 与 list_splice_tail 有何不同？
# a
list_splice_tail_init(&src, &dst) 与 list_splice_tail 类似，都是把 src 拼接到 dst 尾部，但拼接后 src 会被重新初始化为空链表（src: HEAD ⇄ HEAD）。

# q
如何将一个节点添加到链表末尾？使用哪个接口？
# a
使用 ```list_add_tail(&node, &list)``` 将 node 节点插入到 list 链表的尾部。该函数一次只添加一个节点，不是拼接。

# q
这些链表操作是否会直接修改源链表？
# a
- list_splice、list_splice_tail：只改变 src 的链表指针，src 链表头还在，但内容被搬空。
- list_splice_tail_init：拼接后重置 src，src 变为空链表。
- list_add_tail：不影响原有链表，只插入新节点。

