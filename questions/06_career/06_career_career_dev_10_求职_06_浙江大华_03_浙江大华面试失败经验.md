# q
线程和进程之间在资源共享方面有什么区别？
# a
同一个进程内的线程之间资源共享；不同进程之间的资源不共享。

# q
在 Linux 中如何查看当前栈空间大小限制？
# a
使用命令 `ulimit -s` 查看栈空间大小（通常以 KB 为单位）。

# q
如何临时修改 Linux 中的栈空间大小限制？
# a
使用命令 `ulimit -s <stack_size_limit>`，其中 `<stack_size_limit>` 为期望的栈空间大小（KB）。例如：`ulimit -s 8192`。

# q
如何永久修改 Linux 系统的栈空间大小限制？
# a
编辑 `/etc/security/limits.conf` 文件，添加如下行：
```
* hard stack <stack_size_limit>
* soft stack <stack_size_limit>
```
然后保存文件并重启系统使配置生效。

