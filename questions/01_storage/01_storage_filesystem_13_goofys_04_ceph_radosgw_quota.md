# q
如何对指定用户和桶设置存储配额？
# a
使用 `radosgw-admin quota set` 命令：
```bash
radosgw-admin quota set --uid=<用户ID> --bucket=<桶名> --quota-scope=bucket --max-size=<字节数> --max-objects=<对象数量>
```
例如：
```bash
radosgw-admin quota set --uid=qx1 --bucket=qiuxinbucket1 --quota-scope=bucket --max-size=$((1*1024)) --max-objects=20
```

# q
`quota-scope` 参数有哪些取值？分别代表什么？
# a
- `user`：对用户级别的总配额进行控制，限制该用户所有桶的资源使用总量。
- `bucket`：对单个桶级别的配额进行控制，限制指定桶的资源使用上限。

# q
设置桶配额后如何使其立即生效？
# a
使用 `radosgw-admin quota enable` 命令，需指定对应的用户、桶和作用域：
```bash
radosgw-admin quota enable --uid=<用户ID> --bucket=<桶名> --quota-scope=bucket
```
示例：
```bash
radosgw-admin quota enable --uid=qx1 --bucket=qiuxinbucket1 --quota-scope=bucket
```

