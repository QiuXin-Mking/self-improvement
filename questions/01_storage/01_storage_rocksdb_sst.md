# q
RocksDB中SST文件的全称是什么？
# a
Sorted String Table（排序字符串表，简称SST）

# q
SST文件是在RocksDB的哪些流程中生成的？
# a
SST文件在 flush 和 compaction 流程中生成：
- flush 流程生成 Level 0 的 SST 文件
- compaction 流程生成 Level 1 及更高层级的 SST 文件

# q
SST文件的核心用途是什么？
# a
用于存储一系列有序的 key-value 对，以支持高效的读写操作。

