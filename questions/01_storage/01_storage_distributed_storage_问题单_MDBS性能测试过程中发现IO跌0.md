# q
MDBS性能测试中发现IO跌0的典型根因是什么？
# a
典型根因是 **ocache（对象缓存）丢IO**，导致写入的IO未被正确处理而丢失。具体表现为在clog回放阶段，由于 cache id 冲突（相同的 block id 被重复分配）引发 `Failed to set cache btmp for same cacheid` 错误，或触发 `blk info cache id[...] is conflict to clog cache id[...]` 冲突打印，进而导致 block 释放流程异常，最终造成 IO 跌 0。

# q
如何从日志定位MDBS IO跌0是否由ocache丢IO引起？
# a
可通过以下日志关键字和统计信息定位：
- 检查系统日志或 dmesg 中是否出现：
  ```
  !!!!!!!!!!set bit was wrong,
  Failed to set cache btmp for same cacheid, oid:
  ```
- 检查是否有 cache id 冲突打印：
  ```
  blk info cache id[...] is conflict to clog cache id[...], blk id oid:...
  ```
- 对比 `/engine-fs/ocache/*/` 目录下 `send_io_rsp`（发给 osd 的响应）与 `recv_io_req`（接收的请求）计数差异：如果 osd 发出的 IO 数量与 ocache 接收的 IO 数量存在数量级不一致（如差十几个 IO），表明 ocache 丢失了部分 IO。
- 对比 osd 层面 `send_io_rsp`（发给 sio）和接收的请求计数，确认丢IO发生在 osd 与 ocache 之间。

# q
解决MDBS因ocache丢IO导致IO跌0的标准处理流程是什么？
# a
1. **快速恢复**：在集群所有节点执行服务重启：
   ```
   python /opt/macrosan/mdbs/web/restart_services.py cluster restart all
   ```
   重启后可暂时恢复IO，但可能再次触发问题。
2. **版本检查**：确认是否使用了“切换spdk，1M不拆分io”的版本，该版本已知存在ocache丢IO缺陷，需回退至稳定版本。
3. **代码修复**：由开发侧解决 cache id 分配冲突问题，优化 clog 回放逻辑，并修正 `ocache_try_free_block` 释放条件（需无 cache、状态可读、无命令且不在等待队列），确保 block 正确释放。
4. **验证**：修复后重新进行压力测试，观察 IO 波动是否消除，确认无 `set bit was wrong` 或 cache id 冲突打印。

