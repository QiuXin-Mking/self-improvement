# q
如何开启 Lustre 的调试日志功能？
# a
使用 `lctl set_param debug=-1` 开启最详细的调试日志：
```bash
lctl set_param debug=-1
```
可以通过 `lctl get_param debug` 验证当前调试级别。

# q
如何将 Lustre 内核调试日志保存到文件？
# a
使用 `lctl debug_kernel` 命令将当前内核调试缓冲区内容保存到指定日志文件：
```bash
lctl debug_kernel /tmp/lustre.log
```

# q
如何关闭 Lustre 的详细调试日志？
# a
将调试级别恢复为仅记录关键错误和警告：
```bash
lctl set_param debug="ioctl neterror warning error emerg ha config console lfsck"
```
该设置将停止记录过于详细的调试信息，仅保留重要事件。

# q
如何检查所有 MDT 进程的总 CPU 占用率？
# a
使用 `top` 结合 `grep` 和 `awk` 统计所有 `mdt` 相关进程的 CPU 使用率总和：
```bash
top -bn1 | grep mdt | awk 'BEGIN{s=0}{s+=$9}END{print s}'
```
其中 `$9` 表示 CPU 使用率列（取决于 `top` 输出格式，通常如此）。若只查看单个进程，可省略 awk 部分。

# q
如何临时开启 Lustre 调试日志并在指定时间段后自动收集和关闭？
# a
使用以下步骤记录特定时间段的调试信息，并自动恢复到默认日志级别：
```bash
lctl mark "qx1 start "
lctl set_param debug=-1
sleep 10
lctl mark "qx1 end "
lctl debug_kernel /tmp/lustre.log
lctl set_param debug="ioctl neterror warning error emerg ha config console lfsck"
```
`lctl mark` 在日志中插入标记，方便定位；`sleep 10` 控制采集时长；最后将日志保存并关闭调试模式。

