# q
分布式存储OSD容量显示异常的典型根因是什么？
# a
通常与分配位图(alloc_bitmap)的 `alloc_num` 计数不一致有关，该值未能正确反映已分配的块数，进而导致 `totalSize`、`usedSize` 等统计信息错误。相关排查参数为 `&osd->map.alloc_bitmap->alloc_num` 和 `attr_fs_service_osd_info_show`。

# q
如何通过命令检查OSD的详细容量信息以定位显示异常？
# a
执行 `cat /engine-fs/osd/<osd_id>/osd_info` 查看OSD内部状态，重点关注以下字段：
- `total_blkNum`：总块数
- `blockAllocNum`：已分配块数
- `usedSize`：已用容量
正常情况下，`blockAllocNum` 应与实际分配一致，`usedSize` 不应为空或与分配情况矛盾。若出现 `blockAllocNum` 为零但 `usedSize` 不为零等不一致，即为异常。

# q
在“不打流量、无卷”的情况下，OSD的 `osd_info` 中哪些值应为零？
# a
此时应无任何空间占用，关键字段均为零：
- `usedSize: 0KB`
- `blockAllocNum: 0`
- `dirty mem: 0 >> 0MB`
- `seg used mem: 0 >> 0MB`
示例输出片段：
```
totalSize:     291620864KB
usedSize:      0KB
total_blkNum:  4556560
blockAllocNum: 0
```

