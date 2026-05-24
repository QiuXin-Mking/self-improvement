# q
如何对Lustre文件系统执行1M顺序写入性能测试并将结果保存到指定文件？
# a
```bash
nohup fio -numjobs=16 -iodepth=128 -direct=1 -ioengine=libaio -rw=write -bs=1M -size=10G -name=Fio -directory=/mnt/lustre --time_based --timeout=60 -group_reporting --output=/mnt/qiuxin/perfo/1125/1m_write_test_cas1 &
```

# q
如何对Lustre文件系统执行4K随机读性能测试并将结果保存到指定文件？
# a
```bash
nohup fio -numjobs=16 -iodepth=128 -direct=1 -ioengine=libaio -rw=randread -bs=4k -size=10G -name=Fio -directory=/mnt/lustre --time_based --timeout=60 -group_reporting --output=/mnt/qiuxin/perfo/1125/4k_read_test_cas1 &
```

# q
如何使用fio对原始块设备进行全盘顺序写入预热？
# a
```bash
nohup fio --name=warmup --filename=/dev/vdb --rw=write --bs=1M --direct=1 --size=100% --numjobs=1 --iodepth=32 > fio_warmup0.log 2>&1 &
```

# q
如何从fio测试结果文件中提取写入带宽的数值？
# a
```bash
cat /mnt/qiuxin/perfo/1125/1m_write_test | grep WRITE | awk '{print $2}'
```

