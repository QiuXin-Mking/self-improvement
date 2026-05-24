# q
介绍如下术语 LBUG
# a
A fatal error condition detected by the software that halts execution of the kernel thread to avoid potential further corruption of the system state. It is printed to the console log and triggers a dump of the internal debug log. The system must be rebooted to clear this state.
软件检测到致命错误状态，会暂停内核线程的执行以避免系统状态可能出现的进一步损坏。该错误会打印到控制台日志，并触发内部调试日志的转储。必须重启系统才能清除此状态。

# q
介绍如下术语 LDLM
# a
Lustre Distributed Lock Manager.

# q
介绍如下术语 LFSCK (lfs ck，lfs check)
# a
Lustre 文件系统检查（lfsck），是磁盘文件系统检查程序的分布式版本。通常情况下，无需运行 lfsck，仅当文件系统因多磁盘故障等事件损坏，且无法通过文件系统日志恢复功能修复时，才需要运行它。

# q
介绍如下术语 LMV
# a
Logical Metadata Volume. A module that implements a DNE client-side abstraction device. It allows a client to work with many MDTs without changes to the llite module. The LMV code forwards requests to the correct MDT based on name or directory striping information and merges replies into a single result to pass back to the higher llite layer that connects the Lustre file system with Linux VFS, supports VFS semantics, and complies with POSIX interface specifications.
逻辑元数据卷（LMV），是一个实现 DNE 客户端侧抽象设备的模块。它允许客户端在不修改 llite 模块的情况下，与多个 MDT 协同工作。LMV 代码会根据名称或目录条带化信息，将请求转发至正确的 MDT；同时将多个 MDT 的响应合并为单一结果，回传给上层的 llite 层。该 llite 层负责连接 Lustre 文件系统与 Linux VFS，支持 VFS 语义，并遵循 POSIX 接口规范。

# q
介绍如下术语 LNet
# a
Lustre networking. A message passing network protocol capable of running and routing through various physical layers. LNet forms the underpinning of LNETrpc.

