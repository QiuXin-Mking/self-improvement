# q
ceph-kvstore-tool 是什么？
# a
ceph-kvstore-tool 是 Ceph 提供的离线键值数据库操作工具，可以直接操作 LevelDB、RocksDB 和 BlueStore 的 KV 存储。

# q
ceph-kvstore-tool 支持哪些基本操作？
# a
支持 get、set、rm 等基本键值操作，也提供数据压缩、CRC 校验、数据库修复等高级功能。

# q
使用 ceph-kvstore-tool 时有什么前置要求？
# a
必须离线操作，需要先停止相关 Ceph 服务。

