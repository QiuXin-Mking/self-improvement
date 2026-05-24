# q
ARRAY_SIZE宏的作用是什么，它是如何实现的？
# a
```c
#ifndef ARRAY_SIZE
#define ARRAY_SIZE(x) (sizeof((x)) / sizeof((x)[0]))
#endif
```
该宏用于计算数组的元素个数，通过整个数组占用的字节数除以单个元素的字节数得出。注意该宏不能用于指针，只能用于编译期能确定大小的静态数组。

# q
container_of宏的核心原理是什么？
# a
```c
#define container_of(ptr, type, member) ({ \
        const typeof( ((type *)0)->member ) *__mptr = (ptr); \
        (type *)( (char *)__mptr - offsetof(type,member) );})
```
它通过已知的结构体成员指针 `ptr`，反推出包含该成员的结构体实例的起始地址。原理是：先创建一个指向 `member` 的指针 `__mptr`，然后将其地址减去 `member` 在结构体类型中的偏移量（`offsetof`），得到结构体指针。这是 Linux 内核中实现泛型数据结构（如链表）的关键宏。

# q
MALLOC、CALLOC 和 FREE 这些宏的设计意图是什么？
# a
这些宏提供类型安全的内存分配与释放：
```c
#define MALLOC(v) \
    v = (typeof(v))malloc(sizeof(typeof(*v)))
#define CALLOC(v, n) \
    v = (typeof(v))calloc(n, sizeof(typeof(*v)))
#define FREE(v) \
    if (v) { free(v); v = NULL; }
```
它们利用 `typeof` 自动推断变量类型，避免手动指定大小或强制转型，减少因类型不匹配导致的内存错误。`FREE` 还会在释放后将指针置为 `NULL`，防止野指针。

# q
foreach 与 foreach_index 系列宏的主要区别是什么？
# a
- `foreach(i, e, array)`：正向遍历数组，每次循环将 `e` 设置为当前元素的指针（`&(array)[i]`）。
- `foreach_index(i, array)`：正向遍历数组，仅提供索引 `i`，不提供元素指针。
类似地还有反向版本 `foreach_reverse` 和 `foreach_index_reverse`，区别在于遍历方向从大到小。`foreach` 系列适合需要直接操作元素地址的场景，`foreach_index` 适合仅依赖索引的操作。

