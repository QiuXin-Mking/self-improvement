# q
fio测试Lustre纯写性能，要求numjobs=72、iodepth=128、direct IO、libaio引擎、块大小1M、文件大小10G、运行3000秒，写出命令。
# a
```bash
fio -numjobs=72 -iodepth=128 -direct=1 -ioengine=libaio -rw=write -bs=1M -size=10G -name=Fio -directory=/mnt/lustre/$(hostname) --time_based --timeout=3000 -group_reporting
```
参数说明：
- `-numjobs=72`：72个并发job
- `-iodepth=128`：IO队列深度128
- `-direct=1`：使用direct IO绕过系统缓存
- `-ioengine=libaio`：异步IO引擎
- `-rw=write`：纯写模式
- `-bs=1M`：块大小为1MB
- `-size=10G`：每个job写10GB
- `--time_based --timeout=3000`：基于时间运行3000秒（实际运行时覆盖size参数）
- `-group_reporting`：汇总报告

# q
哪些分布式存储系统理论上能支持EB乃至ZB级别？请列举一些对象存储和文件系统，并给出EB规模的大数据项目例子。
# a
业界多数对象存储和文件存储系统通过横向无限扩展，理论上均可支持EB甚至ZB级别，例如：
- 对象存储：Ceph、Amazon S3、Google Cloud Storage、Alibaba Cloud OSS、Huawei OBS、华为OceanStor、EMC ECS、IBM COS等
- 企业级分布式文件系统：Lustre、GPFS (IBM Spectrum Scale)、QFS、Red Hat GlusterFS、HDFS等

EB规模的大数据项目案例：“云上大基因组”、“气象仿真”、“天文望远镜”、“互联网平台归档”等长期在EB级别。

# q
Linux内核参数`vm.dirty_expire_centisecs`和`vm.dirty_writeback_centisecs`的作用分别是什么？如何临时生效和永久生效？
# a
- `vm.dirty_expire_centisecs`：脏页在内存中最长停留时间（单位1/100秒）。设为500表示5秒后仍未写回的脏页将被后台线程主动刷写。默认值通常为3000（30秒）。
- `vm.dirty_writeback_centisecs`：内核后台刷新线程的唤醒周期（单位1/100秒）。设为100表示每隔1秒执行一次脏页刷新。

**临时生效（重启后失效）**
```bash
sudo sysctl -w vm.dirty_expire_centisecs=500
sudo sysctl -w vm.dirty_writeback_centisecs=100
```

**永久生效**
1. 编辑`/etc/sysctl.conf`，添加以下两行：
```
vm.dirty_expire_centisecs = 500
vm.dirty_writeback_centisecs = 100
```
2. 执行`sudo sysctl -p`使其立即生效。

