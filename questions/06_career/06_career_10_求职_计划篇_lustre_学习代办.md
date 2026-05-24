# q
Lustre学习代办中列出的核心主题有哪些？
# a
核心主题包括：FID、layout、xattr、ACL、quota、OSS affinity、对象预取与回写、LNET（跨节点通信与路由）、Client（VFS/LOV/OSC三层）、page cache，以及DIO、aio、mmap、read-ahead/write-back、lockless I/O路径、LDLM分布式锁管理器、MDT/OST恢复流程（日志回放、客户端replay）等。

# q
Lustre客户端涉及哪三个层次？
# a
Client 端分为 VFS、LOV、OSC 三层。

# q
LNET在Lustre中需要掌握哪些知识点？
# a
需熟悉 LNet NID、路由策略、网络容错与多网卡绑定，理解其跨节点通信与路由原理。

# q
Lustre分布式锁管理器（LDLM）相关的学习点包括哪些？
# a
包括锁模式、范围锁（extent lock）以及 glimpse 的语义。

# q
MDT/OST恢复流程涉及什么关键操作？
# a
涉及日志回放（journal replay）和客户端 replay。

