# q
为什么 `ceph config get osd.157 osd_op_num_shards` 和 `ceph config show osd.157 | grep osd_op_num_shards` 对同一参数返回值不一致？
# a
Ceph 配置存在多个层次（全局、守护进程类型、守护进程实例），不同命令读取的配置来源不同。  
- `ceph config get` 返回指定级别的最终生效值，会综合所有层次。例如 `ceph config get osd.157 osd_op_num_shards` 返回 16，说明该 OSD 实例被单独设置过此参数。  
- `ceph config show` 可能仅显示由 monitor 持久化存储的某个特定来源的配置，如案例中返回 10、来源为 `file mon`，说明 mon 侧文件配置仍为旧值。  
典型根因：实例级配置覆盖了类型或全局配置，且 monitor 存储的配置未同步更新，导致不同命令视角下的值不一致。

# q
如何定位 Ceph 配置参数在不同层级的具体来源与覆盖关系？
# a
通过逐级查询配置值来对比差异。示例排查命令：  
```bash
# 查询全局级别
ceph config get global osd_op_num_shards
# 查询 OSD 类型级别
ceph config get osd osd_op_num_shards
# 查询具体实例级别
ceph config get osd.157 osd_op_num_shards
# 查询 monitor 侧配置（非预期来源）
ceph config get mon osd_op_num_shards
```
也可直接使用 `ceph config dump` 查看所有配置项及其来源，快速定位覆盖关系。案例中输出显示只有 `osd.157` 返回 16，其余级别返回 0，表明值来自实例级覆盖，mon 配置文件中残留的值为 10。

