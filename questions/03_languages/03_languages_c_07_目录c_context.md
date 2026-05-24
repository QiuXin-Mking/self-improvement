# q
container_of 宏的核心原理是什么？
# a
通过已知结构体成员的地址反推出整个结构体的起始地址。利用 `offsetof` 获取成员在结构体中的偏移量，再将成员指针减去该偏移量，得到结构体指针。典型实现：
```c
#define container_of(ptr, type, member) \
    ((type *)((char *)(ptr) - offsetof(type, member)))
```

# q
`__attribute__((packed))` 的作用和使用场景是什么？
# a
指示编译器取消结构体的对齐优化，按紧凑方式分配内存，成员之间不留填充字节。常用于网络协议、硬件寄存器等需要精确控制内存布局的场景。示例：
```c
struct __attribute__((packed)) data {
    char a;
    int b;
}; // sizeof 可能为 5 而非 8
```

# q
epoll 系列函数包含哪几个核心操作，各自的作用是什么？
# a
- `epoll_create` / `epoll_create1`：创建 epoll 实例，返回文件描述符。
- `epoll_ctl`：向 epoll 实例添加、修改或删除被监控的文件描述符及其事件。
- `epoll_wait`：阻塞等待注册的事件发生，返回就绪事件列表。

# q
如何使用 `typedef` 定义函数指针类型？
# a
语法为 `typedef 返回类型 (*别名)(参数列表);`。例如：
```c
typedef int (*compare_t)(const void *, const void *);
```
之后可直接用 `compare_t` 声明函数指针变量，常用于回调函数、排序比较函数等。

# q
C 语言中的位域是什么？如何声明和使用？
# a
位域允许在结构体中指定成员占用的二进制位数，用于节省内存或访问寄存器中的特定位。声明时在成员名后加冒号和位数：
```c
struct {
    unsigned int flag : 1;
    unsigned int mode : 3;
} bits;
```
可直接赋值和读取，编译器负责位操作。注意位域的对齐和可移植性问题。

