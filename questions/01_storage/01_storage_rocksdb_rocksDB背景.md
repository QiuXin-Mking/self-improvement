# q
分布式领域的三驾马车指的是哪三种技术？
# a
GFS（分布式文件存储）、Bigtable（分布式KV存储）、MapReduce（基于分布式文件系统和KV存储的大数据处理框架）

# q
LevelDB的作者与哪个系统的作者是同一批人？
# a
Bigtable（Google的分布式KV存储）的作者

# q
RocksDB是由哪家公司开发的，与LevelDB有什么关系？
# a
RocksDB是Facebook基于Google的LevelDB优化而来的

# q
Redis和MySQL分别使用RocksDB作为存储引擎的产品名称是什么？
# a
Redis的版本叫**pika**，MySQL的版本叫**myrocks**

# q
RocksDB的主要缺点和收益分别是什么？
# a
**收益**：提升写入IO性能。  
**缺点**：写放大和读放大（一次读可能先查memtable，再查immutable memtable，最后读磁盘SST，导致多次IO），空间放大，定期合并操作会消耗大量CPU、磁盘和内存等硬件资源。

