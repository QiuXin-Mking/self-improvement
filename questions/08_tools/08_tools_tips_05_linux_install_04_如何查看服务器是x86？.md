# q
使用哪个命令可以快速输出当前系统的CPU架构名称？
# a
```bash
uname -m
```
输出示例：`x86_64`（64位x86）、`i686`（32位x86）、`armv7l` 等。

# q
在 `/proc/cpuinfo` 文件中，哪个标志位表示CPU支持64位（Long Mode）？
# a
`lm` 标志。执行 `cat /proc/cpuinfo` 后在 `flags` 行中查找是否包含 `lm`，若包含则表明服务器支持64位架构。

# q
`lscpu` 命令输出中，哪个字段明确显示了CPU架构信息？
# a
`Architecture` 字段。运行 `lscpu` 后查找该行即可获取架构信息。

