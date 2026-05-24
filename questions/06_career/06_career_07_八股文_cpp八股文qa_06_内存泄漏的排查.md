# q
使用 Valgrind 排查内存泄漏的完整命令是什么？
# a
```sh
valgrind --tool=memcheck --leak-check=full ./your_program
```

# q
在 GDB 中如何动态输出当前进程的内存分配统计信息？
# a
```sh
# 启动 GDB 并附加到进程
gdb -p <pid>

# 执行 malloc_stats()
(gdb) call malloc_stats()

# 执行 malloc_info() 并输出到文件
(gdb) call malloc_info(0, fopen("temp.xml", "w"))
```

# q
如果不使用外部工具，如何在代码层面记录内存分配与释放以便排查泄漏？
# a
将 `malloc` 和 `free` 映射到调试版本，把每次申请内存和释放内存的详细信息（如地址、大小、调用栈等）全部写入一个文件，然后通过分析该文件查找未配对释放的内存。

