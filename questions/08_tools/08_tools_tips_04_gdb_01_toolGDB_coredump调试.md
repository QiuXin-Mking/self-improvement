# q
默认情况下，系统生成的 coredump 文件存放在什么路径？
# a
`/core`

# q
如何使用 gdb 加载一个 coredump 文件并查看崩溃时的调用栈？
# a
```bash
gdb /opt/macrosan/mdbs/bin/engine/engine core_engine_37276
```
然后使用 `bt`（backtrace）查看调用栈，`frame 1` 切换栈帧，`info threads` 查看所有线程。

# q
在使用 gdb 打印结构体时，如何让输出更易读（格式化缩进）？
# a
在 gdb 中执行：
```
set print pretty on
```
之后 `p` 命令输出结构体内容时会自动换行和缩进。

# q
如何在不终止一个正在运行的进程的情况下为其手动生成 coredump 文件？
# a
使用 `gcore` 命令：
```bash
gcore <pid>
```
例如先通过 `systemctl status engine` 或 `ps` 找到进程 PID，然后执行 `gcore <pid>` 即可在当前目录生成 core 文件。

# q
如何启动 gdb 调试一个带调试信息的可执行文件，并查看和运行源码？
# a
1. 编译时添加 `-g` 选项生成调试信息：`gcc -g hello.c -o hello`
2. 启动 gdb（`-q` 可屏蔽版本信息）：`gdb -q hello`
3. 查看源码：`l` 或 `list <行号>`（显示前后共 10 行代码）
4. 运行到断点：`r`
5. 查看调用回溯：`backtrace` 或 `bt`

