# q
起快照策略并持续打流量时，OSD内存异常增大的典型根因是什么？
# a
在segment删除时未及时释放block group内存，以及block_num为0时未释放block_mapping内存，导致内存泄漏。

# q
如何从OSD日志或系统信息中定位快照相关的内存泄漏问题？
# a
持续观测OSD使用的内存信息，执行命令：
```bash
cat /engine-fs/osd/*/osd_info | grep mem
```
观察内存使用量是否持续异常增大（无正常波动）；若旧版本OSD内存持续增长而新版本有上下波动，则可判定存在泄漏。

# q
解决OSD快照删除内存泄漏的标准修复流程包含哪些关键改动？
# a
1. 在segment删除时直接释放block group内存。
2. 在block_num为0时释放block_mapping内存。
3. 修改核心文件：`src/osd/osd.c`、`src/osd/osd.h`、`src/osd/osd_mdlog.c`（Commit: 4c673a570ca），并补充修改 `src/osd/osd_mdlog.c`（Commit: 372b8896dd89）。

