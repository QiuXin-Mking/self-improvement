# q
Linux C API函数异常时，errno变量的作用是什么？需要包含哪个头文件？
# a
当Linux C API函数发生异常时，一般会将全局变量 `errno` 设置为一个整数值，表示不同的错误原因。可以通过查看该值来推测出错原因，调试程序。使用 `errno` 需要包含 `<errno.h>` 头文件。

# q
1-34号错误号和35-132号错误号分别在Linux内核源码的哪个头文件中定义？
# a
- 1-34号错误号在内核源码的 `include/asm-generic/errno-base.h` 中定义。
- 35-132号错误号在 `include/asm-generic/errno.h` 中定义（该文件还会先 `#include <asm-generic/errno-base.h>`）。

# q
请列举 `errno-base.h` 中定义的几个常见错误码及其含义：EPERM、ENOENT、EACCES、ENOMEM。
# a
- `EPERM` (1): Operation not permitted（操作不允许）
- `ENOENT` (2): No such file or directory（文件或目录不存在）
- `EACCES` (13): Permission denied（权限不足）
- `ENOMEM` (12): Out of memory（内存不足）

# q
`EWOULDBLOCK` 错误码是如何定义的？其对应的宏值是多少？
# a
`EWOULDBLOCK` 被定义为 `EAGAIN`，即宏 `EWOULDBLOCK` 与 `EAGAIN` 的值相同（11），含义是 "Operation would block"。

# q
常见的网络相关错误码有哪些？请列举 `ECONNREFUSED`、`ETIMEDOUT`、`ENETUNREACH` 的含义。
# a
- `ECONNREFUSED` (111): Connection refused（连接被拒绝）
- `ETIMEDOUT` (110): Connection timed out（连接超时）
- `ENETUNREACH` (101): Network is unreachable（网络不可达）

