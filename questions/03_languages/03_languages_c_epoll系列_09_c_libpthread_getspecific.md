# q
线程局部存储（TLS/TSD）是什么，它的核心原理是什么？
# a
线程局部存储（TLS，也称线程特有数据 TSD）是一种变量管理机制，每个线程都会维护一份变量的独立副本（copy），该副本长期存在于线程中。对这类变量的读写操作只影响当前线程，不会干扰其他线程，从而在无需使用锁的情况下解决多线程对共享变量的数据竞争问题。

# q
pthread_once函数的作用是什么，如何使用？
# a
`pthread_once` 用于确保某个初始化函数在多线程环境中只被执行一次（常用于单例模式）。函数原型：
```c
#include <pthread.h>
int pthread_once(pthread_once_t *once_control, void (*init)(void));
```
`once_control` 必须指向一个初始化为 `PTHREAD_ONCE_INIT` 的静态变量；`init` 是用户定义的初始化函数，无参数无返回值。无论有多少个线程调用多少次 `pthread_once`，`init` 最多只会执行一次，返回 0 表示成功，正数表示错误。

# q
Linux下操作线程局部数据的主要API有哪些，各自功能是什么？
# a
API 包括：
```c
int pthread_key_create(pthread_key_t *key, void (*destructor)(void *));
int pthread_key_delete(pthread_key_t key);
int pthread_setspecific(pthread_key_t key, const void *value);
void *pthread_getspecific(pthread_key_t key);
```
- `pthread_key_create`：创建一个新键，`destructor` 指向的清理函数在线程终止且与 `key` 关联的 `value` 非 NULL 时自动被调用。
- `pthread_key_delete`：删除键，不调用析构函数，仅释放键供后续重用。
- `pthread_setspecific`：将 `value` 的副本与调用线程和 `key` 关联，`value` 通常指向线程自己分配的内存。
- `pthread_getspecific`：获取当前线程中与 `key` 关联的 `void*` 值，若无关联则返回 NULL。

# q
Linux中线程局部存储机制的内部实现包含哪些关键数据结构？
# a
典型实现包含两层结构：
- 进程级的全局数组 `pthread_keys`，每个元素包含一个标志位（表示该键是否正在使用）和一个析构函数指针（由 `pthread_key_create` 传入的 `destructor`）。`pthread_key_t` 实际上是该数组的索引。
- 每个线程维护一个指针数组，数组的元素与 `pthread_keys` 一一对应，存储由 `pthread_setspecific` 为每个键设置的 `value` 指针。线程终止时，会根据全局数组中的析构函数释放对应数据。

