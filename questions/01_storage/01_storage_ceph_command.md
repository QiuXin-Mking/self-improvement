# q
如何转储 Ceph 集群中 osd.17 的历史操作按持续时间排序的详细信息？
# a
使用 `ceph daemon` 子命令直接与 OSD 守护进程通信，执行以下命令：
```bash
ceph daemon osd.17 dump_historic_ops_by_duration
```

# q
`ceph daemon osd.17 dump_historic_ops_by_duration` 这条命令的主要用途是什么？
# a
该命令用于从指定 OSD (osd.17) 获取历史操作记录，并按操作持续时间（duration）进行排序和转储。通常用于在不停止 OSD 的情况下快速分析慢请求或性能瓶颈。

