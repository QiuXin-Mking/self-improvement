# q
`syscall(SYS_gettid)` 获取的线程 ID 与 `pthread_self()` 获取的线程 ID 有何区别？为什么它被用于跨进程通信场景？
# a
`pthread_self()` 返回的是 POSIX 线程库维护的线程 ID，只在当前进程内保证唯一，不同进程中的线程 ID 可能重复。`syscall(SYS_gettid)` 返回的是内核分配的真实线程标识（tid），在整个系统中全局唯一，因此可以安全地作为跨进程通信中线程的唯一标识。glibc 未直接提供 `gettid()` 函数，需要通过 `syscall(SYS_gettid)` 调用获取。

# q
示例代码中的 `request_id_gen` 函数如何通过 `syscall(SYS_gettid)` 保证请求 ID 的唯一性？
# a
`request_id_gen` 使用内核线程 ID (`tid`) 作为 `req_id.pid` 字段的一部分，该 ID 通过 `syscall(SYS_gettid)` 在第一次调用时静态初始化获取。由于内核 `tid` 全局唯一，结合节点 ID (`g_nid`) 和原子递增的序列号 (`atomic_inc_return`)，即可生成在分布式或跨进程场景下全局唯一的请求 ID，解决了不同进程中 POSIX 线程 ID 可能相同导致的冲突问题。

