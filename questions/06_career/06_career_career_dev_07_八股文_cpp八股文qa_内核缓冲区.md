# q
什么是内核缓冲区？
# a
内核缓冲区（Kernel Buffer）是操作系统内核中用于临时存储数据的内存区域，是内核与硬件、用户空间或不同子系统之间进行数据传输时的“中间暂存区”，其核心作用是协调速度差异、提高IO效率、保证数据一致性。

# q
内核缓冲区按作用方式主要分为哪几类？
# a
主要分为三类：
1. 文件IO缓冲区（Page Cache）
2. 磁盘IO缓冲区（Disk Buffer/Cache）
3. 网络IO缓冲区（收发缓冲区）

# q
应用程序写文件时，数据是如何经过文件IO缓冲区（Page Cache）的？
# a
调用 write() 时，数据先存入 Page Cache，而不是直接写入磁盘。内核会在合适时机（如缓冲区满、调用 fsync()、close() 或定时回写线程）将数据刷到磁盘。

# q
网络IO缓冲区的主要作用是什么？
# a
平衡 CPU 处理速度与网络带宽差异，应对网络拥塞或流量波动：发送数据时先进入内核 Send Buffer，由协议栈分包发送；接收数据时网卡将包存入 Receive Buffer，应用再通过 recv() 读取。

# q
Lustre OST 会启动哪几类线程？
# a
ll_ost_out、ll_ost_seq、ll_ost_io、ll_ost_create、ll_ost

