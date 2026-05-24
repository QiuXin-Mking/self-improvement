# q
clog压缩的整体流程分为哪几个主要阶段？
# a
分为三个阶段：压缩触发、压缩处理、压缩完成判断与通知。具体流程：压缩触发后，压缩请求发送给 ocache，由 `ocache_io_process_req` 接收，调用 `ocache_compact_clog` 执行压缩；然后判断压缩是否完成，未完成则继续压缩下一个 segment，完成后通知 clog 压缩完成。

# q
ocache 接收 clog 压缩请求并处理的入口函数是什么？
# a
`ocache_io_process_req`

# q
clog 侧处理压缩相关的核心函数有哪些？
# a
- `clog_process_compact`
- `clog_compact_log_write_req`
- `clog_update_header_req`
- `ocache_append_clog`

# q
压缩完成后如何通知 clog？
# a
通过 `clog_process_compact` 等函数完成压缩结果的写入和 header 更新（`clog_update_header_req`），并最终通过 `ocache_append_clog` 附加日志，完成压缩完成通知。

