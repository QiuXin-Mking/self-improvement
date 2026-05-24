# q
WAL的全称是什么？
# a
Write Ahead Log（预写日志）。

# q
WAL解决了什么问题？
# a
解决断电或崩溃时内存数据丢失的问题。RocksDB 先将写操作写入 WAL 和 MemTable，然后返回成功。如果断电，MemTable 丢失，可通过回放 WAL 恢复已提交的数据，保证持久性。

# q
RocksDB 的写流程中，什么阶段认为写操作已经成功？
# a
先写 WAL，再写 MemTable，数据写入内存中的 MemTable 后即返回成功。当 MemTable 写满时，会切换成不可修改的 Immutable MemTable，并触发 Flush 到 SST 文件。

# q
RocksDB 中的 Manifest 是什么？
# a
Manifest 文件记录了每一层 SST 文件的元数据及关键的层级概览信息，用于描述 LSM-Tree 的当前状态。

# q
MemTable 最常见的数据结构实现是什么？
# a
跳表（SkipList）。

