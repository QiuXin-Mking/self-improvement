# q
Segment分配/回收流程在执行前需要等待哪些操作完成？冲突点分别是什么？
# a
需要等待以下操作完成：
1. 正在持久化：冲突点在于两者都会更新Segment的block_map状态（持久化更新为PERSIST，分配回收更新为Normal）。
2. 正在压缩：冲突点同样在于都会更新block_map状态（压缩更新为Compact，分配回收更新为Normal）。
3. 正在快照拷贝：冲突点是快照拷贝只有第一次会拷贝，后续命令必须等待拷贝完成。
4. 正在快照合并：冲突点是快照合并会更新元数据，必须等待元数据更新完成。
此外，若元数据不在内存，需从持久化位置还原；若元数据状态为压缩，需重写MDLOG并将状态更新为Normal后才能继续。

# q
Segment压缩流程必须等待哪些操作完成？为什么？
# a
必须等待：
1. 正在回收：回收会先写MDLOG后更新元数据，若压缩在不等待的情况下执行，会导致压缩区缺少回收的MDLOG。
2. 正在持久化：都会更新block_map状态（持久化更新为PERSIST，压缩更新为Compact）。
3. 正在删除：删除会先写MDLOG后删除元数据，若不等删除完成，压缩区会缺少此删除MDLOG。
4. 正在快照合并：快照合并会先写MDLOG后更新元数据，不等其完成会导致压缩区缺少更新元数据的MDLOG。
5. 正在读取/恢复元数据：都会更新Segment的block_map状态。

# q
SEGMENT_SNAPCAPY_HEAD、SEGMENT_PERSIST_HEAD、SEGMENT_COMPACT_HEAD、SEGMENT_SNAPMERGE_HEAD 这些宏的作用是什么？
# a
它们定义了Segment操作冲突等待队列的头部类型，分别对应：
- 快照拷贝等待队列头（SEGMENT_SNAPCAPY_HEAD = 0）
- 持久化等待队列头（SEGMENT_PERSIST_HEAD = 1）
- 压缩等待队列头（SEGMENT_COMPACT_HEAD = 2）
- 快照合并等待队列头（SEGMENT_SNAPMERGE_HEAD = 3）
用于在流程冲突时，将当前操作挂载到对应队列上等待。

