# q
Lustre 是什么？
# a
Lustre 是一个由 DDN 开发和维护的开源分布式并行文件系统，兼容 POSIX，支持 mmap 文件 I/O，广泛部署于超算、石油、金融等领域。后端文件系统支持 ldiskfs（ext4 修改版）和 ZFS，并可使用 RDMA 技术加速通信。

# q
Lustre 中的 MDS 和 OSS 分别负责什么？
# a
MDS（元数据服务器）管理文件系统命名空间，提供文件名、目录、权限、文件布局等元数据服务。OSS（对象存储服务器）存储实际的文件数据对象，向客户端提供数据访问，通常配置多个 OSS 来增加容量和带宽。

# q
MDT 和 OST 的区别是什么？
# a
MDT（元数据目标）是存储元数据的块设备，由 MDS 挂载，至少需要一个 MDT 提供文件系统根。OST（对象存储目标）是存储用户文件数据的块设备，由 OSS 挂载，Lustre 系统总容量为所有 OST 容量之和。同一时刻每个 MDT/OST 只能被一个服务器挂载。

# q
LOV 的作用是什么？
# a
LOV（Logical Object Volume）是位于客户端和 OSS/OST 之间的抽象层，负责管理文件数据如何映射到多个对象，并决定数据在 OST 上的分布方式。客户端通过与 MDS 通信获取 LOV 信息，再直接与 OSS 交互读写数据。

# q
如何使用 lnetctl 管理 Lustre 网络接口？
# a
- 查看所有接口详情：`lnetctl net show --verbose`
- 添加 tcp 网络接口：`lnetctl net add --net tcp --if eth0 --peer-timeout 180 --peer-credits 8`
- 删除 tcp 网络接口：`lnetctl net del --net tcp`
- 查看指定类型网络：`lnetctl net show --net tcp --verbose`

