# q
如何使用fio对NVMe磁盘进行4k随机读性能测试？
# a
使用以下命令，指定块大小为4k，直接IO，16个并发任务，队列深度128，运行60秒：
```bash
fio --name=test --filename=/dev/nvme0n1p1 --rw=read --bs=4k --direct=1 --size=1G --numjobs=16 --iodepth=128 --time_based --runtime=60 --group_reporting
```

# q
如何使用fio对NVMe磁盘进行4k随机写性能测试？
# a
使用以下命令，将读写模式改为 `--rw=write`：
```bash
fio --name=test --filename=/dev/nvme0n1p1 --rw=write --bs=4k --direct=1 --size=1G --numjobs=16 --iodepth=128 --time_based --runtime=60 --group_reporting
```

# q
进行磁盘预热的目的及常用fio命令是什么？
# a
目的：通过全盘顺序写入，消除SSD初次写入的“写入放大”或性能波动，使后续测试结果更稳定。常用命令：
```bash
fio --name=warmup --filename=/dev/nvme0n1 --rw=write --bs=1M --direct=1 --size=100% --numjobs=1 --iodepth=32
```
若需后台运行，可追加 `> fio_warmup.log 2>&1 &` 或使用 `nohup`。

# q
在本次测试中，NVMe磁盘4k随机读和随机写的IOPS及带宽各是多少？
# a
- 随机读：IOPS=132k，带宽 BW=514MiB/s (539MB/s)
- 随机写：IOPS=383k，带宽 BW=1495MiB/s (1567MB/s)

