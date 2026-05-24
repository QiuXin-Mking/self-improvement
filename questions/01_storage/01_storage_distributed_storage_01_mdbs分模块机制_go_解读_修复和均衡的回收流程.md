# q
修复或均衡的回收流程包含哪些关键步骤？
# a
流程从 `newPoolSyncInfo` 启动，分别进入修复协程 `repairingGoroutine` 或均衡协程 `balancingGoroutine`，然后执行 `segSyncData` 进行段数据同步，之后调用 `segRepairRecycle` 和 `oRecycle` 完成回收准备，最终通过 `osd_process_req` 或 `osd_process_system_req` 发出 `OMT_RECYCLE` 系统请求，由 `osd_process_recycle` 处理实际回收。

# q
`OMT_RECYCLE` 消息在回收流程中起什么作用？
# a
`OMT_RECYCLE` 是一条系统请求消息，在 `oRecycle` 阶段后经由 `osd_process_req` 或 `osd_process_system_req` 发出，用于驱动最终的实际回收操作，由 `osd_process_recycle` 函数处理。

# q
修复流程和均衡流程分别由哪个协程执行？
# a
修复流程由 `repairingGoroutine` 协程执行，均衡流程由 `balancingGoroutine` 协程执行。

# q
`osd_process_req` 和 `osd_process_system_req` 在回收流程中扮演什么角色？
# a
它们用于将 `oRecycle` 阶段生成的 `OMT_RECYCLE` 请求作为系统请求发送出去，从而触发 `osd_process_recycle` 的执行。

