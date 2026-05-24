# q
sgpdd-survey 是什么工具，使用时必须指定什么参数？
# a
sgpdd-survey 是 Lustre 提供的磁盘性能测试工具，使用时必须通过环境变量 `scsidevs` 或 `rawdevs` 指定待测设备，且两者只能指定其一。例如：
```bash
scsidevs=/dev/vdb sh sgpdd-survey
```
或
```bash
rawdevs=/dev/raw/raw1 sh sgpdd-survey
```

# q
使用 sgpdd-survey 测试 raw 设备时可能遇到哪些常见错误，如何解决？
# a
常见错误及解决：
- **“Can't find SG device... Do you have the sg module configured for your kernel?”**：需要加载 sg 内核模块（`modprobe sg`），确认 `lsmod | grep sg` 有输出。
- **“READ CAPACITY (16) failed... device ... not big enough”**：raw 设备可能未正确绑定到块设备，需先用 `raw` 命令绑定，如 `sudo raw /dev/raw/raw1 /dev/nvme0n1`，并可通过 `size` 参数限制测试区域（如 `size=50`）。
- 使用 `-v` 选项可查看详细错误信息。

# q
obdfilter-survey 工具的主要用途是什么，如何执行一个简单的测试？
# a
obdfilter-survey 用于测试 OST 的性能。执行示例：
```bash
nobjhi=2 thrhi=2 size=1024 targets="lustre1-OST0000" sh obdfilter-survey
```
其中 `targets` 指定要测试的 OST 名称，`nobjhi` 和 `thrhi` 控制并发对象和线程数，`size` 指定文件大小（单位 MB）。

# q
根据文档，作者对 sgpdd-survey 的最终评价是什么？有什么替代建议？
# a
作者结论：sgpdd-survey 工具箱不好用，存在诸多兼容和使用问题，不如直接使用 FIO 工具测试裸盘性能。推荐用 FIO 进行磁盘基准测试。

