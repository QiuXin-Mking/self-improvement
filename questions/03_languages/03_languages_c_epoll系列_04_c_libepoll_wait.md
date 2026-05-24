# q
epoll_wait 函数的核心作用是什么？
# a
epoll_wait 用于等待 epoll 实例中已注册的文件描述符上的 I/O 事件，它会阻塞调用线程，直到有事件发生、被信号中断或超时，然后将就绪事件从内核复制到用户提供的缓冲区。

# q
events 和 maxevents 参数分别有什么作用与要求？
# a
- `events` 指向一个 `struct epoll_event` 数组，用于接收内核返回的就绪事件。
- `maxevents` 指定最多可以返回的事件数量，必须大于 0，内核会按就绪顺序将最多 `maxevents` 个事件写入 `events` 缓冲区。

# q
timeout 参数的含义是什么？epoll_wait 在哪些条件下会停止阻塞？
# a
`timeout` 参数指定 epoll_wait 阻塞等待的超时时间（毫秒）。  
epoll_wait 在以下任一条件满足时返回：
- 感兴趣的 fd 产生了 I/O 事件；
- 被信号处理函数打断；
- 阻塞时间达到 `timeout` 毫秒。

# q
epoll_pwait 和 epoll_pwait2 相比 epoll_wait 增加了哪些参数？
# a
- `epoll_pwait` 增加了 `const sigset_t *sigmask` 参数，用于原子地设置信号掩码。
- `epoll_pwait2` 在此基础上将 `timeout` 参数的类型由 `int` 改为 `const struct timespec *timeout`，可以提供更精确的超时控制（纳秒级）。

