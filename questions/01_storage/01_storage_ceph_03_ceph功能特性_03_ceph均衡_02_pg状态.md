# q
Ceph PG 的健康状态是什么？
# a
active + clean，表示 PG 处于活跃且数据一致的状态。

# q
active + clean + scrubbing 状态表示什么？
# a
表示 PG 正在进行浅度清理（scrubbing），即集群当前未执行深度清理（deep scrub）操作。

