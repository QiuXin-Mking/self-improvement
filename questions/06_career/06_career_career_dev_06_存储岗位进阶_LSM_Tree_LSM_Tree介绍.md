# q
LSM Tree是什么，主要用于哪些系统？
# a
LSM Tree（Log-Structured Merge Tree）是为大规模、高并发写入而设计的数据结构，核心技术是“写先缓冲-批量顺序落盘-多级合并”。被广泛应用于 LevelDB、RocksDB、TiKV、HBase、Cassandra 等系统。

# q
LSM Tree的核心写入思想是什么？
# a
LSM Tree 将所有更新操作（插入、删除、修改）都转换为追加写操作，首先缓存在内存（MemTable）中，然后定期批量、顺序地刷写到磁盘，从而充分利用磁盘顺序写性能远高于随机写性能的特性。

# q
LSM Tree如何管理磁盘数据以减少碎片？
# a
磁盘数据按多个层级（Level）分散存放，当某个层级的数据量达到阈值时，会触发合并（merge/compaction），将新老数据顺序化融合，以减小碎片、提升读写性能。

