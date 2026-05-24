# q
使用fio对Lustre进行纯写测试（direct io）的完整命令是什么？要求numjobs=72，iodepth=128，ioengine=libaio，块大小1M，文件大小10G，目录为/mnt/lustre/$(hostname)，timeout=3000，并输出分组报告。
# a
```bash
fio -numjobs=72 -iodepth=128 -direct=1 -ioengine=libaio -rw=write -bs=1M -size=10G -name=Fio -directory=/mnt/lustre/$(hostname) --time_based --timeout=3000 -group_reporting
```

# q
使用fio对Lustre进行读写混合测试（direct io）的完整命令是什么？与纯写测试相比只改变读写模式，其他参数相同。
# a
```bash
fio -numjobs=72 -iodepth=128 -direct=1 -ioengine=libaio -rw=rw -bs=1M -size=10G -name=Fio -directory=/mnt/lustre/$(hostname) --time_based --timeout=3000 -group_reporting
```

# q
Linux内核参数vm.dirty_expire_centisecs和vm.dirty_writeback_centisecs分别表示什么？如何临时生效和永久生效？
# a
- vm.dirty_expire_centisecs：脏页在内存中可逗留的最长时间，单位为1/100秒，超时会被后台线程刷写。例如设为500表示5秒。
- vm.dirty_writeback_centisecs：后台线程发起脏页刷新操作的间隔，单位为1/100秒。例如设为100表示1秒一次刷新。
临时生效（重启失效）：
```bash
sudo sysctl -w vm.dirty_expire_centisecs=500
sudo sysctl -w vm.dirty_writeback_centisecs=100
```
永久生效：在/etc/sysctl.conf末尾添加：
```
vm.dirty_expire_centisecs = 500
vm.dirty_writeback_centisecs = 100
```
然后执行 `sudo sysctl -p` 使其生效。

