# q
BlueStore 存储引擎在 Ceph 中的核心特点是什么？
# a
BlueStore 是 Ceph 新一代 OSD 存储引擎，直接管理裸盘，使用 RocksDB + BlueFS + WAL，替代 FileStore 以减少写放大。

# q
CRUSH Rule 在 Ceph 中的作用及如何创建？
# a
CRUSH Rule 是 Ceph 的数据分布策略，通过 `ceph osd crush rule create-replicated` 创建，用于控制数据副本在不同 host/rack 级别的物理放置。

# q
Ceph 集群中 PG 数量的推荐计算公式是什么？
# a
PG 总数 = OSD 数量 × 100 ÷ 副本数。单个 pool 的 PG 数不宜过大，以免触发 `mon_max_pg_per_osd` 限制。

# q
Lustre 并行文件系统的核心组件及其功能是什么？
# a
MGS (Management Server) 管理集群配置；MDT (Metadata Target) 存储元数据；OST (Object Storage Target) 存储数据块。客户端通过 MGC/MDC/OSC 连接对应服务，各组件通过 ptlrpc 通信，底层网络由 lnet 抽象支持 TCP/RDMA。

# q
对象存储（RGW）中配额管理的正确步骤是什么？
# a
配额生效需要两步：先用 `radosgw-admin quota set --quota-scope=bucket --max-size=...` 设置限制值，再用 `quota enable` 启用配额。仅 set 不 enable 配额不会实际生效。

