# q
大规模分布式存储研发的核心工作技能主要包括哪些维度？
# a
- 数据结构与算法：B/B+ Tree、LSM Tree、一致性哈希、分布式拓扑、并发算法等
- 分布式系统原理：CAP/BASE理论、共识协议（Paxos/Raft/ZAB）、复制机制、分区分割、故障恢复、分布式事务与2PC/3PC等
- 存储介质与IO：SSD/NVMe/HDD原理、顺序/随机写优化、写放大、GC、Cache/Buffer管理、冷热分层、RocksDB/LevelDB等存储引擎
- 网络与高性能编程：epoll/kqueue、多线程/协程、Zero-Copy、RDMA、低延迟优化
- 高可靠/高可用设计：容错、分布式锁/租约、幂等、自动扩缩容
- 工程实现：精通C++/Rust/Go、性能分析工具、阅读Ceph/TiKV/etcd等开源代码
- 线上实践：大规模部署、监控、故障注入与排查
- 理论转化：将Google BigTable/Spanner/GFS/Raft等论文转化为产品

# q
分布式系统中CAP理论与BASE理论分别是什么？
# a
CAP：一致性（Consistency）、可用性（Availability）、分区容忍性（Partition Tolerance）三者不可兼得，系统需做取舍。BASE：基本可用（Basically Available）、软状态（Soft state）、最终一致性（Eventual consistency），是CAP的实际权衡方案。

# q
分布式存储领域常用的共识协议有哪些？
# a
Paxos、Raft、ZAB（ZooKeeper Atomic Broadcast）等，用于实现选主、日志复制和一致性。

# q
在存储介质和IO层面需要掌握哪些关键知识？
# a
- 介质特性：SSD、HDD、NVMe的IO模式、顺序写与随机写性能差异
- 优化方向：减少写放大、优化垃圾回收（GC）、利用缓存/缓冲区
- 存储分层：数据冷热分层存储
- 典型引擎：RocksDB、LevelDB 等 LSM 结构引擎的实现原理

