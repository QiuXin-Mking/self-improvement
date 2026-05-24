# q
调整Ceph恢复速率需要设置哪些配置参数？
# a
```shell
ceph config set osd osd_recovery_max_active 5
ceph config set osd osd_max_backfills 2
ceph config set osd osd_recovery_op_priority 10
ceph config set osd osd_recovery_sleep 0.1
```

# q
如何开关Ceph的backfilling和recovery？
# a
```shell
# 关闭
ceph osd set nobackfill
ceph osd set norecover

# 开启
ceph osd unset nobackfill
ceph osd unset norecover
```

# q
修改 /sys/block/$bcache_device/bcache/sequential_cutoff 的作用是什么？
# a
当顺序I/O流量大于该设定值时，数据不会缓存到SSD，而是直接写入后端慢盘。这样有助于提高SSD空间利用率，让SSD专注随机I/O，提升整体IO性能。该参数也称为“顺序访问停止缓存阈值”。

# q
请解释条带宽度（stripe width）、条带深度（stripe depth）和条带长度（stripe size）的含义。
# a
- 条带宽度（Stripe Width）：参与并行的磁盘数量，即多少块盘一起并行。
- 条带深度/条带单元（Stripe Depth/Unit）：每次在一块盘上连续写入的数据量。
- 条带长度（Stripe Size）= 条带深度 × 条带宽度，即轮询所有盘一次所写入的数据总量。

