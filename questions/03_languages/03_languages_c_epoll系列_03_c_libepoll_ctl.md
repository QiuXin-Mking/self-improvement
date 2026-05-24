# q
epoll_ctl函数的作用是什么，它的参数分别代表什么？
# a
`epoll_ctl` 用于向 `epfd` 引用的 epoll 实例添加、修改或删除要监控的文件描述符 `fd`。其声明为：
```c
int epoll_ctl(int epfd, int op, int fd, struct epoll_event *event);
```
- `epfd`：epoll 实例的文件描述符。
- `op`：操作类型（`EPOLL_CTL_ADD`、`EPOLL_CTL_MOD`、`EPOLL_CTL_DEL`）。
- `fd`：要监控的目标文件描述符。
- `event`：关联的事件配置（`struct epoll_event` 指针）。

# q
`epoll_event` 结构体的定义和 `events` 成员可设置的事件宏有哪些？
# a
```c
typedef union epoll_data {
    void *ptr;
    int fd;
    uint32_t u32;
    uint64_t u64;
} epoll_data_t;

struct epoll_event {
    uint32_t events;    /* Epoll events */
    epoll_data_t data;  /* User data variable */
};
```
`events` 成员可以是以下宏的集合：
- `EPOLLIN`：可读
- `EPOLLOUT`：可写
- `EPOLLPRI`：有紧急数据可读（带外数据）
- `EPOLLERR`：发生错误
- `EPOLLHUP`：对端挂断
- `EPOLLET`：边缘触发模式
- `EPOLLONESHOT`：只监听一次，再次监听需重新加入

# q
`epoll_ctl` 的 `op` 参数有哪些有效值，分别用于什么操作？
# a
- `EPOLL_CTL_ADD`：向 epoll 实例注册 `fd`，并关联 `event` 中的事件。
- `EPOLL_CTL_MOD`：更改 `fd` 已注册的事件设置。
- `EPOLL_CTL_DEL`：从 epoll 实例注销 `fd`，此时 `event` 可忽略（可设为 NULL）。

# q
`epoll_ctl` 常见的返回值和错误码有哪些含义？
# a
成功返回 0，失败返回 -1 并设置 `errno`：
- `EBADF`：`epfd` 或 `fd` 无效
- `EEXIST`：`EPOLL_CTL_ADD` 时 `fd` 已注册
- `EINVAL`：`epfd` 不是 epoll 描述符，或 `fd` 与 `epfd` 相同，或操作不支持
- `ENOENT`：`MOD`/`DEL` 时 `fd` 未注册
- `ENOMEM`：内存不足
- `ENOSPC`：达到 `/proc/sys/fs/epoll/max_user_watches` 限制
- `EPERM`：`fd` 不支持 epoll

