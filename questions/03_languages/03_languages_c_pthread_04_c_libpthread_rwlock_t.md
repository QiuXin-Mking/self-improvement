# q
pthread_rwlock_t 是什么，它包含哪些锁类型？
# a
pthread_rwlock_t 是 POSIX 线程库中的读写锁类型。它分为“读锁”（阻止写内存）和“写锁”（阻止读、写内存）。读写不能同时进行，且写锁优先级高于读锁。

# q
读写锁的写优先策略如何体现？
# a
写锁优先级高于读锁，当有写锁等待时，即使有新的读锁请求也会被阻塞，直到写锁完成。例如：线程 A 持有读锁，线程 B 请求写锁，线程 C 请求读锁 → B、C 均阻塞；A 解锁后 B 获得写锁，C 继续阻塞；B 解锁后 C 获得读锁。

# q
如何使用 pthread 读写锁（pthread_rwlock_t）？
# a
动态创建并初始化读写锁的典型步骤：
```c
// 分配读写锁结构
pthread_rwlock_t *rwlock = (pthread_rwlock_t*)malloc(sizeof(pthread_rwlock_t));
// 初始化读写锁
pthread_rwlock_init(rwlock, NULL);
// 加写锁
pthread_rwlock_wrlock(rwlock);
// ... 临界区操作 ...
// 解写锁
pthread_rwlock_unlock(rwlock);
```
加读锁使用 `pthread_rwlock_rdlock`，解锁均使用 `pthread_rwlock_unlock`。

