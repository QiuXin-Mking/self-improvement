# q
pthread_create 函数的核心功能是什么？
# a
pthread_create 是 POSIX 线程库提供的创建线程的函数，用于在调用进程中创建一个新线程。该线程从指定的起始函数地址开始运行，可以接收一个 `void*` 类型的参数。创建成功后线程自动执行，无需手动启动。

# q
pthread_create 函数的参数有哪些，含义分别是什么？
# a
```c
int pthread_create(pthread_t *tidp, const pthread_attr_t *attr,
                   void *(*start_rtn)(void *), void *arg);
```
- `tidp`：指向 `pthread_t` 的指针，用于保存新线程的 ID。
- `attr`：指向线程属性的指针，通常传入 `NULL` 使用默认属性。
- `start_rtn`：线程函数的起始地址，即线程要执行的函数指针。
- `arg`：传递给线程函数的参数，类型为 `void*`，多参数时需要打包成结构体，并传递结构体的地址。

# q
pthread_create 的返回值说明了什么？如何判断成功或失败？
# a
- 成功时返回 `0`，同时 `*tidp` 所指向的内存被设置为新线程的 ID。
- 失败时返回非零的错误编号，此时 `*tidp` 的内容未定义。错误信息通过返回值获取，不修改全局 `errno`，因此不能使用 `perror()` 打印错误信息。

# q
编译使用 pthread 库的程序时需要注意什么？
# a
pthread 不是 Linux 默认链接的库，编译时需要显式指定，常用 `-lpthread` 或 `-pthread` 选项。例如：
```bash
gcc -o prog prog.c -lpthread
```

# q
示例代码中为何要使用 `sleep()` 让主线程等待？
# a
如果主线程执行结束，整个进程会立即终止，所有未完成的子线程会被强制结束。使用 `sleep()` 延迟主线程，可以保证子线程有足够时间完成执行。更好的做法是使用 `pthread_join()` 来同步等待线程结束。

