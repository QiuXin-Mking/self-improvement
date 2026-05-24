# q
conn_send 函数在处理消息时，如何根据连接类型分发消息？
# a
- **CONN_MULTICONN**：直接调用 `multiconn_send`，作为 engine 与 Go 的桥梁。
- **本地节点**（`conn_id.type == CONN_NODE && conn_id.id == g_cluster_md.nid` 或等于 `LOCAL_CONN_ID`）：设置 `conn_id` 为 `LOCAL_CONN_ID`；若为请求消息且存在回调或上下文，将消息插入 `msg_base_map`；最终调用 `engine_conn_send`。
- **远程节点**：遍历 `sock_type`（从 `SOCK_TYPE_MAX-1` 向下到 `SOCK_TYPE_TCP`），通过 `conn_get_socket` 获取可用 socket，若为空则返回 `ENOBUFS`，否则通过 `conn_ctx.conn_fn_map[sock_type].send_fn` 发送；发送后调用对应的 `put_fn` 释放 socket。
- **消息过滤**：请求消息先经过 service 后端过滤器，若匹配成功则直接释放消息并返回 `S_OK`，不继续发送。

# q
conn_send 在节点间通信时，选择 socket 的优先级顺序是怎样的？
# a
按 `sock_type` 从高到低遍历，起始于 `SOCK_TYPE_MAX-1` 直到 `SOCK_TYPE_TCP`。对每个类型调用 `conn_get_socket(conn_id, sock_type)`，返回第一个非 NULL 的 socket。若所有类型均无可用 socket，则返回错误码 `ENOBUFS`。

# q
CONN_MULTICONN 在 conn_send 中的作用是什么？
# a
当 `conn_id.type` 为 `CONN_MULTICONN` 时，`conn_send` 直接调用 `multiconn_send(conn_id, msg)` 发送消息，该通道作为 engine 与 Go 部分通信的桥梁。

# q
本地节点消息处理时，msg_base_map 何时被使用？
# a
当连接标识为本节点（本地 CONN_NODE 或 LOCAL_CONN_ID）且消息是请求类型，并且 `msg_context` 或 `msg_callback` 非空时，`conn_send` 会获取 `msg_base_map`，将该消息插入映射表（`msg_base_map_insert`），以支持后续的上下文或回调匹配。

