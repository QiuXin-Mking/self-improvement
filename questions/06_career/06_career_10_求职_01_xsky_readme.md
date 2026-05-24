# q
Lustre 客户端发起一次文件 I/O 到 OST 的完整流程是怎样的？
# a
流程顺序如下：
1. **VFS 层入栈**：用户调用 read()/write()，Linux VFS 查找到 Lustre Lite (llite) 提供的 file_operations/address_space_operations，进入 Lustre 客户端栈。
2. **客户端元数据层 (MDC/LMV)**：若文件无句柄或需要元数据，llite 通过 MDC→MDT 交互获取 FID、布局 (lov_stripe_md) 等，更新本地 inode 缓存。
3. **锁与数据访问准备**：llite 向 LDLM 请求 extent lock，保证本地缓存数据一致性，并建立对对应 OST 的 OSC 句柄。
4. **页缓存与 OSC**：实际读写由 OSC 处理，将要传输的页面排队到 rpc_page_list，决定使用已有锁或重新申请。
5. **RPC 封装**：OSC 根据条带信息拆分请求，对每个 OST 构造 ptlrpc_request，将页描述变成 bulk 传输或普通 RPC。
6. **LNet 发送**：ptlrpc 调度器将请求交给 LNet，LNet 选择 LND（如 TCP），处理重试、流控和网络路由，发往目标 OST。
7. **服务器端处理**：OST 收到请求后，由 ptlrpc/obdfilter 解包，执行 dt_object 层读写，访问后端存储（ldiskfs/ZFS）。
8. **结果返回**：结果沿原路径返回客户端，客户端更新页缓存状态、清理/延长锁并唤醒应用线程；写入时可能触发后续写回或刷新。

控制面（元数据）、数据面（条带 I/O）与锁管理协同工作，保证跨多个 OST 的并行访问和一致性。

