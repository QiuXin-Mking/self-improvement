# q
lustre 如果在配置文件中选择tcpip网卡，会走rdma嘛
# a
不会。配置里把网络类型选成 tcp（或指定 TCP/IP 网卡）时，LNet 会加载 socklnd ，走的是普通的 TCP 套接字栈，不会自动切换到 RDMA。
要用 RDMA，需要显式把网络类型设为对应的 RDMA LND（例如 o2ib, v2ib, infinihost 等），并确保 RDMA 驱动/库已启用。

[root@lustre_back1 ~]# modinfo  ksocklnd | grep description
description:    TCP Socket LNet Network Driver

[root@lustre_back1 ~]# lsmod | grep lnd
ksocklnd              188416  1
lnet                  933888  12 osc,ost,mgs,obdclass,ofd,ptlrpc,mgc,ksocklnd,lmv,mdt,lustre
libcfs                172032  22 fld,lnet,osc,fid,lod,mdd,ost,mgs,obdclass,osp,ofd,ptlrpc,mgc,ksocklnd,lov,mdc,lmv,mdt,osd_ldiskfs,lustre,lquota,lfsck

# q
RDMA 比 tcp ip链接优秀在什么地方
# a
1. 零拷贝（Zero Copy）
    RDMA： 允许一台主机直接将内存中的数据拷贝到另一台主机的内存中，全程无需CPU干预，无需内核上下文切换和应用层/内核层之间的数据拷贝。这就可以做到数据直接用户空间→用户空间，是真正的“零拷贝”（Zero Copy）。
    TCP/IP： 网络IO通常需经过多次内核缓冲数据拷贝（应用空间→内核空间→网卡），且需要CPU参与数据收发。

2. 超低延迟
    RDMA： 带宽高、时延低，由于省去了协议栈处理和内核中断、拷贝等环节，端到端延迟极低，微秒级甚至亚微秒级，非常适合对时延极其敏感的场景（如金融交易、HPC等）。
    TCP/IP： 时延高，涉及较多的协议堆栈处理、上下文切换。

3. 高吞吐
    RDMA： 可以发挥网络硬件的极限带宽（比如IB/Infiniband/ROCE等），并发能力强，资源开销极小。
    TCP/IP： 性能很大程度上受系统协议栈和资源消耗影响，吞吐不如RDMA高。

4. CPU占用低
    RDMA： 数据拷贝和协议卸载全由网卡（智能网卡，例如RDMA支持的HCA）完成，CPU几乎无需参与。
    TCP/IP： 数据流经协议栈、处理上下文切换和缓存，CPU消耗较高。

5. 支持更高的IO密度
    RDMA： 因为CPU占用低，可以支撑更多的连接并发和更高的IO密度，非常适合分布式存储、分布式数据库等高并发场景。
    TCP/IP： 连接数一多，CPU容易成为瓶颈。

6. 适合大规模分布式与高性能计算
    Massively parallel computing（HPC）、AI训练、大数据分布式存储等场景，几乎都会首选RDMA作为底层高效互联通道。

