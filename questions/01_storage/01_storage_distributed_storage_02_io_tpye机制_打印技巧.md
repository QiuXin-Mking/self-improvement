# q
`LOG_DBG_MSG` 宏在存储系统调试中的典型用法是什么？
# a
典型用法示例如下：
```c
LOG_DBG_MSG("%s process##qx blk_in_osd:%u cmd_begin:%lu", 
            ocache_io->pri_req,
            (char *)(__FUNCTION__),
            blk_in_osd,
            ocache_cmd->cmd_begin);
```
该宏用于记录 IO 请求处理的关键状态，其中 `%s` 打印请求标识 `ocache_io->pri_req`，`process##qx` 为固定流程标记，`%u` 打印 OSD 块索引 `blk_in_osd`，`%lu` 打印命令开始时间/序列 `ocache_cmd->cmd_begin`。

# q
在缓存与刷新链路中，建议在哪些关键函数插入打印日志？
# a
建议在 `ocache_cmd_done_blk_refresh`（命令完成且块刷新时）追踪数据块刷新状态，确保落盘一致性；以及在 `CACHE_UPDATE` 事件标记缓存变更，用于调试缓存命中与淘汰策略。

# q
在日志回放与恢复链路中，有哪些关键打印点？
# a
关键打印点包括：
- `clog_load_active_space_log_rsp`：活跃空间日志加载响应，确认加载成功与否；
- `ocache_clog_replay_update`：提交日志回放更新，追踪 Commit Log 回放进度；
- `clog_block_map_replay_update`：块映射回放更新，追踪块映射表恢复过程。

# q
添加调试日志时，如何控制日志长度以避免关键信息被截断？
# a
应遵守长度限制原则：
- 避免打印大段内存数据（Hex Dump），除非必要；
- 使用缩写或哈希值代替长字符串标识；
- 确保单行日志保持在合理范围内，建议小于 1024 字节。

