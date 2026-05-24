# q
什么是MapReduce？MapReduce的核心概念是什么？
# a
MapReduce（映射-归约）是Google提出的一种分布式计算编程模型和框架，用于处理大规模数据集的并行计算。核心概念包括：Map（映射）将输入数据拆分成多个独立的子任务，并行处理生成中间 key-value 对；Reduce（归约）对 Map 阶段产生的中间结果进行聚合、合并、计算，生成最终结果。设计思想为“分而治之”、数据本地性、容错性、可扩展性。执行流程分为 Map、Shuffle、Reduce 三个阶段，适用于日志分析、文本处理、ETL、图计算等大规模数据处理场景。

# q
什么是ReduceTask？它的作用是什么？
# a
ReduceTask 是 MapReduce 框架中 Reduce 阶段的执行单元，负责对 Map 阶段输出的中间结果进行聚合、合并与计算，生成最终结果。其核心作用包括：接收并合并来自不同 Map Task 的数据、按 key 排序、执行用户定义的 reduce 函数，以及将计算结果写入 HDFS 或其它存储系统。ReduceTask 数量可由 `job.setNumReduceTasks(N)` 设置，合理值通常为集群 Reduce 槽数的 0.95–1.75 倍。特殊情况下，若 ReduceTask 设为 0 则无 Reduce 阶段，Map 直接输出结果。

# q
Map Task 的并行度是如何决定的？
# a
Map Task 的并行度由切片（Split）个数决定，每个切片启动一个 Map Task。切片个数根据输入文件大小和切片大小计算，切片大小通常等于 HDFS 的块大小（默认 128MB）。计算公式为 `切片个数 = 向上取整(文件大小 / 切片大小)`。例如，300MB 文件在默认配置下产生 3 个切片，对应 3 个 Map Task。

# q
ReduceTask 的并行度为什么会影响 Job 的执行效率？
# a
ReduceTask 的并行度显著影响 Job 执行效率：设置过少会导致数据倾斜、负载不均，成为性能瓶颈；设置过多会造成输出小文件碎片化、启动和管理开销增大，以及网络传输增加。合理的数量能在并行度与资源开销之间取得平衡，经验值为集群 Reduce 槽数的 0.95–1.75 倍，或根据数据量动态计算，避免数据倾斜和资源浪费。

# q
ReduceTask 数量为 0 代表什么？此时输出文件个数是多少？
# a
ReduceTask 数量为 0 表示没有 Reduce 阶段，只执行 Map 阶段，Map 处理完成后直接将结果写入 HDFS。此时输出文件个数等于 Map Task 个数，每个 Map Task 生成一个以 `part-m-` 开头的输出文件。这种配置适用于只需数据过滤、转换等无需聚合的场景，能够跳过 Shuffle 和 Reduce 过程，提升执行效率。

