# q
如何使用gdb初步定位Ceph OSD进程hang的问题？
# a
使用 `gdb attach <osd-pid>` 进入调试会话，执行 `info thread` 查看所有线程的等待状态。重点关注是否出现异常的锁等待，例如 `__lll_lock_wait`，如以下片段所示：

```
Thread 0x7f00da8f8700 (LWP 97951) "safe_timer" 0x00007f00e365a75d in __lll_lock_wait () from /lib64/libpthread.so.0
```

大量线程处于 `pthread_cond_wait`、`epoll_wait` 或 `poll` 属于正常的等待行为，而出现 `__lll_lock_wait` 则提示可能发生锁争用或死锁。

# q
从gdb的info thread输出中，看到一个线程卡在__lll_lock_wait，这通常意味着什么？
# a
表示该线程正在尝试获取一个互斥锁（pthread mutex），但该锁已被其他线程持有，导致本线程进入内核态等待锁释放。这通常意味着存在锁竞争，如果持锁线程同样在等待某个条件得不到满足，就可能形成死锁，导致进程hang。

# q
Ceph OSD hang时，正常的线程等待状态有哪些？
# a
大部分后台线程正常等待时，会显示为：
- `pthread_cond_wait`：等待条件变量（如 rocksdb 线程、safe_timer、service 线程等）
- `epoll_wait`：等待网络事件（如 msgr-worker 线程）
- `poll`：等待文件描述符事件（如 admin_socket、signal_handler）
- `nanosleep`：主动休眠（如 OpHistorySvc）

这些状态本身不表示异常，需要结合业务逻辑和数量综合判断。

