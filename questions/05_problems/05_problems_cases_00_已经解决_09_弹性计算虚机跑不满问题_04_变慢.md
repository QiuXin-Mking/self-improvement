# q
如何临时开启Ceph RGW的调试日志来排查请求？
# a
使用 `ceph daemon` 命令动态调整 `debug_rgw` 参数：
```bash
ceph daemon /var/run/ceph/client.rgw.<hostname>.asok config set debug_rgw 20
```
排查结束后关闭调试：
```bash
ceph daemon /var/run/ceph/client.rgw.<hostname>.asok config set debug_rgw 0
```

# q
从 iostat 输出如何判断 NVMe 磁盘已成为性能瓶颈？
# a
关注 `%util` 指标。若 `%util` 接近 100%（如案例中 nvme0n1 的 `%util` 为 93.20），且伴随高读吞吐（`rkB/s` 达 192040.00），说明磁盘 IO 已接近饱和，即使响应时间（`await` 仅 0.30ms）较低，高并发读也会导致排队，成为虚机 IO 变慢的根本原因。

# q
系统 CPU iowait 不高（如 0.48%）但磁盘 %util 极高，这说明什么？
# a
说明磁盘响应快（低 `await`）但并发请求量大，磁盘带宽或 IOPS 已接近上限。NVMe 设备在处理大量读请求时，即使单个 I/O 等待时间短，总体利用率仍可打满，导致虚机感知到的存储延迟增加，计算资源无法被充分利用（虚机“跑不满”）。

# q
Ceph 集群中虚机变慢的典型存储侧根因是什么？
# a
Ceph 集群面临极大的客户端读负载（如 `ceph -s` 显示 1.4 GiB/s 读、12.73k op/s），导致底层 NVMe 磁盘的 `%util` 超过 90%。磁盘 I/O 饱和使得虚机 I/O 延迟变大，虚机内的应用等待 I/O 而无法跑满分配的 CPU，表现为虚机“跑不满”。

