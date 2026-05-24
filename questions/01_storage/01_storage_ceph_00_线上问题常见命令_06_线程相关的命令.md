# q
如何查询指定进程（例如PID为85821）中所有safe_timer线程的内核调用栈？
# a
使用以下命令，它会遍历目标进程的每个轻量级线程（task），筛选出线程名包含 `safe_timer` 的线程ID，并打印其 `/proc/<pid>/task/<tid>/stack` 内容：
```bash
ps -T -p 85821 | grep safe_timer | awk '{print $2}' | while read tid; do echo "Thread $tid:"; cat /proc/85821/task/$tid/stack; echo "-----------------------------"; done
```

