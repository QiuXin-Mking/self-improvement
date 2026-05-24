# q
Ceph PG恢复失败、日志中出现“missing primary copy of”错误的典型根因是什么？
# a
典型根因是主OSD上的PG缺少主副本（primary copy），通常由于OSD进程异常、磁盘 I/O 故障或数据损坏导致 PG 无法正常恢复，使 PG 长期处于 `active+undersized+degraded+remapped+backfilling` 状态。

# q
如何通过日志定位Ceph PG缺少主副本导致恢复卡住的问题？
# a
执行 `ceph -w` 观察集群事件，查找包含 `missing primary copy of` 的 ERR 日志，例如：
```
2025-09-18T10:30:48.107610+0800 osd.10 [ERR] 4.157 missing primary copy of 4:eae490b5:::rbd_data.1bfefcc348fd86.000000000001e7d7:d8c, will try copies on 26
```
该日志表明 PG `4.157` 在 `osd.10` 上缺少主副本，集群尝试从 `osd.26` 拷贝。结合 `ceph pg dump | grep ^4.157` 可查看 PG 详细状态与 acting set。

# q
解决Ceph PG因缺少主副本而无法恢复的标准流程是什么？
# a
1. 定位问题 PG 和异常 OSD：  
   ```bash
   ceph pg dump | grep ^<pgid>
   ```
2. 通过 RBD 数据前缀（如 `1bfefcc348fd86`）查找受影响的镜像：  
   ```bash
   for img in $(rbd ls -p vms); do
     if rbd info vms/"$img" | grep -q 1bfefcc348fd86; then
       echo "$img"
     fi
   done
   ```
3. 确认镜像快照情况：  
   ```bash
   rbd snap ls vms/<image>
   ```
4. 重启异常的 OSD 服务（本例中为 `osd.10`）：  
   ```bash
   systemctl restart ceph-osd@10
   ```
5. 监控恢复进度：  
   ```bash
   ceph -w
   ```
   等待 PG 状态恢复为 `active+clean`。

