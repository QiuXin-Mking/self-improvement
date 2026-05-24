# q
什么是LSM Tree？它的全称是什么？设计目标是什么？
# a
LSM Tree（Log-Structured Merge Tree）是为大规模、高并发写入而设计的数据结构。其核心技术是“写先缓冲-批量顺序落盘-多级合并”，通过将更新操作转换为追加写，充分利用磁盘顺序写性能远高于随机写性能的特性。

# q
LSM Tree的写入流程是怎样的？数据如何从内存持久化到磁盘？
# a
写入（插入和更新）首先缓存在内存中的MemTable里，然后定期批量、顺序地刷写到磁盘上。磁盘上的数据按多个层级分散存放，当数据量达到阈值时会触发合并（merge/compaction），将新老数据顺序化融合，减小碎片并提升读写性能。

# q
常见使用LSM Tree的存储系统有哪些？
# a
LevelDB、RocksDB、TiKV、HBase、Cassandra 都使用了LSM Tree结构。

