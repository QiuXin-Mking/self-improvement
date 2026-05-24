# q
什么是线程私有数据（TSD），实现TSD的关键步骤是什么？
# a
线程私有数据 (Thread-specific Data, TSD) 是仅对单个线程有效但可跨多个函数访问的全局变量，同一变量名在不同线程中指向不同的内存空间。实现步骤：
1. 调用 `pthread_key_create` 创建一个 `pthread_key_t` 变量作为键；
2. 在线程内需要存储特殊值时调用 `pthread_setspecific` 将该键与一个 `void*` 值关联；
3. 通过 `pthread_getspecific` 用该键取出关联的指针，从而在同一线程的不同函数间共享私有数据。

# q
pthread_key_create 的参数和用途是什么？
# a
`pthread_key_create` 用于创建线程私有数据的键，原型为：
```c
int pthread_key_create(pthread_key_t *key, void (*destructor)(void*));
```
第一个参数是 `pthread_key_t` 变量的指针；第二个参数是清理函数指针，当线程退出且该键关联的数据非空时会被调用以释放资源，可以设为 `NULL` 使用默认清理。

# q
pthread_setspecific 和 pthread_getspecific 的工作原理及使用注意事项是什么？
# a
- `pthread_setspecific(key, value)` 将 `value` 指针本身（而非指针指向的内容）与 `key` 关联，存入线程私有存储区。
- `pthread_getspecific(key)` 返回该线程中与 `key` 关联的 `void*` 指针，不同线程用同一个 `key` 会得到各自的私有数据，互不影响。
- 注意事项：`pthread_getspecific` 返回 `void*`，必须强制转换为实际数据类型后才能正确使用；每个线程需先调用 `pthread_setspecific` 建立关联，且同一 `key` 在不同线程中指向的内存独立。

