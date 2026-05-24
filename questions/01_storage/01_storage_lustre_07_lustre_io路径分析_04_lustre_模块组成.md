# q
Lustre 文件系统的高层组件大类包括哪些？
# a
主要包括 LNET/网络相关（通信、网络驱动）、通用库/基础子系统（libcfs、obdclass）、核心客户端/服务端功能（OSC、MDC、MDS、OST、OSD、lov、lmv、ldlm 等）、安全/日志/管理（ptlrpc_gss、gks、mgs、mgc 等）、后端存储适配（ldiskfs、obdfilter、osd* 等）、以及维护/测试组件（obdecho、lfsck 等）。

# q
客户端挂载 Lustre 文件系统时会加载哪些典型模块？
# a
客户端挂载会加载 libcfs、lnet、ksocklnd、lustre（llite）、osc、mdc 等模块。

# q
LNET、PTLRPC 和 LDLM 在 Lustre 中分别承担什么功能？
# a
LNET 是底层网络抽象子系统，负责通信；PTLRPC 是远程过程调用协议层，处理网络请求与回复；LDLM 是分布式锁管理器，负责锁和一致性控制。

# q
LOV 与 LMV 的核心区别是什么？
# a
LOV（Lustre Object Volume）负责数据的条带化与多 OST 聚合管理；LMV（Logical Metadata Volume）负责元数据在多 MDT 上的聚合管理。

# q
MGS 和 MGC 在 Lustre 集群中的角色是什么？
# a
MGS（Management Server）管理所有 Lustre 节点的配置信息；MGC（Management Client）负责客户端与 MGS 的配置通信。

