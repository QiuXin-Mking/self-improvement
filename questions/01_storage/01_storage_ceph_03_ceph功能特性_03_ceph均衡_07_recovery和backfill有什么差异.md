# q
Ceph中的recovery（恢复）是什么？它的典型触发场景是怎样的？
# a
recovery用于修复丢失或损坏的PG副本，保证数据完整性。典型场景：一个三副本PG的某个OSD掉线一段时间后重新上线，期间只有其他副本发生了数据变更，此时该OSD需要从健康OSD对比并补回缺失的变化数据块。

# q
Ceph中的backfill（回填）是什么？它的典型触发场景是怎样的？
# a
backfill用于OSD加盘后迁移数据，实现存储均衡。典型场景：添加新OSD后，CRUSH map将部分PG重新分配到新OSD，此时需要将PG内的全部数据对象分批拷贝到新盘（新成员），旧OSD上的副本在backfill完成后才会删除。

# q
recovery和backfill的主要区别是什么？
# a
recovery是修复丢失/损坏的副本，目标OSD仅需补回变化的数据块；backfill是向新加入的空OSD全量拷贝PG数据，实现存储均衡。recovery针对部分数据差异，backfill是整体数据迁移。两者机制不同，都会影响集群性能。

