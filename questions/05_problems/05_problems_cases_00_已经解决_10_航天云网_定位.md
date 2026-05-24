# q
如何通过命令快速发现Ceph集群中存在异常的OSD？
# a
使用以下命令查看OSD树，状态为 `down` 即为异常 OSD：
```bash
docker exec ceph_mon ceph osd tree
```
输出示例中 `osd.4` 状态为 `down`：
```
 4    ssd    7.31740          osd.4          down         0  1.00000
```
结合 `docker exec ceph_mon ceph -s` 可以获取集群整体健康状态。

# q
OSD down 且日志中出现“Unexpected IO error”的典型根因是什么？
# a
通常指向**硬件问题**，例如磁盘故障、控制器异常或连接问题。关键日志片段：
```
/mnt/ceph-15.2.10.3/.../KernelDevice.cc: 575: ceph_abort_msg("Unexpected IO error. This may suggest a hardware issue. Please check your kernel log!")
```
此时应检查对应主机的内核日志（`dmesg`、`/var/log/messages`）中的磁盘 I/O 错误，以确认磁盘或链路是否损坏。

