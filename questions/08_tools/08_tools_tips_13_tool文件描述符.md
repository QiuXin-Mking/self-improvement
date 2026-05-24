# q
如何通过 lsof 命令统计每个进程打开的文件描述符数量并按从高到低排序？
# a
```bash
lsof | awk '{print $1}' | sort | uniq -c | sort -nr
```

# q
如何利用 /proc 文件系统查找文件描述符使用过多的进程及其数量？
# a
```bash
find /proc/*/fd -type f | cut -d'/' -f3 | uniq -c | sort -nr
```
该命令列出每个进程的文件描述符数量并按数量从高到低排序。

# q
如何查看系统中当前监听的 TCP 和 UDP 套接字及其关联进程？
# a
```bash
ss -tuln
```

