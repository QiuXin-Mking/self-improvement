# q
Lustre客户端发起一次文件I/O到OST的完整路径是怎样的？
# a
主要流程（自上而下）：  
1. 用户调用 `read()/write()`，VFS 进入 Lustre 客户端栈的 `llite`。  
2. 客户端通过 MDC→MDT 获取文件的 FID、布局（`lov_stripe_md`）等元数据，并更新本地 inode 缓存。  
3. 向 LDLM 请求相应的 `extent lock` 保证缓存一致性，建立对目标 OST 的 OSC 句柄。  
4. OSC 将页面排队到 `rpc_page_list`，按条带信息拆分请求。  
5. 每个 OST 独立构造 `ptlrpc_request`，底层将页描述转为 bulk 传输或普通 RPC。  
6. ptlrpc 调度器把请求交给 LNet，LNet 根据网络类型（如 TCP）处理重试、流控和路由，发送到目标 OST。  
7. 服务器端 OST 由 ptlrpc/obdfilter 解包，执行 `dt_object` 层的读写（访问后端 ldiskfs/ZFS）。  
8. 操作结果沿原网络路径返回客户端，客户端更新页缓存、清理/延长锁并唤醒应用线程；写操作按策略触发写回或刷新流程。

# q
Lustre中负责客户端数据面读写和条带拆分的组件是什么？
# a
OSC（Object Storage Client），它根据条带布局将 I/O 请求按 OST 拆分，管理页面排队和 RPC 封装。

# q
提高Lustre元数据服务性能的常见方法有哪些？
# a
部署多个 MDT（元数据目标）并利用远程 MDT 分担负载，从而提升元数据操作的并行处理能力。

# q
宏杉科技在OSD位图回收方面做了什么优化？
# a
将 OSD bitmap 位图的异步回收改为同步回收，以改善回收行为或脏数据管理。

