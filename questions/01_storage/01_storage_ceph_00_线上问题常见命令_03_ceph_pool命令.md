# q
如何创建Ceph pool？
# a
使用以下命令创建pool，需指定pool名称和PG数量：
```bash
ceph osd pool create your_new_pool_name your_pg_num
```

# q
如何删除Ceph pool？
# a
使用以下命令删除pool，需要提供pool名称两次并加上确认标志：
```bash
ceph osd pool rm my_pool my_pool --yes-i-really-really-mean-it
```

# q
PG数量的设置如何影响Ceph pool的负载均衡和故障恢复？
# a
PG数量越小，每个PG包含的数据对象越多，可能导致数据分布不均匀，某些OSD负载较重，影响性能，因此具有较小PG数量的pool权重较低。相反，PG数量较大的pool会将数据分成更多PG，分布更均衡，负载均衡性和故障恢复效率更高，权重较高。PG数量的设置需要根据集群规模、硬件配置和性能需求合理规划调整，较大的pool应设置较大的PG数量以确保均衡分布和高效恢复。

# q
如何修改已有Ceph pool的PG数量？
# a
使用以下命令修改pool的PG数量（示例中将qxpool的pg_num设置为320）：
```bash
ceph osd pool set qxpool pg_num 320
```

