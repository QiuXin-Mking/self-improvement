# q
如何挂载Lustre客户端文件系统（支持多MGS节点）？
# a
使用 `mount -t lustre` 并指定 MGS NID 和文件系统名，可同时列出多个 MGS 节点以实现故障切换：
```bash
mount -t lustre 192.168.6.172@tcp:/nas_test /mnt/lustre
mount -t lustre 192.168.6.172@tcp:192.168.6.174@tcp:192.168.6.175@tcp:/nas_test /mnt/lustre
```

# q
如何查看文件在Lustre中的条带分布和所在OST对象信息？
# a
使用 `lfs getstripe` 命令，可查看常规文件或指定文件：
```bash
lfs getstripe /mnt/lustre/q/*
lfs getstripe /mnt/lustre/q/Fio.7.0
```
输出包含 `lmm_stripe_count`、`lmm_stripe_size`、`lmm_pattern` 以及每个 stripe 对应的 `obdidx`、`objid` 等。

# q
如何运行FIO对Lustre挂载点进行混合读写性能测试（并获取尾部结果）？
# a
典型命令：在 `/mnt/lustre/$(hostname)` 下执行 70% 读混合随机读写，1M 块，16 任务，队列深度 8，30 秒时长，最后两行结果：
```bash
fio --name=Fio --directory=/mnt/lustre/$(hostname) --rw=rw --rwmixread=70 \
    --bs=1M --size=1m --numjobs=16 --iodepth=8 --direct=1 \
    --ioengine=libaio --time_based --timeout=30 --group_reporting | tail -n 2
```

# q
如何开启Lustre所有调试信息以排查问题？
# a
通过 `lctl set_param` 全局开启调试：
```bash
lctl set_param debug=+all
```

# q
如何修改Lustre目标（如OST）的服务节点和MGS节点参数？
# a
使用 `tunefs.lustre` 擦除旧参数并写入新服务节点（支持多节点 failover）：
```bash
tunefs.lustre --erase-params --servicenode=192.168.6.172@tcp \
              --servicenode=192.168.6.174@tcp /dev/rbd0
```

