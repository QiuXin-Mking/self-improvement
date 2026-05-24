# q
如何查找Ceph集群中处于unfound状态的PG？
# a
使用命令 `ceph pg dump | grep unfound` 或 `ceph pg dump | grep backfill_unfound` 排查。

# q
解决Ceph unfound PG问题的标准修复流程是什么？
# a
首先尝试对指定PG执行修复：`ceph pg repair <pgid>`（例如 `ceph pg repair 5.18c`）。该命令会指示对应OSD进行修复，输出类似 `instructing pg 5.18c on osd.101 to repair`。

# q
当常规repair无法恢复unfound对象时，可使用什么最终手段？
# a
使用 `ceph pg <pgid> mark_unfound_lost revert`（例如 `ceph pg 5.18c mark_unfound_lost revert`）。注意：该操作会丢弃unfound对象，可能导致数据丢失，需谨慎使用。

