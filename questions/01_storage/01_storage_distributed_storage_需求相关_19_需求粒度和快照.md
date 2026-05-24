# q
在快照创建流程中，Snap 如何与 DS（数据存储）进行联立交互？
# a
Snap 通过调用 `Snap.request_ds_create` 来请求 DS 创建，并发送 `ds_create_req` 消息（在 C 和 Python 中分别对应类 `DsCreateReq` 和操作类型 `DSMT_CREATE`）。DS 端通过 REST API `/ds/<string:vid>` 接收 post 请求，入口函数为 `send_ds_create_messages`，Python 类 `Ds` 注册在 API 资源中。后续 C 层调用 `ds_process_create` → `ds_load_segments`，通过消息 `SEGMGR_MSG_QUERY_DS_SEGMENT` 查询 segment 分布信息，在收到 `ds_process_query_ds_segmgr_rsp` 后调用 `ds_fetch_info_from_msg` 提取信息，最终执行 `ds_create` 创建 `ds_t` 结构并更新 segment 分布（`ds_update_segment_distr`）。

# q
DS 创建请求的消息类型标识符是什么？
# a
在 C 和 Python 侧分别使用常量 `DSMT_CREATE`（Python 和 C 的操作类型）来表示 DS 创建消息。

# q
DS 创建时如何获取 segment 分布信息？
# a
DS 创建过程通过 `ds_load_segments` 发起查询，发送 `SEGMGR_MSG_QUERY_DS_SEGMENT` 消息，然后在响应处理函数 `ds_process_query_ds_segmgr_rsp` 中调用 `ds_fetch_info_from_msg` 从响应消息中提取 segment 分布信息。

# q
DS 创建的 Python REST API 路由是如何定义的？
# a
Python 中通过 Flask-RESTful 注册资源：
```python
api.add_resource(Ds, "/ds/<string:vid>/<string:argument>", "/ds/<string:vid>")
```
`Ds` 类处理 post 请求，对应函数为 `send_ds_create_messages`。

