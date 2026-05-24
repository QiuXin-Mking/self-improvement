# q
Lustre 系统中 MGS、MDS 和 OSS 分别负责什么功能？
# a
MGS 管理集群配置；MDS 管理命名空间、目录与元数据，元数据存储在 MDT 上，需理解 FID、layout、xattr、ACL、quota 等元数据路径与一致性模型；OSS 负责对象存储，提供 OST 存储设备，实现对象级条带化与回写。

# q
Lustre 客户端内部 VFS/LOV/OSC 三层在 I/O 路径上如何分工，支持哪些 I/O 特性？
# a
VFS 提供标准文件系统接口；LOV 实现逻辑卷与条带化映射；OSC 负责与 OSS 的对象存储 RPC 通信。I/O 路径支持 page cache、DIO、aio、mmap、read-ahead、write-back 以及 lockless I/O。

# q
配置 Lustre 条带化时如何选择合适的 stripe size 和 stripe count？
# a
需匹配后端存储介质（SSD/NVMe/HDD）和网络 MTU，对齐 I/O 以减少开销；同时权衡并发写放大效应与锁争用，根据应用的并发模式和访问特征进行调整以提升吞吐。

# q
Lustre 排障和性能调优主要使用哪些工具，各自的作用是什么？
# a
lctl 用于调试和配置；lfs 管理文件条带化与布局；debugfs 查看内核内部信息；/proc/lnet 监控 LNet 状态；strace/perf 分析系统调用和热点；tcpdump 抓包诊断网络；lustre_rsync 用于数据同步。这些工具配合可定位条带化、网络与磁盘瓶颈。

