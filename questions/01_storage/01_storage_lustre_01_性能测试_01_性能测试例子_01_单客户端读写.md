# q
单客户端Lustre性能基线测试中，使用 fio 256 jobs、1024 iodepth、1M 块大小的读写带宽分别是多少？
# a
读带宽为 2687MiB/s (2817MB/s)，写带宽为 2635MiB/s (2763MB/s)。

# q
Lustre 性能调优时，如何通过 modprobe.conf 增加 ptlrpc 每 CPU 分区线程数和 ksocklnd 的 credit？
# a
在 `/etc/modprobe.d/modprobe.conf` 中添加以下两行，然后重启节点：
```
options ptlrpc ptlrpcd per cpt max=32
options ksocklnd credits=2560
```

# q
在 Lustre 客户端使用 `lctl set_param` 调优时，`osc.*.max_dirty_mb` 参数的作用是什么？测试中将其设置为何值？
# a
`max_dirty_mb` 控制客户端缓存的未写入脏数据上限（单位 MB），用于限制回写速度和减少内存压力。测试中设置为 64：
```bash
sudo lctl set_param osc.*.max_dirty_mb=64
```

# q
如何通过 `lctl set_param` 调整 Lustre 客户端与 OST、MDT 之间的并发 RPC 数量？
# a
可分别设置以下参数：
- 调整与每个 OST 的最大并发 RPC：`sudo lctl set_param osc.*.ost*.max_rpcs_in_flight=32`
- 调整与每个 MDT 的最大并发 RPC：`sudo lctl set_param mdc.*.max_rpcs_in_flight=64`
- 调整修改类 RPC 的最大并发数：`sudo lctl set_param mdc.*.max_mod_rpcs_in_flight=50`

