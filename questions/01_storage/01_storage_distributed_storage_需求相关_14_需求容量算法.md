# q
ds_query_volume_used_capacity 函数在分布式存储中的主要作用是什么？
# a
该函数由 proxy 调用，用于触发一次对卷（volume）已用容量的扫描，是容量查询流程的入口点。

# q
DS 在什么时机被创建？创建过程的关键步骤有哪些？
# a
DS 在为其添加访问路径时由 master 触发创建。创建流程的关键步骤包括：
- 执行 DS 创建操作，产生日志 `ds create successful, ds id:<id>`
- 加载 segment 分布信息（日志 `ds load segment distr info, ds id:<id>, seg nr:<数量>`）
- 完成创建并上报结果和初始容量（日志 `ds create finish when ds process create cmd. ds:<id>, result:0 used cap:0`）

# q
DS 加载的 segment 分布信息中 `seg nr` 的值代表什么含义？
# a
`seg nr` 表示 DS 当前加载的 segment 数量。当 DS 尚未挂载到任何访问路径时该值为 0；挂载到访问路径后，该值反映实际为该卷分配的 segment 数量，例如 `seg nr:4` 表示卷包含 4 个 segment。

# q
proxy 发起的容量扫描与 master 后续的容量查询如何衔接？
# a
proxy 调用 `ds_query_volume_used_capacity` 发出扫描请求；volume 对应的 master 随后调用 `ds_process_query_ds_segmgr_rsp` 处理查询响应，从而完成容量信息的收集。

