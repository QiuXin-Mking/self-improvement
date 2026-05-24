# q
快照删除的起点函数是什么？
# a
`snap_ds_delete`

# q
在快照删除流程中，什么情况下快照不能被删除？
# a
被依赖的快照不能删除

# q
ds 层处理消息的回调函数是哪个？
# a
`ds_process_msg_fn`

# q
sio 使用什么数据结构组织 segment？
# a
跳表（skip list）

# q
ocache_unmap_wcmd 在哪个流程中被调用？
# a
unmap 流程

