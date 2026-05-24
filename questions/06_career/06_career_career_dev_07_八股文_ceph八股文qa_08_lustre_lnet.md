# q
LNet 源代码中 lnet/lnet/ 目录下各个 .c 文件的主要功能是什么？
# a
api-ni.c: 核心 API 和网络接口管理
lib-move.c: 消息移动/路由/发送
lib-msg.c: 消息生命周期管理
lib-md.c: 内存描述符管理
lib-me.c: 匹配入口管理
lib-ptl.c: Portal 管理
peer.c: 对等节点管理
router.c: 路由功能
config.c: 配置解析
acceptor.c: 连接接受器
lib-socket.c: 套接字工具
nidstrings.c: NID 字符串解析
udsp.c: 用户定义选择策略
lo.c: 回环网络驱动
lnet_rdma.c: RDMA 支持
net_fault.c: 网络故障模拟
router_proc.c: 路由器 proc 接口
module.c: 模块初始化/退出

