# q
Ceph RGW出现msgr-worker线程CPU占用高的根本原因是什么？
# a
根因是RGW使用的librados中messenger模块的异步网络通信线程存在处理瓶颈，无法应对当前负载，典型表现为多个`msgr-worker-*`线程持续占满CPU。

# q
如何通过系统工具定位Ceph RGW的msgr-worker线程瓶颈？
# a
执行`top -H`查看线程级CPU使用，关注`COMMAND`为`msgr-worker-*`的线程。若多个此类线程（如msgr-worker-0/1/2）的`%CPU`持续接近或超过100%，且`RES`很高（例如11.9g），则说明messenger线程成为瓶颈。

例：
```
   PID USER      PR  NI    VIRT    RES    SHR S %CPU %MEM     TIME+ COMMAND
 34214 ceph      20   0   53.0g  11.9g  25832 R 55.8  4.8   2986:09 msgr-worker-2
 34212 ceph      20   0   53.0g  11.9g  25832 R 48.2  4.8   3775:57 msgr-worker-1
 34211 ceph      20   0   53.0g  11.9g  25832 R 33.5  4.8   3048:23 msgr-worker-0
```

# q
`top -H`输出中，msgr-worker线程的`RES`和`S`状态分别表示什么？对诊断有何意义？
# a
- `RES`（Resident Set Size）：线程实际占用的物理内存大小（示例中为11.9g），高内存占用与高CPU结合表明线程正在大量处理网络数据。
- `S`（状态）：`R`表示Running（运行中），若多个msgr-worker同时处于`R`状态且%CPU接近饱和，说明异步网络线程持续繁忙，存在处理瓶颈。

