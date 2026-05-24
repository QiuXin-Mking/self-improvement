# q
在Lustre单客户端性能测试中，如何压榨出极限I/O性能？
# a
使用 fio 配置 `-numjobs=128`（128个并发Job）来最大化客户端并发处理能力，并配合 `-iodepth=128`（高队列深度）与 `-direct=1`（绕过页缓存）实现对存储系统的极限压力。

# q
测试中提供的 fio 读性能测试用例的关键参数是什么？
# a
关键参数：
- `-numjobs=128`，`-iodepth=128`
- `-direct=1`，`-ioengine=libaio`
- `-rw=read`，`-bs=1M`，`-size=1G`
- `-directory=/mnt/lustre/$(hostname)`（每个节点独立目录避免冲突）
- `-runtime=30`（快速测试）或 `--time_based --timeout=120`（长时间稳定测试）

# q
测试节点 174 和 175 的单客户端读带宽分别是多少？
# a
- 节点 174：平均读带宽 `3193MiB/s`（约 3348 MB/s）
- 节点 175：平均读带宽 `3123MiB/s`（约 3274 MB/s）
两台节点均使用 `-numjobs=128 -iodepth=128` 的相同 fio 参数完成测试。

