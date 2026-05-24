# q
mdlog回放流程的整体步骤是什么？
# a
mdlog回放始于osdmgt_scan_osds，依次经历：
```
osdmgt_scan_osds_load → MDLOG_OSD_SCAN → mdlog_process_osd_scan → mdlog_head_info_read → LOG_HEADER_LOAD → mdlog_load_header_rsp → mdlog_send_osdscan_to_ocache → OCACHE_OSD_SCAN → ocachemgt_process_osd_scan → ocachemgt_osd_scan_data → MDLOG_CACHE_SCAN → mdlog_process_cache_scan → mdlog_load_space_mgt_make_req → mdlog_space_persist_read → LOG_SPACE_MGT_LOAD → mdlog_load_space_mgt_rsp → mdlog_scan → LOG_SPACE_REPLAY_LOAD → mdlog_load_active_space_log_rsp → mdlog scan successful → mdlog_scan_end
```
扫描结束后通过`mdlog_send_osdmgt_rsp`将结果返回给`osdmgt_mdlog_rsp`，最终OSD上线。

# q
mdlog回放中`ocachemgt_osd_scan_data`步骤的作用是什么？
# a
`ocachemgt_osd_scan_data`负责将扫描请求转发给mdlog处理，触发`MDLOG_CACHE_SCAN`，进而由`mdlog_process_cache_scan`加载空间管理元数据并执行实际日志回放（`mdlog_scan`→`LOG_SPACE_REPLAY_LOAD`）。

# q
OSD上线过程中，mdlog扫描完成后如何通知OSD上线？
# a
mdlog扫描成功后，通过`osdmgt_mdlog_rsp`将扫描结果返回给`osdmgt_send_osd_scan_ready_req`，再经由`osd_make_mgt_request`发送`OMT_OSD_SCAN_READY`消息，最终由`osd_process_osd_scan_ready`设置OSD为online并激活存储引擎。

