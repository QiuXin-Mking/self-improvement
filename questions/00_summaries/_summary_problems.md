# q
单桶对象数量过大（~970万）导致 list 操作极慢的典型根因是什么，如何解决？
# a
根因：bucket index shard 数量不足（如仅32个），导致大量对象集中在少量 shard 上，list 操作严重堵塞。
解决：增大 `rgw_max_dynamic_shards` 参数，必要时执行 bucket reshard 增加 shard 数量。关键命令：
```bash
radosgw-admin bucket reshard --bucket upload --num-shards 8
rados -p obj_index_xxx stat .dir.xxx
```

# q
如何从日志定位 OSD 因心跳超时导致 down 的问题？
# a
在 OSD 日志中搜索 `heartbeat_map is_healthy had timed out after 30`，判断是否存在心跳超时。同时通过以下命令检查集群状态和开启调试：
```bash
ceph osd tree
ceph daemon osd.X config set debug_osd 20
grep "heartbeat_map is_healthy had timed out" ceph-osd.X.log
```

# q
在 bcache 设备环境下，如何重建 OSD 设备、bcache 与底层磁盘的完整映射关系？
# a
通过以下步骤建立映射链：
1. 使用 `docker inspect ceph_osd_xxx | grep OSD` 查看容器 OSD_DEVICE。
2. 使用 `bcachectl cachelist` 获取 bcache 缓存设备列表。
3. 使用 `pvs; vgs; lvs` 和 `lsblk` 检查 LVM 层。
4. 使用 `ll /dev/disk/by-partuuid/` 通过 partuuid 关联设备。
综合以上信息即可重建 OSD → bcache → 物理盘/VG/LV 的完整对应关系。

# q
Ceph 扩容后出现大量 slow ops，如何通过命令行定位 delayed ops 和受影响的 PG？
# a
使用以下命令定位：
```bash
ceph daemon osd.X dump_ops_in_flight      # 查看该 OSD 上排队的 op
ceph pg dump | grep undersized            # 查找处于 undersized 状态的 PG
ceph tell osd.X injectargs '--debug_osd 20'  # 临时开启 debug 日志分析
```
重点关注 `recovery_wait+undersized+degraded+remapped` 的 PG。

# q
RGW 生命周期（LC）过期规则卡在 process 状态不删除对象，解决的标准流程是什么？
# a
步骤：
1. 查看 LC 状态：`radosgw-admin lc list`
2. 强制推进 LC 处理：`radosgw-admin lc process`
3. 若锁定导致卡死，执行修复：`radosgw-admin lc reshard fix`
4. 必要时查看 LC 日志命名空间：`rados -p .product.rgw.log --namespace=lc ls`
5. 使用带 debug 的 LC 命令：`CEPH_ARGS="--debug-rgw=20" radosgw-admin lc process > lc_process.log 2>&1`

