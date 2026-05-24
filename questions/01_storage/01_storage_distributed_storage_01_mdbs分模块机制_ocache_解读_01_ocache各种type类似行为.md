# q
ocache_cmd_type 枚举中定义了哪些读写相关的命令类型？
# a
定义了以下命令类型：
- `OCACHE_READ`：普通读命令
- `OCACHE_READ_LIFT`：读提升命令
- `OCACHE_WRITE_BACK`：写回命令
- `OCACHE_WRITE_THROUGH`：透写命令
- `OCACHE_WRITE_AROUND`：写穿命令（ocache 进入只读状态）
- `OCACHE_FLUSH`：刷盘命令
- `OCACHE_SPEED_READ`：快速读命令
- `OCACHE_SPEED_WRITE`：快速写命令
- `OCACHE_UNMAP`：unmap 命令
- `OCACHE_READ_AHEAD`：预读命令

# q
OCACHE_WRITE_BACK、OCACHE_WRITE_THROUGH 与 OCACHE_WRITE_AROUND 三者的核心区别是什么？
# a
三者都是写命令，区别在于：
- `OCACHE_WRITE_BACK`（写回）：先将数据写入缓存，再异步写回后端存储，写操作延迟低但存在数据丢失风险。
- `OCACHE_WRITE_THROUGH`（透写）：同步将数据写入缓存和后端存储，保证一致性但写延迟较高。
- `OCACHE_WRITE_AROUND`（写穿）：数据直接写入后端存储，绕过缓存，并会使 ocache 进入只读状态。

# q
OCACHE_READ_LIFT 是什么，它可能用于什么场景？
# a
`OCACHE_READ_LIFT` 是读提升命令，可能用于将缓存中较低层级（如机械盘）的数据主动提升到更高性能的缓存层（如 SSD 或内存），以加速后续访问。具体行为可参考 [[【ocache】读提升机制]]。

# q
OCACHE_FLUSH 的作用是什么？
# a
`OCACHE_FLUSH` 是刷盘命令，用于将缓存中的脏数据强制写回后端存储，以保证数据持久化。

