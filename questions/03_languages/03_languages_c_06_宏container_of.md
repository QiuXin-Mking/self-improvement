# q
container_of宏的作用是什么？它的核心原理是什么？
# a
container_of宏用于通过结构体成员的指针，获取整个结构体的起始地址。其核心原理是：先通过 `offsetof(type, member)` 计算出成员在结构体中的偏移量，然后将成员指针 `ptr` 减去该偏移量，从而得到结构体指针。宏定义为：
```c
#define container_of(ptr, type, member) ({ \
        const typeof( ((type *)0)->member ) *__mptr = (ptr); \
        (type *)( (char *)__mptr - offsetof(type,member) );})
```

# q
container_of宏中为什么要使用 `const typeof( ((type *)0)->member ) *__mptr = (ptr);` 这一行？
# a
这一行用于类型安全检查：`typeof( ((type *)0)->member )` 获取成员的类型，声明一个与该成员类型一致的常量指针 `__mptr`，并将传入的 `ptr` 赋值给它。如果 `ptr` 的类型与 `member` 的类型不匹配，编译器会给出警告，从而避免因类型错误导致的运行时问题。同时，使用 `const` 可以防止在计算过程中意外修改指针指向的内容。

# q
rb_entry宏与container_of宏有什么关系？blk_rb_entry函数如何利用它们？
# a
`rb_entry` 宏就是 `container_of` 的一个封装，完全等价于 `container_of(ptr, type, member)`，可能用于为红黑树节点提供更语义化的名称。在 `blk_rb_entry` 函数中，它通过传入红黑树节点指针 `node`、外层结构体类型 `blk_info_t` 以及结构体中红黑树节点的成员名 `blk_rbnode`，调用 `rb_entry`（即 `container_of`）计算出包含该红黑树节点的 `blk_info_t` 结构体指针，从而从树节点反向获得外层管理结构体的地址。

