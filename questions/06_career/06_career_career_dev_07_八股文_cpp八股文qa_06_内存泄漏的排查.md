# q
如何使用 valgrind 检测内存泄漏？
# a
```sh
valgrind --tool=memcheck --leak-check=full ./your_program
```

# q
在 GDB 中如何获取进程的内存分配统计信息？
# a
启动 GDB 附加到进程，然后调用相应的库函数：
```sh
gdb -p <pid>
(gdb) call malloc_stats()
(gdb) call malloc_info(0, fopen("temp.xml", "w"))
```
`malloc_stats()` 输出统计到标准错误，`malloc_info()` 将详细信息写入指定文件。

# q
在不使用外部工具时，排查内存泄漏的基础方法有哪些？
# a
1. 将 `malloc` 与 `free` 映射到调试版本，记录每次内存申请和释放的详细信息到日志文件，再离线分析。
2. 在无法直接使用 valgrind 的多服务环境中，可以编写启动命令或通过打桩（stub）返回的方式进行间接排查。

