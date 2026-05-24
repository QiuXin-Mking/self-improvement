# q
什么是Reactor模式？其核心原理是什么？
# a
Reactor模式（反应堆模式）是一种事件驱动的高性能服务器开发模式，核心思想是将非阻塞IO与IO复用相结合。它不再主动等待IO就绪，而是通过epoll等IO复用技术监听多个套接字的可读/可写事件，当事件发生时调用对应的回调函数。其典型循环如下：
```cpp
while(!stop) {
    int time_out = Max(1000, getNextTimerCallback());
    int rt = epoll_wait(epfd, fds, ...., time_out);
    if (rt > 0) {
        // 处理发生IO事件的fd，调用其回调函数
    }
}
```

# q
单线程阻塞模型和多线程阻塞模型分别有什么缺点？
# a
单线程阻塞模型：一次只能处理一个客户端连接，完全不支持并发；主线程阻塞在`accept`上，处理完一个连接后才能处理下一个。
多线程阻塞模型：系统最大线程数有限，大量连接时无法无限创建线程；线程频繁创建和切换极度浪费系统资源。此外，两种模型都存在阻塞问题：`accept`、`read`、`write`等慢系统调用会阻塞线程，导致线程资源浪费；同时，一个线程上串行处理多个客户端连接时，若前面的`socket`阻塞，后面的`socket`将永远得不到处理（即使有事件就绪）。

# q
为什么需要使用IO多路复用（如epoll）？
# a
IO多路复用用于解决阻塞问题。它允许单个线程同时监听多个套接字的可读/可写事件，当某个套接字上产生就绪事件时再通知程序进行处理，此时调用`read`/`write`/`accept`等函数不会阻塞。使用epoll只需三步：`epoll_create`创建epoll句柄，`epoll_ctl`注册需要关心的事件（如`listenfd`的可读事件），`epoll_wait`等待事件发生并返回就绪的fd列表。这样就避免了线程被单个阻塞IO耗费，并能高效管理大量并发连接。

# q
主从Reactor模型与单Reactor模型的主要区别是什么？
# a
单Reactor模型：只有一个Reactor线程，既负责监听连接事件（`listenfd`可读），又负责处理已连接套接字（`clientfd`）的IO事件，业务处理也在同一线程或交给线程池，但连接建立与IO事件仍集中在一个Reactor上。
主从Reactor模型：包含一个`mainReactor`和多个`subReactor`。`mainReactor`只负责监听`listenfd`，接收新连接后将`clientfd`的读写事件注册到某个`subReactor`的epoll上；`subReactor`各自运行在独立线程中，专门处理已连接套接字的IO事件和业务逻辑。该模型将连接建立与IO处理解耦，进一步提升了并发能力和可扩展性。

