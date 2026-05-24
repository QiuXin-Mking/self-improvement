# q
pthread_create 函数的功能是什么？
# a
pthread_create 用于创建新线程，确定调用该线程函数的入口点，创建后线程立即开始运行相关的线程函数。它是 POSIX 线程库（pthread）的一部分。

# q
pthread_create 的参数有哪些？各自的含义是什么？
# a
```c
int pthread_create(pthread_t *tidp, const pthread_attr_t *attr, void *(*start_rtn)(void *), void *arg);
```
- `tidp`：指向线程标识符的指针，成功后存储新线程的 ID。
- `attr`：设置线程属性，一般为 NULL 使用默认属性。
- `start_rtn`：线程运行函数的起始地址，函数形如 `void *func(void *)`。
- `arg`：传递给线程运行函数的参数，单个万能指针；若需多个参数，将其放入结构体并传递结构体地址。

# q
pthread_create 的返回值有何含义？错误处理应注意什么？
# a
- 返回 0 表示线程创建成功。
- 返回非 0 出错编号表示失败，此时 `*tidp` 内容未定义。
- 错误信息通过返回值返回，不修改全局 `errno`，因此不能使用 `perror()` 打印错误，需直接检查返回值并处理。

# q
编译使用 pthread_create 的程序时需要注意什么？
# a
pthread 不是 Linux 默认库，编译时需要显式链接 POSIX 线程库。通常使用编译选项 `-lpthread` 或 `-pthread` 进行链接。

