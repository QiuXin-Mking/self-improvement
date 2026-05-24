# q
如何使用 fio 对块设备进行顺序读性能测试（1M块大小、4任务、16队列深度）？
# a
```bash
fio --name=test --filename=/dev/vdb --rw=read --bs=1024k --direct=1 --size=1G --numjobs=4 --iodepth=16 --time_based --runtime=60 --group_reporting
```
该命令对 `/dev/vdb` 执行直接 I/O 的 1 MiB 顺序读，4 个并发任务，每个任务 i/o 深度为 16，持续 60 秒并输出聚合报告。

# q
obdfilter-survey 测试 Lustre OST 性能时如何指定测试参数？
# a
使用格式：`nobjhi=2 thrhi=2 size=1024 case=disk targets="oss:lustre1-OST0000" sh obdfilter-survey`，其中：
- `nobjhi`：最大对象数
- `thrhi`：最大线程数
- `size`：测试文件大小（单位 MB，此处 1024 MB）
- `case=disk`：测试类型为磁盘级性能
- `targets`：指定要测试的 OST（格式为 `oss:<OST名称>`）

# q
如何在 OSS 节点上列出所有 obdfilter 设备并检查其状态？
# a
在 OSS 节点上执行：
```bash
lctl dl | grep obdfilter
```
输出示例：`3 UP obdfilter lustre1-OST0000 lustre1-OST0000_UUID 5`，表明 OST0000 设备处于 UP 状态。

