# q
扩容后集群出现大量slow ops，典型根因是什么？
# a
典型根因是新增OSD后PG发生重新映射（remapped），但旧OSD上的PG数量已达到`osd_max_pg_per_osd`上限（例如阈值750），导致新PG被阻止创建，PG卡在`active+recovery_wait+undersized+degraded+remapped`状态，IO操作长时间阻塞，引发slow ops。日志关键字：`withhold creation of pg ... >= 750`。

# q
如何从原生日志定位扩容引起的slow ops问题？
# a
可按以下步骤定位：
1. 查看OSD日志中的slow ops报告：`tail -f ceph-osd.12.log` 观察到 `reporting 52 slow ops`。
2. 使用 `ceph daemon osd.X dump_ops_in_flight` 查看在途操作，检查高`duration`的op及其描述，例如 `"duration": 1580.815...` 和 `"flag_point": "delayed"`。
3. 在OSD日志中搜索 `withhold creation of pg`，如：
   ```
   2025-01-13T21:47:38.341+0800 7fdab44f9700  1 osd.0 3424 withhold creation of pg 2.61f: 752 >= 750
   ```
   这直接表明PG数量超过限制（752 >= 750）。
4. 执行 `ceph pg dump | grep undersized` 确认存在大量`undersized+degraded+remapped`的PG，且`acting`集合中只有一个OSD，说明副本不足。

# q
解决因`osd_max_pg_per_osd`限制导致扩容后slow ops的标准流程是什么？
# a
1. 临时提高所有OSD的PG上限，允许PG创建：
   ```bash
   ceph tell osd.* injectargs '--osd_max_pg_per_osd 800'
   ```
2. 监控PG状态，等待`undersized` PG完成恢复，直至变为`active+clean`：
   ```bash
   ceph pg dump | grep undersized
   ```
3. 数据恢复完成后，可逐步将参数调回原值以避免未来新OSD再次触发限制；如需长期解决方案，应规划合理的PG数量或调整相关阈值。

