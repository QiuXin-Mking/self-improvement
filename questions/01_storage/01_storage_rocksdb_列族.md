# q
RocksDB的列族是什么？
# a
列族是将一个RocksDB数据库在逻辑上分割成多个独立的键值存储空间。每个列族都有自己的内存表、SST文件集合和配置选项，但它们共享同一个WAL日志和同一个后台Compaction进程。

# q
如何利用列族分别优化热数据和冷数据的存储？
# a
创建一个 hot_data 列族，配置较小的 write_buffer_size 和较快的压缩算法（如Snappy），以优化写入和内存占用。创建一个 cold_data 列族，配置较大的 write_buffer_size 和更强的压缩算法（如ZSTD），以最大化节省存储空间。两者在同一个RocksDB实例中管理。

# q
RocksDB默认的列族名称是什么？
# a
default

# q
RocksDB中单列族的配置和整个DB的配置分别通过什么对象设置？
# a
单列族的配置通过 ColumnFamilyOptions 设置，整个数据库的配置通过 DBOptions 设置。

# q
在有列族的情况下，RocksDB的写I/O流程是怎样的？
# a
1. 先写WAL文件（全列族共享）。
2. 写入对应列族的MemTable。如果该列族的MemTable写满，会创建新的MemTable，并将旧的MemTable转化为Immutable MemTable，等待刷盘。

