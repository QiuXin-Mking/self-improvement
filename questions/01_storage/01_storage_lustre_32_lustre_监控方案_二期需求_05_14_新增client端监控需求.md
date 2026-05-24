# q
如何获取Lustre client端的cache使用量指标？
# a
执行命令 `cat /sys/kernel/debug/lustre/llite/*/max_cached_mb | grep '^used_mb'`，输出单位为 MB。

# q
在Lustre client端监控中，读带宽和写带宽的数据获取方式及单位是什么？
# a
读取 `/etc/lustre_mon/op_bw.log` 文件，在 `[client]` 节点下取 `read_bytes`（读带宽）和 `write_bytes`（写带宽），单位均为 Byte。

# q
对于部分Lustre client机器缺少 `/etc/lustre/lustre_mon.conf` 和 `/etc/lustre/lustre_mon.py` 的情况，应如何部署监控采集？
# a
从已存在这些文件的机器（如 10.176.101.171）同步两个文件到目标机器，并在目标机器上新增 1 分钟采集的定时任务：`python3 lustre_mon.py local`。

# q
原有监控节点名称做了哪些重命名和移动？
# a
- “ost 时延读” → “ost 读时延”
- “ost 时延写” → “ost 写时延”
- “mdt cpu usage” → “mdt cpu使用率”
- “MDT总OPS” → “mdt 总ops”
- “MDT 各个op操作值” → “mdt 各类op操作数”
- “MDT 时延” → “mdt 各类op时延”
将原 node 小节的“client的 read ahead 使用量”和“client的cache 使用量”移至 client 小节，并分别更名为“client read ahead 使用量”和“client cache 使用量”。

