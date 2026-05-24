# q
如何查看 Ceph 集群中 OSD 的数量与基本状态？
# a
使用以下命令查看 OSD 数量以及 up/in 状态，以及当前的 epoch：
```bash
ceph osd stat
```
示例输出：`4 osds: 3 up (since 23m), 3 in (since 13m); epoch: e345`  
OSD 状态含义：`in` 表示在集群内，`out` 表示在集群外，`up` 表示进程正在运行，`down` 表示进程已停止。

# q
如何将某个 OSD 手动标记为 down 或 up？
# a
将 OSD（例如 ID 为 0）标记为 down，使其不接受读写请求但进程仍存活（变为 down in 状态）：
```bash
ceph osd down 0
```
将已 down 的 OSD（例如 ID 为 3）重新标记为 up：
```bash
ceph osd up 3
```

# q
如何暂停和恢复整个 Ceph 集群的数据写入？
# a
暂停集群 IO，集群将不再接受任何数据：
```bash
ceph osd pause
```
恢复集群，重新接收数据：
```bash
ceph osd unpause
```

# q
如何查看 OSD 的请求延时性能统计？
# a
使用以下命令查看所有 OSD 的平均延时（主要用于排查单块磁盘问题）：
```bash
ceph osd perf
```
统计中：`commit_latency` 是从接收请求到设置 commit 状态的时间间隔，`apply_latency` 是从接收请求到设置 apply 状态的时间间隔。如有异常应及时剔除有问题的 OSD。

# q
如何从 CRUSH map 中移除一个 OSD？
# a
使用以下命令从 CRUSH map 中移除指定 OSD（例如 `osd.3`）：
```bash
ceph osd crush rm osd.3
```

