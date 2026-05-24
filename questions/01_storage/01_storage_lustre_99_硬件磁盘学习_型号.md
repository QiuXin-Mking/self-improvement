# q
如何从 `nvme list` 输出中解读 NVMe 设备的扇区大小格式信息？
# a
`Format` 列显示扇区格式，如 `512   B +  0 B` 表示扇区大小为 512 字节，没有额外的元数据。

# q
Micron 9300 系列 NVMe SSD（型号 MTFDHAL7T6TDP）的顺序读写性能指标是多少？
# a
顺序读 3.5 GB/s，顺序写 3.1 GB/s。

# q
下面命令中 `fio` 的 `--direct=1` 参数有什么作用？
```bash
nohup fio -numjobs=32 -iodepth=128 -direct=1 -ioengine=libaio -rw=write -bs=128K -size=10G -name=Fio --filename=/dev/nvme0n1 --time_based --timeout=120 -group_reporting ... &
```
# a
`--direct=1` 表示使用直接 I/O（direct I/O），绕过操作系统的页缓存，直接读写块设备。

