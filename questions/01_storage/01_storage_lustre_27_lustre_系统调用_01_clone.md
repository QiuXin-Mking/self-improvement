# q
clone 系统调用是什么？
# a
clone 是 Linux 中的系统调用，用于创建新进程或线程，可以精细控制父子进程间的共享资源，也是底层 fork 和 pthread 的实现基础。

# q
clone 调用中 child_stack=NULL 表示什么？
# a
`child_stack=NULL` 表示没有为子进程指定独立的用户栈，子进程将与父进程共享栈。这通常用于创建基于 fork 的进程，而不是用户级线程（创建用户级线程通常需要指定独立栈）。

# q
示例中 clone 的返回值 4159 代表什么？
# a
clone 返回的是新创建的子进程的进程 ID（PID），示例中为 4159。此时父进程继续执行，并根据该 PID 管理和等待子进程。

# q
示例中 wait4 的输出结果说明了什么？
# a
`wait4` 的输出表明子进程（PID 4159）已经终止，且退出状态码为 0（`WIFEXITED` 为真，`WEXITSTATUS` 为 0），表示正常退出。

