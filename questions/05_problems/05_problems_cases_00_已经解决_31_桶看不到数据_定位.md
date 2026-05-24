# q
如何通过radosgw-admin确认桶的实际数据量和对象数量？
# a
使用 `radosgw-admin bucket stats --bucket=<bucket-name>` 命令，返回桶的详细统计，包括 `num_objects`（对象总数）和各存储使用量 `size`、`size_actual`、`size_utilized`。例如：
```bash
radosgw-admin bucket stats --bucket=mysql-backup
```
输出中 `num_objects` 为 196004，`rgw.total` 的 `size` 约为 1.5 TB，表明数据确实存储在集群中。

