# q
Ceph集群扩容时为什么会引发数据恢复和重平衡？
# a
扩容（增加新OSD、节点等）会改变集群的数据分布，Ceph需要重新放置数据以确保数据在新配置下的**一致性和冗余度**（如副本的完整性），因此自动触发数据恢复（recovery）和重平衡（rebalance）操作。

# q
如何通过参数调整Ceph OSD的恢复并行度？osd_recovery_max_active和osd_max_backfills的区别是什么？
# a
通过以下命令调整恢复并行度：
```sh
ceph config set osd osd_recovery_max_active 5   # 每个OSD同时进行的最大恢复操作数
ceph config set osd osd_max_backfills 2         # 每个OSD同时进行的最大回填操作数
```
区别：
- `osd_recovery_max_active` 控制的是**恢复**（recovery）操作的并发数，主要处理因副本不一致造成的增量数据修复。
- `osd_max_backfills` 控制的是**回填**（backfill）操作的并发数，主要处理完整的PG数据迁移（如新增OSD后整体数据重新放置）。

# q
osd_recovery_op_priority和osd_recovery_sleep参数的作用是什么？
# a
- **osd_recovery_op_priority**：设置恢复/回填操作的优先级（数值越高优先级越高）。高优先级可加速恢复，但可能抢占客户端I/O资源，影响正常业务性能。示例：  
  ```sh
  ceph config set osd osd_recovery_op_priority 10
  ```
- **osd_recovery_sleep**：控制每次恢复操作之间的等待时间（秒）。增大该值可减少恢复操作对客户端I/O的冲击；设为0或较小值则恢复更激进。示例：  
  ```sh
  ceph config set osd osd_recovery_sleep 0.1
  ```

# q
如何暂停和恢复Ceph集群的数据回填与恢复操作？
# a
暂停回填和恢复（常用于维护窗口）：
```bash
ceph osd set nobackfill
ceph osd set norecover
```
恢复回填和恢复：
```bash
ceph osd unset nobackfill
ceph osd unset norecover
```

