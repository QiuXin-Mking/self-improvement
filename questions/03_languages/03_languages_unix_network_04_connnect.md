# q
connect函数的作用是什么？
# a
connect函数用于将指定的套接字连接到指定的目标地址，通常用于客户端建立与服务器的连接。

# q
connect函数的三个参数分别是什么？
# a
- `sockfd`：套接字文件描述符，由 `socket` 函数创建
- `addr`：指向目标地址结构体的指针，类型为 `struct sockaddr*`
- `addrlen`：目标地址结构体的长度

# q
connect函数调用成功和失败时分别返回什么？
# a
成功返回 0；失败返回 -1 并设置全局变量 `errno` 指示错误类型。

# q
在 connect 示例中，如何填充 `struct sockaddr_in` 以指定目标服务器？
# a
- `sin_family` 设置为 `AF_INET`
- `sin_port` 设置为 `htons(12345)`（主机字节序到网络字节序的端口转换）
- `sin_addr.s_addr` 设置为 `inet_addr("127.0.0.1")`（将点分十进制 IP 转换为网络字节序）
- 调用 connect 时将结构体指针强制转换为 `(struct sockaddr*)`

