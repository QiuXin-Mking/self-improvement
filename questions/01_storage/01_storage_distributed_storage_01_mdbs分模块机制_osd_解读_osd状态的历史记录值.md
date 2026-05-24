# q
查询etcd中OSD状态历史值的整体调用流程是怎样的？
# a
1. 由 `autoconn main` 发起  
2. 调用 `osdmgr_query_osdlist` 查询OSD列表  
3. 进入 `adapter_process_req` 处理请求，启动 `amtWatchOsds go` 和 `osdWatch go` 协程观察etcd变更  
4. 观察收到数据后，`adapter_process_rsp` 处理响应  
5. 最终 `osdmgr_process_get_osd_list_rsp` 处理 `AMT_GET_OSD_LIST` 响应消息

# q
`osdmgr_process_get_osd_list_rsp` 在OSD状态历史查询中的作用是什么？
# a
它负责处理 `AMT_GET_OSD_LIST` 消息的响应，即从etcd返回的OSD状态历史数据。

