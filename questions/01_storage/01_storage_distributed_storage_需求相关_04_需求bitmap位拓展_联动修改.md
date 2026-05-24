# q
在分布式存储的IO路径中，bitmap的位数扩展是如何实现的？
# a
bitmap的位数可以直接扩展（"这个位数是可以直接拓展的"），无需大幅改造数据结构。

# q
bitmap在什么场景下才会被使用？
# a
仅在速读（speed read）操作时才会用到bitmap，用于记录和匹配块状态。

# q
`osd_reply->bitmap`的信息流如何传递到后续处理逻辑？
# a
信息流来自`osd_reply->bitmap`，通过条件`msg.data==osd_reply.bitmap`进行匹配和分发。

# q
bitmap位数扩展作为诱因，会联动修改哪些关键处理函数？
# a
涉及osd层和sio层的多个函数，包括`osd_process_io`、`osd_process_speed_write`、`sio_process_sync_write`、`osd_process_speed_read`、`osd_speed_read_io_split`等。

