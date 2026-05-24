# q
互斥锁的核心原理是什么？
# a
互斥锁保证多线程同一时间只有一个线程能访问共享数据。线程访问共享变量前先加锁，操作完后解锁；其他线程发现变量被锁住则等待（阻塞）直到锁释放，从而让对共享变量的操作串行化，避免数据不一致。

# q
pthread 互斥锁有哪两种初始化方式？
# a
1. 动态初始化：调用 `pthread_mutex_init(&mtx, NULL)`，属性参数为 NULL 则使用默认属性。需要配套调用 `pthread_mutex_destroy` 销毁。
2. 静态初始化：使用宏 `pthread_mutex_t mtx = PTHREAD_MUTEX_INITIALIZER;`，无需显式销毁。

# q
pthread_mutex_lock 和 pthread_mutex_trylock 的区别是什么？
# a
- `pthread_mutex_lock`：阻塞调用。如果锁正被占用，调用线程进入排队队列并阻塞，直到获取锁后才返回。
- `pthread_mutex_trylock`：非阻塞尝试。如果锁正被占用，立即返回错误码 `EBUSY`，不会阻塞。

# q
释放互斥锁使用的函数是什么？
# a
使用 `pthread_mutex_unlock(&mtx);` 释放互斥锁，用完后必须释放，否则其他等待线程无法获取锁。

