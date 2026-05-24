# q
`ceph osd rm osd.x` 和 `ceph osd crush rm osd.x` 命令在 Ceph 集群中分别起什么作用？
# a
- `ceph osd rm osd.x`：从 Ceph 集群中彻底删除指定 OSD（包括其认证和 ID）。
- `ceph osd crush rm osd.x`：从 CRUSH map 中移除该 OSD 的条目，使其不再参与数据分布计算。

# q
为什么说 `ceph osd rm` 是为缩容设计，而不是为临时换盘设计的？
# a
因为 OSD 被 `rm` 删除后，存储池的 PG 映射会发生改变，导致大量数据重新分布和迁移。如果是临时坏盘换盘，短时间内替换损坏的 OSD 可以保持原有 PG 映射，避免不必要的数据迁移，从而降低对业务的影响。所以缩容时才应移除 OSD，而换盘时不应删除，应等待替换后让集群自行恢复。

