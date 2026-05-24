# q
Ceph 集群中 Monitor 和 OSD 分别使用了哪种一致性协议？
# a
- **Monitor 元数据层**：使用 Paxos 协议，通过多数派投票保证集群元数据（如 OSD 状态、CRUSH 映射）的强一致性，每个 Monitor 均可接受写请求。  
- **OSD 数据层**：采用类 Raft 的日志复制机制，为数据存储提供高性能的一致性保障。

# q
Ceph 的三副本写 I/O 流程如何保证数据一致性？确认时机是什么？
# a
写入采用**同步复制**，Primary OSD 必须等待所有副本（包含自己）返回 **CEPH_OSD_FLAG_ONDISK**（表示数据已落盘）后，才会向客户端回复 ACK。  
流程：  
1. Client 向 Primary 发送写请求。  
2. Primary 将写操作发送给所有副本，同时本地落盘。  
3. Primary 收集所有副本的 ONDISK 确认（通过 `waiting_for_commit` 集合管理）。  
4. 当 `waiting_for_commit` 清空（所有副本确认），Primary 才触发 `on_commit` 回调，向客户端返回 ACK（包含 ACK + ONDISK 标志）。

# q
LevelDB 和 RocksDB 的核心差异及改进点是什么？
# a
- **LevelDB**：Google 开发的高性能键值嵌入式存储，基于 LSM 树，适合大量写入，采用单线程 compaction。  
- **RocksDB**：由 Facebook 基于 LevelDB 改进，主要差异：  
  - **多线程**：支持多线程压缩和 memtable 写入，提升并发写入能力。  
  - **写放大优化**：显著降低写放大问题，提高高并发写入吞吐。  
  - **硬件适配**：针对 SSD 优化，可处理更大数据量和更高 IOPS，适合大数据与分布式环境。  
  - **社区活跃**：持续增加功能，解决 LevelDB 主动限制写入等实际问题。

