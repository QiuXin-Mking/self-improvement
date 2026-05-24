# q
Ceph监控项目主要包括哪些性能指标？
# a
ops（操作数）、throughput（吞吐量）、usage（使用率）、time cost（耗时）、req W（写请求）、req R（读请求）、success req W（成功写请求）、success req R（成功读请求）、iops（每秒操作数）

# q
在Lustre系统中，如何查找rpc_stats文件的位置？
# a
使用命令 `find / -name "rpc_stats"`，常见路径包括 `/proc/fs/lustre/mdc/*/rpc_stats`（客户端到MDT的RPC统计）、`/proc/fs/lustre/osc/*/rpc_stats`（客户端到OST的RPC统计）以及 `/sys/kernel/debug/lustre/osp/*/rpc_stats`（MDT到OST的统计）

# q
如何通过lctl命令获取Lustre的详细性能统计？
# a
可使用 `lctl get_param -n obdfilter.*.stats` 查看服务器端过滤器统计（如 write_bytes、read_bytes、操作耗时等）；使用 `lctl get_param -n osc.*.stats` 查看客户端OSC统计（包括 `req_waittime`、`req_active`、`ost_read`、`ost_write` 等指标）。用 `lctl list_param -F osc.*` 可以列出所有可用的OSC参数

