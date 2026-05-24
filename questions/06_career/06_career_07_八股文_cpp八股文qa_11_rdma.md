# q
Lustre 如果在配置文件中选择 TCP/IP 网卡，会走 RDMA 吗？
# a
不会。配置里把网络类型选成 tcp 时，LNet 会加载 ksocklnd，走普通 TCP 套接字栈，不会自动切换到 RDMA。要用 RDMA，需显式设置网络类型为 RDMA LND（如 o2ib, v2ib 等），并确保 RDMA 驱动/库已启用。
```
[root@lustre_back1 ~]# modinfo  ksocklnd | grep description
description:    TCP Socket LNet Network Driver

[root@lustre_back1 ~]# lsmod | grep lnd
ksocklnd              188416  1
lnet                  933888  12 osc,ost,mgs,obdclass,ofd,ptlrpc,mgc,ksocklnd,lmv,mdt,lustre
libcfs                172032  22 fld,lnet,osc,fid,lod,mdd,ost,mgs,obdclass,osp,ofd,ptlrpc,mgc,ksocklnd,lov,mdc,lmv,mdt,osd_ldiskfs,lustre,lquota,lfsck
```

# q
RDMA 相比 TCP/IP 连接有哪些优势？
# a
1. 零拷贝（Zero Copy）：允许一台主机直接将内存中的数据拷贝到另一台主机的内存中，无需 CPU 干预和内核切换，实现数据用户空间到用户空间。
2. 超低延迟：省去协议栈处理和内核中断、拷贝等环节，端到端延迟极低（微秒级甚至亚微秒级）。
3. 高吞吐：能够发挥网络硬件的极限带宽，并发能力强，资源开销极小。
4. CPU 占用低：数据拷贝和协议卸载由网卡（如 RDMA 支持的 HCA）完成，CPU 几乎无需参与。
5. 支持更高的 IO 密度：CPU 占用低，可以支撑更多并发连接和高 IO 密度，适合分布式存储、分布式数据库等高并发场景。
6. 适合大规模分布式与高性能计算：HPCC、AI 训练、大数据分布式存储等场景，几乎都会首选 RDMA 作为底层高效互联通道。

