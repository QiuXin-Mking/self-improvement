# q
在Lustre性能测试的fio命令中，如何配置直接IO和使用libaio异步IO引擎？
# a
使用参数 `--direct=1 --ioengine=libaio`。

# q
fio执行读写混合测试时，如何指定读写比例为70%读、30%写？
# a
设置 `--rw=rw --rwmixread=70`。

# q
在空载与半载性能测试脚本中，fio的numjobs参数值有何不同？
# a
空载测试使用 `--numjobs=16`，半载测试使用 `--numjobs=32`。

# q
使用ansible批量运行fio测试时，如何为每个Lustre客户端节点创建独立的测试目录？
# a
通过 `--directory=/mnt/lustre/$(hostname)` 在命令中插入节点主机名变量，实现目录隔离。

