# q
RGW生命周期处理无法清理过期未完成multipart上传的典型根因是什么？
# a
bucket index统计信息严重不一致，`rgw.multi_uncomplete`记录了海量未完成分片对象（例如298378个），而实际简单对象仅305个。这种巨大差异导致生命周期（LC）处理遍历索引条目时陷入循环或始终处于sleep状态，无法正常完成过期multipart的清理。

# q
如何从日志定位RGW LC处理multipart过期删除失败的问题？
# a
使用调试日志复现：`CEPH_ARGS="--debug-rgw=20" radosgw-admin lc process`，在输出中搜索`cls_bucket_list_unordered`。若日志反复出现相同的`start_after` multipart键前缀，且每次仅步进少量条目（如`_multipart_18370244178/20250710113148868.7z.2~…`），表明LC在遍历海量无效索引项时形成性能瓶颈或逻辑阻塞，无法推进至需要清理的时间窗。

# q
解决RGW bucket中大量未完成multipart导致LC失效的标准流程是什么？
# a
1. 确认统计差异：`radosgw-admin bucket stats --bucket <bucket>`，关注`rgw.multi_uncomplete.num_objects`与`list-objects-v2`结果的差值。
2. 修复索引：`radosgw-admin bucket check --bucket <bucket> --fix`以校正bucket index头部统计。
3. 调整LC配置：增大`rgw_lc_max_worker`（如10）和`rgw_lc_max_objs`（如64），并设置`rgw_lifecycle_work_time = "00:00-23:59"`全时段工作。
4. 强制触发LC处理：`radosgw-admin lc process`或`timeout 1200 env CEPH_ARGS="--debug-rgw=20" radosgw-admin lc process`后台运行。
5. 必要时处理分片：`radosgw-admin lc reshard fix`。
6. （可选）手动清理：通过`radosgw-admin bucket radoslist --bucket <bucket>`列出所有multipart对象，确认后可执行`rados -p <pool> rm <object>`删除。

