# q
项目中的chan提供了哪些核心操作接口？
# a
```c
int32_t chan_new(chan_t** ch, uint32_t count, const char* name);   // 创建channel
int32_t chan_send(chan_t* ch, void* e);                            // 发送元素
int32_t chan_recv(chan_t* ch, void** e);                           // 接收元素
fd_t chan_fd(chan_t* ch);                                          // 获取文件描述符
int32_t chan_wait(chan_t* ch);                                     // 等待通道就绪
void chan_free(chan_t* ch);                                        // 释放channel
int32_t chan_count(chan_t* ch);                                    // 获取当前元素计数
```

# q
chan_fd函数的主要用途是什么？
# a
返回chan对应的文件描述符（fd_t），用于将channel集成到事件循环（如epoll）中，实现异步通知或多路复用。

# q
chan_new函数的参数分别表示什么？
# a
```c
int32_t chan_new(chan_t** ch, uint32_t count, const char* name);
```
- `ch`：输出参数，指向新创建的chan指针的指针
- `count`：channel的容量（缓冲区大小）
- `name`：channel的名称（用于调试或标识）
返回值：0成功，非0失败。

