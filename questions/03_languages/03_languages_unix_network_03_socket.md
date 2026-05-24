# q
socket 系统调用的作用及其核心参数是什么？
# a
socket 用于创建一个套接字，返回文件描述符。参数包括：
- domain：协议族，如 `AF_INET`（IPv4）、`AF_INET6`（IPv6）
- type：套接字类型，如 `SOCK_STREAM`（TCP流）、`SOCK_DGRAM`（UDP数据报）
- protocol：协议，通常设为 0 使用默认协议
成功返回套接字 fd，失败返回 -1。

# q
setsockopt 系统调用的作用及其参数 level 和 optname 的常见取值是什么？
# a
setsockopt 用于设置套接字选项，配置套接字行为。参数：
- sockfd：套接字文件描述符
- level：选项协议层，常见 `SOL_SOCKET`（通用选项）、`IPPROTO_TCP`（TCP选项）
- optname：具体选项名
- optval：指向选项值的指针
- optlen：选项值的大小
成功返回 0，失败返回 -1。

# q
SO_REUSEADDR 选项的作用是什么？如何通过 setsockopt 启用它？
# a
SO_REUSEADDR 允许重用本地地址，典型地用于服务器重启时快速绑定到之前使用的端口。通过 setsockopt 设置：
```c
int optval = 1;
setsockopt(sockfd, SOL_SOCKET, SO_REUSEADDR, &optval, sizeof(optval));
```
这样可以在 bind 前开启地址重用。

