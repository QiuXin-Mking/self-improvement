# q
卷容量查询不到（一直为0）的典型根因是什么？
# a
典型根因是 etcd 中缺少容量统计专用键 `/volumes/<vol_id>/cap`。正常环境中该键会存储 `totalcap` 和 `usedcap`，问题环境中仅有卷基本信息中的 `used_cap` 字段且恒为 0，容量统计信息未同步至 etcd。

# q
如何从 etcd 日志定位卷容量查询不到的问题？
# a
在 etcd 节点执行以下命令检查目标卷的键值结构：
```bash
etcdctl get /volumes/<卷ID> --prefix
```
观察是否存在 `/volumes/<卷ID>/cap` 键。如果缺失（如问题日志中只有 `used_cap:0` 而无 `/cap` 子键），说明容量统计数据未上报。正常参考输出应包含：
```json
/volumes/1/cap
{"totalcap":"4194304","usedcap":"5184424368"}
```

# q
解决卷容量查询不到的标准排查流程是什么？
# a
1. 使用 `etcdctl get /volumes --prefix` 获取所有卷信息，确认目标卷的 etcd 键树中是否缺少 `/volumes/<ID>/cap`。
2. 检查容量统计相关服务（如 mdbs 的数据统计模块）是否正常运行，查看组件日志是否有上报失败的错误。
3. 核对卷的 `used_cap` 字段是否为初始值 0，辅证容量统计未生效。
4. 若服务异常，将其恢复后观察 `/volumes/<ID>/cap` 是否生成并更新。

