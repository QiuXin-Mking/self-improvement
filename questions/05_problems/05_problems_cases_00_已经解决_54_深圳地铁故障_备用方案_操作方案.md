# q
如何查看Ceph集群中处于down状态的PG？
# a
执行以下命令过滤出状态为down的PG：
```bash
ceph pg dump | grep down
```

# q
在Ceph故障恢复中，如何临时停止自动恢复和回填，并强制创建down状态的PG？
# a
标准操作流程：
1. 查看down PG：
   ```bash
   ceph pg dump | grep down
   ```
2. 停止自动恢复与回填：
   ```bash
   ceph osd set norecover
   ceph osd set nobackfill
   ```
3. 强制创建指定PG（以pgid为例）：
   ```bash
   ceph pg force_create_pg <pgid>
   ```
4. 检查集群状态：
   ```bash
   ceph -s
   ```
操作完成后需根据情况使用 `ceph osd unset norecover` 和 `ceph osd unset nobackfill` 恢复自动恢复。

# q
如何查看指定OSD上承载的PG详情？
# a
使用 `ceph daemon` 命令直接查询OSD：
```bash
ceph daemon osd.<id> show | grep pg
```
将 `<id>` 替换为实际的OSD编号（如 0），输出结果中会包含该OSD上的PG列表及其状态。

# q
排查Ceph OSD进程异常时，如何获取进程的父进程及cgroup信息？
# a
1. 查看服务状态：
   ```bash
   systemctl status <PID>
   ```
2. 获取cgroup路径：
   ```bash
   cat /proc/<PID>/cgroup
   ```
3. 查看进程树及父进程PID：
   ```bash
   pstree -p <PID>
   # 或
   ps -ef | grep <PID>
   ```
替换 `<PID>` 为实际的OSD进程ID。

