# q
bpftrace 如何安装？
# a
在 CentOS/RHEL 系统上使用包管理器安装：
```bash
sudo yum install bpftrace
```

# q
bpftrace 如何查看当前版本？
# a
```bash
bpftrace --version
```

# q
bpftrace 如何通过 tracepoint 追踪特定进程的块设备 I/O 请求？
# a
可以使用 `block:block_rq_issue` tracepoint，并通过 pid 过滤。示例：
```bash
sudo bpftrace -e '
tracepoint:block:block_rq_issue /pid == 55068 || pid == 55070/ {
    printf("BLOCK I/O (PID:%d) dev=%d:%d sector=%llu size=%d\n",
        pid,
        MAJOR(args->dev),
        MINOR(args->dev),
        args->sector,
        args->bytes
    );
}
'
```
该脚本打印出发起 I/O 的进程 PID、设备主次编号、起始扇区和字节数。

