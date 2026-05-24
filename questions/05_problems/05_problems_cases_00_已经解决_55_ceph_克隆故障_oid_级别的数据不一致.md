# q
如何通过导出PG对象列表并 diff 排查 Ceph OSD 间 oid 级别数据不一致？
# a
先从各 OSD 上导出指定 PG（如 4.112）的对象列表并排序，例如：
```bash
rados -p <pool> ls > 4.112_osd0
sort 4.112_osd0 > 4.112_osd0_sort
```
对每个 OSD 重复操作，然后两两对比：
```bash
diff 4.112_osd0_sort 4.112_osd10_sort
diff 4.112_osd0_sort 4.112_osd16_sort
```
diff 输出中 `<` 表示仅在左侧 OSD 存在的对象，`>` 表示仅在右侧 OSD 存在的对象，从而直接定位不一致的 oid。

# q
Ceph 克隆（snapshot）操作引发的 oid 级别不一致在 diff 结果中有哪些典型特征？
# a
典型特征为与快照相关的对象仅在部分 OSD 上出现。例如：
```diff
< ["4.112",{"oid":"rbd_data.1ba309a4e274f1.000000000001ced9","key":"","snapid":3453,...}]
< ["4.112",{"oid":"rbd_data.1ba309a4e274f1.000000000001ced9","key":"","snapid":3461,...}]
```
在 osd0 与 osd10 对比中，osd10 多出 `snapid=3453` 和 `snapid=3461` 的对象，而 osd0 上缺失；同时存在 `snapid=-2` 的 head 对象在部分 OSD 上缺失或多余的情况，说明克隆期间对象复制未完成或出现分裂。

