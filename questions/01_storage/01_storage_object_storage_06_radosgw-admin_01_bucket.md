# q
radosgw-admin bucket check 命令的作用是什么，如何让它自动修复问题？
# a
`radosgw-admin bucket check --bucket <bucket_name>` 用于检查指定桶的索引一致性。添加 `--fix` 选项后，命令会在检查的同时自动重建不一致的索引。
```
radosgw-admin bucket check --bucket <bucket_name> --fix
```

# q
如何在使用 radosgw-admin 时开启详细的调试日志？
# a
通过设置环境变量 `CEPH_ARGS` 可以临时开启调试日志。例如：
```
CEPH_ARGS="--debug-rgw=20 --debug-ms=1" radosgw-admin bucket check --bucket upload --fix
```
其中 `--debug-rgw` 设置 RGW 日志级别，`--debug-ms` 设置消息层日志级别，值越大越详细。

# q
如何通过 ceph daemon 命令查看与设置 RGW 的运行时配置（以 LC 参数为例）？
# a
使用 `ceph daemon <admin_socket> config get <参数名>` 查看配置，使用 `config set <参数名> <值>` 动态修改。常见与生命周期（LC）相关的参数包括：
- `rgw_lc_max_objs`: 存储生命周期索引的数据对象数量；
- `rgw_lc_max_wp_worker`: 每个 LCWorker 内部的工作线程数；
- `rgw_lc_max_worker`: LC 工作线程总数；
- `rgw_lc_thread_delay`: LC 线程处理延迟。

示例：
```
ceph daemon /var/run/ceph/client.rgw.ees-stxx-176-2-cxsj.asok config get rgw_lc_max_objs
ceph daemon /var/run/ceph/client.rgw.ees-stxx-176-2-cxsj.asok config set rgw_lc_thread_delay 0
```

