# q
什么是 pthread_mutex_t？它的核心作用是什么？
# a
`pthread_mutex_t` 是 POSIX 线程库中的互斥锁数据类型。它是一种同步机制，用于保护多线程环境中的共享资源，确保同一时刻只有一个线程可以访问或修改资源，从而避免竞态条件（race condition）。

# q
pthread_mutex_t 的基本操作有哪些？对应的函数是什么？
# a
基本操作及对应函数：
- 初始化：`pthread_mutex_init(&my_mutex, NULL);`
- 加锁：`pthread_mutex_lock(&my_mutex);`
- 解锁：`pthread_mutex_unlock(&my_mutex);`
- 销毁：`pthread_mutex_destroy(&my_mutex);`

# q
在使用 pthread_mutex_t 时，加锁和解锁的典型代码结构是什么？
# a
典型结构如下：
```c
pthread_mutex_lock(&my_mutex);
// 访问或修改共享资源的代码
pthread_mutex_unlock(&my_mutex);
```
加锁后执行临界区代码，完成后必须解锁，以保证其他线程能够获取锁。

