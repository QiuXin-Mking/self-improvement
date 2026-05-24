# q
如何通过 strace 定位 radosgw-admin 进程挂住问题？
# a
使用 strace 跟踪进程及其所有线程，并设置超时采集，命令如下：
```bash
timeout 5m strace -p <PID> -f -tt -o strace_output.txt
```
- `-p`：指定目标进程 PID
- `-f`：跟踪所有子线程
- `-tt`：打印微秒级时间戳
- `-o`：将输出写入文件，避免干扰终端
- `timeout 5m`：限制采集 5 分钟，防止长时间阻塞

