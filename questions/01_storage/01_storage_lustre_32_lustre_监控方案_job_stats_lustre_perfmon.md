# q
Lustre的job_stats是什么？
# a
job_stats是Lustre元数据目标(MDT)上按作业(job)维度统计元数据操作性能的功能，通过`lctl get_param mdt.*.job_stats`获取。它可以展示每个进程(如bash、cp、ls)发起的元数据操作类型、次数、耗时分布和I/O字节数，用于性能监控和问题定位。

# q
如何获取Lustre MDT的job_stats？
# a
使用命令`lctl get_param mdt.*.job_stats`即可获取。示例输出会列出每个MDT(如`mdt.nas_test-MDT0000`)下的`job_stats`项，每一项代表一个已结束作业的完整操作统计。

# q
job_stats中每个操作(如open)的统计字段有哪些，分别代表什么？
# a
每个操作会记录以下字段：
- `samples`: 该操作的采样次数(即请求次数)
- `unit`: 耗时单位，通常为`usecs`(微秒)或`bytes`(字节)
- `min`: 最小耗时
- `max`: 最大耗时
- `sum`: 总耗时累加和
- `sumsq`: 耗时平方和(用于计算方差)
对于`read_bytes`和`write_bytes`，unit为`bytes`，还包含`hist`(直方图)字段。

