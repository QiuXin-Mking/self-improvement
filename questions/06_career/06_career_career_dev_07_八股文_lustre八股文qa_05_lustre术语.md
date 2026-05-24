# q
什么是 LBUG？
# a
LBUG（Lustre Bug）是 Lustre 软件检测到的致命错误条件，会停止内核线程执行，避免系统状态进一步损坏。该错误会输出到控制台日志，并触发内部调试日志转储（dump）。只有重启系统才能清除 LBUG 状态。

# q
什么是 LFSCK？
# a
LFSCK（Lustre 文件系统检查）是 Lustre 的分布式文件系统一致性检查与修复工具。通常无需运行，仅当文件系统因多磁盘故障等事件损坏，且无法通过日志恢复时，才需要通过 lfsck 进行修复。

# q
什么是 LMV？
# a
LMV（Logical Metadata Volume，逻辑元数据卷）是实现 DNE（分布式命名空间）的客户端抽象层模块。它使客户端无需修改 llite 模块即可与多个 MDT（元数据目标）协同工作：根据名称或目录条带信息将请求转发到正确的 MDT，并将多个应答合并为单一结果返回上层 llite，从而与 Linux VFS 对接并满足 POSIX 语义。

# q
什么是 LNet？
# a
LNet（Lustre Networking）是 Lustre 的消息传递网络协议，能够在多种物理层（如 TCP、InfiniBand）上运行和路由，构成 LNET RPC 的底层通信基础。

# q
什么是 LDLM？
# a
LDLM（Lustre Distributed Lock Manager）是 Lustre 分布式锁管理器，负责协调集群中各个节点对文件系统资源的并发访问。

