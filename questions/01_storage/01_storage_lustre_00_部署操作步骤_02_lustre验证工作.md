# q
在Lustre中，如何临时和持久化修改jobid_var参数？
# a
临时修改：
```bash
lctl set_param jobid_var=procname_uid
```
持久化修改（在MGS上执行，配置会自动传播到MDS、OSS和客户端）：
```bash
lctl conf_param testfs.sys.jobid_var=SLURM_JOB_ID
```

# q
运行obdfilter-survey时提示 `./iokit-libecho: No such file or directory` 应如何解决？
# a
需要进入包含iokit-libecho等依赖脚本的正确目录（如 `/usr/bin/` 或lustre-iokit的安装目录），再执行测试命令。例如：
```bash
cd /usr/bin/
nobjhi=2 thrhi=2 size=1024 targets="lustre1-OST0000" sh obdfilter-survey
```

# q
sgpdd-survey工具运行时要求必须指定什么参数？
# a
必须指定 `--scsidevs` 或 `--rawdevs` 参数来明确要测试的设备，例如：
```bash
sgpdd-survey --scsidevs=/dev/vdb
```
如果未提供任一参数，命令行会输出 `Must either specify scsidevs or rawdevs` 并退出。

# q
如何查看或调整块设备的内核队列参数，例如最大扇区数和I/O调度器？
# a
查看最大扇区大小（单位KB）：
```bash
cat /sys/block/vdb/queue/max_sectors_kb
```
查看并修改调度器（如设置为deadline）：
```bash
cat /sys/block/vdb/queue/scheduler
echo deadline > /sys/block/vdb/queue/scheduler
```
其他常用参数包括 `max_phys_segments` 等，建议根据性能需求将调度器设为 `deadline` 或 `noop`。

