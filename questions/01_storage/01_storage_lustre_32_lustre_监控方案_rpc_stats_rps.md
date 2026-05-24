# q
如何查看Lustre中不同组件（OSC、MDC、OSP）的RPC统计信息？
# a
使用不同的命令或文件路径：
- OSC（对象存储客户端）：`lctl get_param osc.<fsname>-OSTxxxx-osc-*.rpc_stats`  
  例如：`lctl get_param osc.nas_test-OST0000-osc-*.rpc_stats`
- MDC（元数据客户端）：`lctl get_param mdc.<fsname>-MDTxxxx-mdc-*.rpc_stats`  
  例如：`lctl get_param mdc.*.rpc_stats`
- OSP（对象存储代理）：`cat /sys/kernel/debug/lustre/osp/<fsname>-MDTxxxx-osp-MDTxxxx/rpc_stats`  
  例如：`cat /sys/kernel/debug/lustre/osp/nas_test-MDT0001-osp-MDT0000/rpc_stats`

# q
Lustre的rpc_stats输出中，“pages per rpc”分布表示什么？
# a
“pages per rpc”分布记录了每次RPC请求聚合的页面数量（pages per rpc）及其占比和累计占比。例如：
```
pages per rpc         rpcs   % cum %
1:                       0   0   0
16:                      2  22  88
```
该分布反映I/O合并效率：数值越大表明更多页面被合并到一个RPC中传输，有助于减少网络交互次数。

# q
Lustre的rpc_stats输出中，“rpcs in flight”分布表示什么？
# a
“rpcs in flight”分布记录了同时处于飞行状态（正在处理）的RPC请求数量及其占比。例如：
```
rpcs in flight        rpcs   % cum %
1:                       3  33  33
2:                       2  22  55
6:                       1  11 100
```
它体现客户端的RPC并发度。如果大量RPC聚集在较高并发数，说明系统正在并行处理大量请求。

# q
Lustre的rpc_stats输出中，`modify_RPCs_in_flight`字段有什么作用？
# a
`modify_RPCs_in_flight`表示当前正在进行的元数据修改RPC数量。值为0表明没有正在处理的修改请求，非0则说明有修改操作正在执行。该指标对监控元数据操作负载和客户端请求排队情况有参考意义。

