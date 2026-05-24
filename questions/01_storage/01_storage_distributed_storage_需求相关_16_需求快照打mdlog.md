# q
block_map_segment_copy函数的作用是什么？
# a
在快照首次写入时，负责复制 segment map。

# q
快照打mdlog需求中涉及哪些与mdlog直接相关的处理函数？
# a
包括 `mdlog_process_append`（处理mdlog追加）和 `mdlog_make_request`（构造mdlog请求）等。

