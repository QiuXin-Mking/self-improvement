# q
Ceph集群中OSD全部down且节点已不可用的典型根因是什么？
# a
节点被意外关机或该节点上的Ceph RPM包被全部卸载，导致OSD守护进程无法运行，所有归其管理的OSD均显示为down状态。

# q
如何排查Ceph PG中对象丢失（unfound）的问题？
# a
使用命令 `ceph pg <pgid> query` 查看PG的详细信息，确认是否存在 `unfound` 对象。如果确认这些对象无法从任何副本恢复，可执行 `ceph pg <pgid> mark_unfound_lost revert` 将PG回退到上一个可用版本，或使用 `delete` 参数彻底丢弃丢失对象。注意：此操作会导致对应时间点后的数据丢失，需谨慎使用。

# q
当Ceph OSD所在主机被卸载但其底层LV仍然存在时，恢复数据的标准流程是什么？
# a
1. 确认底层逻辑卷完好，未被格式化或删除。
2. 如果原节点可恢复，重装Ceph软件包并重新激活OSD（如 `ceph-volume lvm activate`）。
3. 若原节点无法恢复，可将该磁盘挂载到其他节点，使用 `ceph-objectstore-tool` 导出数据。
4. 对于已经标记为 `unfound` 且无法修复的PG，只能使用 `ceph pg <pgid> mark_unfound_lost revert` 或 `delete` 强制完成PG状态恢复，以恢复集群可用性。

