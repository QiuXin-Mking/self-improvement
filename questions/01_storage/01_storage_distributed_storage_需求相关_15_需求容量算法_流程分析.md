# q
DS容量查询流程的起点是什么？涉及哪些关键组件？
# a
流程起点是由proxy发起的 `ds_query_volume_used_capacity` 请求，用于扫描并查询DS（数据存储）的使用容量。涉及的关键标识包括 `SMT_SEGMENT_QUERY_USED_CAP`、`OMT_QUERY_USED_CAP`，以及后续在OSD侧的 `osd_process_query_used_capacity`。

# q
当并发查询容量过程中任意一条子请求失败时，流程如何处理？
# a
任何一条并发查询失败时，必须立即中止整个并发流程，并向上层返回错误。同时，系统可能需要减小最大并发数以避免问题。获得错误结果后，会通过 **adapter服务** 向 **go 服务**查询 etcd 中记录的历史容量值作为回退数据。

# q
adapter服务在容量查询流程中扮演什么角色？
# a
adapter服务作为回退机制，在查询失败或返回错误结果时，负责从 etcd 中读取之前存储的容量历史值，保证上层仍能获得可用的容量信息。

# q
OSD端处理segment容量查询的核心函数有哪些？
# a
OSD 端的主要函数包括：
- `osd_make_query_request`（构造查询请求）
- `osd_process_query_used_capacity`（处理容量查询）
- `osd_segment_get_block_num`（获取segment的block数量，用于容量计算）
读取这些数据可能耗时较长。

