# q
stress-ng 工具的主要用途是什么？
# a
stress-ng 是一个 Linux 基准测试与压力测试工具，用于评估系统在极端条件下的性能和稳定性。它可以对 CPU、内存、I/O 子系统等硬件组件施加受控且可重复的高强度负载，帮助识别潜在瓶颈。

# q
如何通过命令行安装 stress-ng（基于 Debian/Ubuntu）？
# a
```bash
sudo apt update
sudo apt install stress-ng
```

# q
如何编写一条同时测试 CPU、内存和磁盘 I/O 的综合压力命令，并限制运行时间？
# a
```bash
stress-ng --cpu 4 --io 2 --vm 2 --vm-bytes 512M --timeout 60s
```
该命令启动 4 个 CPU 负载、2 个 I/O 负载和 2 个虚拟内存负载（各分配 512MB），总运行时长为 60 秒。

# q
stress-ng 输出中的 “bogo ops” 代表什么？其用途是什么？
# a
bogo ops 表示压力测试期间完成的操作数量，它不是一个标准化的精确性能指标，仅用于粗略估计系统性能或观察相对趋势，不能作为绝对基准值使用。

# q
使用 `--hdd` 选项进行硬盘压力测试时需要注意什么？
# a
`--hdd` 会对磁盘执行大量读写操作，存在数据丢失或硬盘损坏风险。只能在受控环境下使用，且必须确保目标设备上没有重要数据。

