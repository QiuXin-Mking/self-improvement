# q
如何使用命令查看Ceph集群中各pool的实时IO负载？
# a
执行 `ceph osd pool stats` 命令，可列出所有 pool 的客户端 IO 速率和操作数。示例输出中高负载的 pool 会显示较高的读写带宽和 ops，如 `obj_data_3fb_045b35ee` 达到 952 MiB/s 读、33 MiB/s 写，`obj_index_3fb_8b640e3d` 达到 7.38k op/s 读。

# q
从 `ceph osd pool stats` 输出中，如何判断某个 pool 正在承受高压力？
# a
观察以下指标：
- 高读写带宽（MiB/s rd/wr），如 `952 MiB/s rd, 33 MiB/s wr`
- 高操作数（op/s rd/wr），如 `453 op/s rd, 374 op/s wr`
- 索引 pool 尤其关注 op/s，因为小 IO 频繁，如 `7.38k op/s rd` 表示大量索引读取。

