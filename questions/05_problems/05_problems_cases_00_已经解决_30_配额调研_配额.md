# q
如何用 radosgw-admin 为指定用户设置桶级别的存储配额（限制最大容量和对象数）？
# a
使用以下命令，将 uid 和 bucket 替换为实际值：
```bash
radosgw-admin quota set --uid=example-user --bucket=example-bucket --quota-scope=bucket --max-size=$((1 * 1024 * 1024 * 1024)) --max-objects=500
```
该命令将 example-user 在 example-bucket 桶上的配额设为最大 1 GiB 容量和 500 个对象。

# q
如何为 Ceph 对象存储用户设置用户级全局配额？
# a
使用以下命令，将 uid 替换为实际用户名：
```bash
radosgw-admin quota set --uid=example-user --quota-scope=user --max-size=$((1 * 1024 * 1024 * 1024)) --max-objects=500
```
该命令设置 user 级别的配额，限制该用户在所有桶中使用的总容量和总对象数。

# q
设置完配额后，如何让配额对某个用户的特定桶生效？
# a
使用 `radosgw-admin quota enable` 命令启用配额：
```bash
radosgw-admin quota enable --uid=qx1 --bucket=qiuxinbucket1 --quota-scope=bucket
```
该命令针对用户 qx1 的 qiuxinbucket1 桶启用之前设定的桶级别配额限制。

