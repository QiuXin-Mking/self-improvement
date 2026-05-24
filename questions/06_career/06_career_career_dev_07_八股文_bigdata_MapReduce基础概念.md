# q
什么是MapReduce？MapReduce的核心概念和执行流程是什么？
# a
**MapReduce（映射-归约）**：是Google提出的分布式计算编程模型和框架，用于处理大规模数据集的并行计算。

**核心概念：**
- **Map（映射）**：将输入数据拆分成多个独立的子任务，并行处理，生成中间结果（key-value对）
- **Reduce（归约）**：对Map阶段产生的中间结果进行聚合、合并、计算，生成最终结果

**设计思想：**
- **分而治之**：将大问题分解成小问题，并行处理
- **数据本地性**：尽量在数据所在节点进行计算，减少网络传输
- **容错性**：自动处理节点故障，重新执行失败的任务
- **可扩展性**：可以扩展到数千个节点的集群

**执行流程：**
1. **输入阶段**：将输入文件切分成多个切片（Split）
2. **Map阶段**：读取输入数据，执行用户定义的map函数，生成中间结果（key-value对）
3. **Shuffle阶段**：将Map输出按照key进行分区、排序、合并
4. **Reduce阶段**：对相同key的数据执行reduce函数，生成最终结果
5. **输出阶段**：将结果写入HDFS或其他存储系统

**示例（WordCount）：**
输入: "hello world hello"
→ Map阶段: (hello,1), (world,1), (hello,1)
→ Shuffle: hello→[1,1], world→[1]
→ Reduce: (hello,2), (world,1)
→ 输出: hello 2, world 1

# q
MapReduce中Map Task的并行度是如何决定的？
# a
**Map Task并行度由切片（Split）个数决定**，每个切片会启动一个Map Task。

**计算公式：**
```
切片个数 = 向上取整(文件大小 / 切片大小)
Map Task个数 = 切片个数
```

**切片规则：**
- 切片大小通常等于HDFS块大小（默认128MB）
- 切片是逻辑概念，不是物理切分

**示例：**
- 100MB文件，切片大小128MB → 1个切片 → 1个Map Task
- 300MB文件，切片大小128MB → 3个切片 → 3个Map Task

**影响因素：**
- HDFS块大小（block size）
- InputFormat的实现方式
- 文件是否可切分（如压缩文件可能不可切分）

# q
ReduceTask的并行度如何影响Job的执行效率？
# a
**ReduceTask的并行度会显著影响整个Job的执行效率。**

**影响机制：**
1. **ReduceTask太少**：导致数据倾斜，部分节点负载过重，成为性能瓶颈；其他节点资源闲置；执行时间延长。
2. **ReduceTask太多**：导致输出文件碎片化，每个Reduce输出一个文件；启动和管理Task开销增加；网络传输次数增加；调度负担增加。

**合理设置：**
- 经验值：集群可用Reduce槽数的0.95～1.75倍
- 根据数据量动态计算：
```java
long totalInputSize = ...;
long reduceTaskSize = 500 * 1024 * 1024; // 每个Reduce处理500MB
int numReduceTasks = (int)(totalInputSize / reduceTaskSize);
job.setNumReduceTasks(numReduceTasks);
```

**注意事项：**
- 需要在并行度和资源开销之间找到平衡
- 可以使用自定义Partitioner避免数据倾斜

# q
MapReduce中ReduceTask数量为0表示什么？输出文件个数是多少？
# a
**ReduceTask数量为0表示没有Reduce阶段**，只执行Map阶段。

**执行流程：**
- 只执行Map阶段，跳过Shuffle和Reduce
- Map阶段处理完成后，直接将结果写入HDFS

**输出文件个数：**
- 输出文件个数 = Map Task个数
- 每个Map Task输出一个结果文件，格式为 `part-m-XXXXX`

**适用场景：**
- 只需数据过滤、转换，不需要聚合
- 不需要排序
- 简单的ETL操作

**设置方法：**
```java
job.setNumReduceTasks(0);
```

**对比：**
- ReduceTask > 0：有Reduce阶段，输出文件数 = ReduceTask个数，文件名为 `part-r-XXXXX`
- ReduceTask = 0：无Reduce阶段，输出文件数 = Map Task个数，文件名为 `part-m-XXXXX`

