# q
Lustre 调试日志中每条记录的输出格式包含哪些字段，各字段如何对应？
# a
Lustre 调试日志单条记录由 `debug.c` 中的 `scnprintf` 按如下格式输出：
```
%08x:%08x:%u.%u%s:%u.%06llu:%u:%u:%u:(%s:%u:%s()) %s
```
- `ph_subsys`（子系统号）、`ph_mask`（掩码）
- `ph_cpu_id`（CPU ID）、`ph_type`（类型）
- `ph_flags & PH_FLAG_FIRST_RECORD` 若为真则输出 "F"
- `ph_sec`（秒）、`ph_usec`（微秒）
- `ph_stack`（栈深度）、`ph_pid`（PID）、`ph_extern_pid`（外部 PID）
- `line->file`（源文件名）、`ph_line_num`（行号）、`line->fn`（函数名）
- `line->text`（日志文本）

示例：
```
00010000:00010000:2.0:1746496640.686162:0:11859:0:(ldlm_lockd.c:2841:ldlm_bl_get_work()) Process entered
```
对应关系为：
```
ph_subsys:ph_mask:cpu_id.ph_type:sec.usec:ph_stack:pid:extern_pid:(file:line:function()) text
```

# q
`lctl filter all_types` 和 `lctl filter all_subs` 的作用分别是什么？
# a
- `lctl filter all_types`：禁用所有日志类型的输出（如 trace、inode、super、iotrace、malloc、cache 等）。
- `lctl filter all_subs`：禁用所有子系统（如 mdc、mds、osc、ost、class、log、llite、lnet、ldlm、lov 等）的日志输出。

这两条命令用于全局关闭调试信息，后续可通过单独的 `lctl filter` 启用特定类型或子系统。

# q
如何查看 Lustre 支持的日志类型（types）和子系统（subsystems）？
# a
使用以下命令：
```bash
lctl debug_list types   # 列出所有日志类型
lctl debug_list subs    # 列出所有子系统
```
常用类型包括 `trace, inode, super, iotrace, malloc, cache, info, ioctl, neterror, net, warning, error, emerg, dlmtrace, rpctrace, vfstrace, config, console, quota, sec, hsm, layout` 等。  
常用子系统包括 `mdc, mds, osc, ost, class, log, llite, rpc, mgmt, lnet, lnd, ldlm, lov, lquota, osd, lfsck, sec, mgc, mgs, fid, fld` 等。

