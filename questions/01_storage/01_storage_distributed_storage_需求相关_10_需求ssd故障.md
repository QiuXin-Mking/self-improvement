# q
SSD故障后，修复流程是由什么事件触发的？
# a
adapter发现SSD下线后，通知ocache mgt开始修复。

# q
SSD故障修复开始时，需要执行哪些关键操作？
# a
1. 暂停刷盘（`OCACHE_FLUSH_PAUSE`）
2. 通知mdlog切换到osd（`MDLOG_SHIFT_OSD`）
3. 限制osd的io下发（`OMT_IO_PAUSE`）
4. 设置ocache_io模式：禁止刷盘启动、取消下发读SSD、取消下发写、所有写设置成写穿、写回命令重新分配cache

# q
修复过程中如何区分`ocache_io_A`和`ocache_io_B`两种场景？
# a
`ocache_io_A`：元数据不丢失；`ocache_io_B`：元数据丢失。通过检查元数据是否脏以及对应脏位置的元数据在内存中是否存在来进行区分。

# q
修复完成后，osd需要经历哪些关键步骤才能恢复正常服务？
# a
1. 在无写流量和修复流量期间重建alloc bitmap（`OMT_REBUILD_BITMAP`）
2. 修复完成，将osd数据设置成有效状态
3. 解除io限制，允许osd的io正常下发（`OMT_IO_RESUME` / `OMT_IO_NOT_PAUSE`）

# q
该流程中涉及了哪些新增或改动的消息类型？
# a
新增type：`OCACHE_REPAIRE_START`、`OCACHE_FLUSH_PAUSE`、`OMT_IO_PAUSE`、`OMT_FREE_SEGMENT`、`OMT_REPAIR_SEG_META`、`OMT_META_LOST`、`OMT_REBUILD_BITMAP`等。紫色为改动流程，黄色为固有master流程。

