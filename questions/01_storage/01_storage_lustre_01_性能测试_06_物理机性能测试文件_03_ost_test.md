# q
obdfilter-survey工具的作用是什么？
# a
obdfilter-survey 是 Lustre 中用于测试 OST 性能的工具，可对单个或多个 OST 执行顺序读写、重写等操作，输出带宽结果（MB/s）。

# q
obdfilter-survey测试命令中的 `nobjhi=2 thrhi=2 size=1024 targets="lustre1-OST0000"` 各参数含义是什么？
# a
- `nobjhi`：最大对象数量（此例为 2）
- `thrhi`：最大线程数量（此例为 2）
- `size`：数据集大小，单位为 MB（1024 即 1 GB）
- `targets`：要测试的 OST 名称，格式可为 `OST名称`、`oss:OST名称` 或 `OST_UUID`

# q
obdfilter-survey 输出中 `ost`、`sz`、`rsz`、`obj`、`thr`、`write/rewrite/read` 各表示什么？
# a
- `ost`：被测试的 OST 编号
- `sz`：数据集总大小（如 1048576K 表示 1 TB）
- `rsz`：单次请求的数据块大小（如 1024K）
- `obj`：并发对象数
- `thr`：并发线程数
- `write`：顺序写入带宽（MB/s），方括号内为最小/最大值
- `rewrite`：重写（覆盖写）带宽（MB/s）
- `read`：顺序读带宽（MB/s）
- `SHORT`：标记表示未达到足够长度，结果可能不完整

# q
该文档对 sgpdd-survey 和 fio 的结论是什么？
# a
sgpdd-survey 不好用，不如 fio 方便，并且在测试中曾发现一个磁盘异常。文档推荐使用 fio 进行磁盘性能测试。

