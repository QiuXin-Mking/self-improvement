# q
删除快照操作会立即触发底层存储段的回收（osd_process_recycle）吗？
# a
不会。仅删除快照时，etcd 中的 ds 条目以及 OSD 日志均无变化。只有当从回收站彻底删除卷时，才会触发 `osd_process_recycle`，日志中会出现 `osd process recycle handle` 和 `osd process delete segment` 等回收行为。

# q
在测试环境中，etcd 中 `/ds/{id}` 下的 `state` 字段与 `ref` 字段分别代表什么含义？
# a
- `state`: `0` 表示快照（只读），`1` 表示活动卷（可写）。
- `ref`: 表示该 ds 被依赖的计数，通常快照 ds 的 ref 为 0，而活动卷的 ref 可能为 1（表示被快照依赖）。

# q
从回收站删除卷后，合并快照过程中 ds 回收的现象是什么？
# a
回收站删除卷后，触发 OSD 回收。如果存在快照合并（ds 融合），被依赖的父 ds（如 ds3）的 segment 会被一并回收，etcd 中对应的 ds 条目会被移除，`osd_info` 中的 segment num 减少，并输出 `Segment sl delete` 等日志。

