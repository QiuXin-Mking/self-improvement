# q
如何调整Ceph恢复速率？可以调整哪些配置参数？
# a
```bash
ceph config set osd osd_recovery_max_active 5
ceph config set osd osd_max_backfills 2
ceph config set osd osd_recovery_op_priority 10
ceph config set osd osd_recovery_sleep 0.1
```

# q
Ceph 开关 backfilling 和 recovery 的命令是什么？
# a
```bash
ceph osd set nobackfill
ceph osd set norecover

ceph osd unset nobackfill
ceph osd unset norecover
```

# q
/sys/block/$bcache_device/bcache/sequential_cutoff 的作用是什么？
# a
顺序数据流量大于该设定值时，不会用SSD缓存，只走后端慢盘，有助于提高SSD空间利用率，提升系统整体IO性能。适合让SSD专注随机I/O。sequential_cutoff 是“顺序I/O临界值”，也叫“顺序访问停止缓存阈值”。

# q
解释一下条带宽度(stripe width)、条带深度(stripe depth)、条带长度(stripe size/stripe unit size)。
# a
- Stripe Width = “多少块盘一起并行”
- Stripe Depth/Unit = “每次在一块盘连续写入数据量”
- Stripe Size = “轮一次所有盘一共写的总量”
- 条带长度 = 条带深度 × 条带宽度

