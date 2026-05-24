# q
什么是IO复用，其核心解决的问题是什么？
# a
IO复用是一种高效的I/O操作处理方式，通过同时监控多个I/O事件（如读、写、连接等），允许单个线程在多个文件描述符上等待事件发生，从而避免进程在等待某个I/O操作完成时被阻塞。它解决了传统阻塞I/O模型中“一个连接占用一个线程/进程”导致的并发能力低和资源浪费问题，提高系统的并发能力和响应性能。常见机制有 select、poll 和 epoll。

# q
epoll机制的核心原理是什么？
# a
epoll 将需要监控的文件描述符及其事件注册到一个内核事件表中（使用红黑树管理），当某个文件描述符上有 I/O 事件发生时，内核会将该就绪事件放入一个就绪链表，应用程序通过 `epoll_wait()` 直接获取已就绪的事件列表，无需遍历所有被监控的描述符，从而实现高效的事件通知。epoll 支持边缘触发（ET）和水平触发（LT）两种模式。

# q
epoll 相比 select 和 poll 的主要优势有哪些？
# a
1. 高效：使用红黑树管理事件表，事件注册和修改效率高；`epoll_wait` 直接返回就绪事件列表，无需像 select/poll 那样遍历全部描述符。
2. 可扩展：支持大量文件描述符，描述符数量增大时性能不会显著下降，没有 select 那样的 FD_SETSIZE 限制。
3. 避免惊群：通过边缘触发（ET）模式可以避免多个线程/进程被同时唤醒的惊群问题。
4. 懒惰删除：已删除的文件描述符会被缓存，在下次 `epoll_wait` 时才真正移除，减少系统调用开销。

# q
epoll 的 LT（水平触发）和 ET（边缘触发）模式的核心区别是什么？
# a
LT 模式（水平触发）：只要文件描述符上还有未处理的事件，`epoll_wait` 每次调用都会通知该事件，直到数据被完全读取或处理。应用程序若未及时处理，会导致重复通知，可能引起性能问题。
ET 模式（边缘触发）：只在文件描述符的状态从“无事件”变为“有事件”的瞬间通知一次，如果应用程序没有一次性处理完所有数据（例如读时未读到 EAGAIN），后续 `epoll_wait` 不会再通知，直到下次有新事件发生。ET 模式必须配合非阻塞 I/O 使用，需循环读取/写入直到返回 EAGAIN，以避免事件丢失，同时能减少无效唤醒和系统调用，提高性能。

# q
epoll 编程常用的 C 函数有哪些，各自的基本作用是什么？
# a
- `int epoll_create(int size)` / `int epoll_create1(int flags)`：创建一个 epoll 实例，返回 epoll 文件描述符。
- `int epoll_ctl(int epfd, int op, int fd, struct epoll_event *event)`：向 epoll 实例中添加（EPOLL_CTL_ADD）、修改（EPOLL_CTL_MOD）或删除（EPOLL_CTL_DEL）要监控的文件描述符及其事件。
- `int epoll_wait(int epfd, struct epoll_event *events, int maxevents, int timeout)`：等待就绪事件，返回就绪文件描述符个数，并将事件信息存入 events 数组中。
- `int epoll_pwait(...)`：与 `epoll_wait` 类似，但可额外通过 `sigmask` 参数在原子操作中屏蔽指定信号。

