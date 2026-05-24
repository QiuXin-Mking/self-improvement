# q
什么是Buffer pool size？它表示什么？
# a
Buffer pool size 是 InnoDB 缓冲池的总容量，以页为单位（默认每页 16KB）。它决定了缓冲池中能存放的最大页面数量，直接影响数据库性能。可通过 `innodb_buffer_pool_size` 配置，建议设置为系统内存的 50%-70%。查看命令：`SHOW STATUS LIKE 'Innodb_buffer_pool_pages_total';`

# q
什么是Modified db pages？它对性能有何影响？
# a
Modified db pages 是缓冲池中的脏页数，即已被修改但尚未刷写到磁盘的页面。脏页由 UPDATE/INSERT/DELETE 等写操作产生。脏页过多会增加崩溃时数据丢失的风险，并可能在需要空间时引发强制刷写，导致 I/O 峰值。监控指标：`SHOW STATUS LIKE 'Innodb_buffer_pool_pages_dirty';`。优化可通过调整 `innodb_max_dirty_pages_pct` 和 `innodb_io_capacity` 参数。

# q
什么是Pages read？如何利用它计算缓存命中率？
# a
Pages read 是自 MySQL 启动以来从磁盘读取到缓冲池的页的累计数。它代表未能命中缓存的物理读次数，值越小说明缓存越有效。缓存命中率计算公式：`(Innodb_buffer_pool_read_requests - Innodb_buffer_pool_reads) / Innodb_buffer_pool_read_requests * 100%`。通过 `SHOW STATUS LIKE 'Innodb_buffer_pool_reads';` 查看。

# q
InnoDB Buffer Pool 的各项指标之间有何基本关系？
# a
基本容量关系：Buffer pool size ≈ Free buffers + Database pages + Modified db pages + 其他系统页。其中 Free buffers 为空闲页，Database pages 为干净数据页，Modified db pages 为脏页。已使用页数 = Database pages + Modified db pages，利用率 = 已使用页数 / Buffer pool size。脏页比例 = Modified db pages / 已使用页数。这些指标共同反映内存利用情况、脏页负担和缓存健康度。

# q
如何查看 InnoDB Buffer Pool 的实时状态？
# a
主要方法包括：
- `SHOW STATUS LIKE 'Innodb_buffer_pool%';` 查看各具体指标。
- `SHOW ENGINE INNODB STATUS\G` 查看 BUFFER POOL AND MEMORY 段落。
- 查询 `INFORMATION_SCHEMA.INNODB_BUFFER_POOL_STATS` 表获得更结构化的数据。
- 使用 Performance Schema（MySQL 5.7+）监控相关事件。
示例计算使用率：`SELECT ROUND((VARIABLE_VALUE)/(SELECT VARIABLE_VALUE FROM INFORMATION_SCHEMA.GLOBAL_STATUS WHERE VARIABLE_NAME='Innodb_buffer_pool_pages_total')*100,2) AS usage_pct FROM INFORMATION_SCHEMA.GLOBAL_STATUS WHERE VARIABLE_NAME='Innodb_buffer_pool_pages_data';`

