# q
如何查看Ceph集群的当前状态？
# a
使用以下命令查看Ceph集群状态：
```bash
ceph -s
```
该命令会输出集群健康状态、监视器信息、OSD、PG等综合概况。

# q
如何快速获取Ceph集群的健康状态？
# a
使用以下命令直接获取健康状态信息：
```bash
ceph health
```
它会返回 HEALTH_OK、HEALTH_WARN 或 HEALTH_ERR 等简要健康状态。

# q
如何实时监控Ceph集群事件与状态变化？
# a
使用监控模式命令：
```bash
ceph -w
```
该命令会持续输出集群变化信息，类似 `tail -f` 的效果，按 `Ctrl+C` 退出。

# q
如何以不同调试级别实时监控Ceph集群？
# a
可以指定调试级别进行实时监控：
```bash
ceph --watch-debug -w
ceph --watch-info -w
```
`--watch-debug` 会输出调试级别的事件，而 `--watch-info` 输出信息级别的事件，用于更细粒度的监控。

