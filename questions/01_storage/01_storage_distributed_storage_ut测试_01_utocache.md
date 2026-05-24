# q
ocache_alloc_io 函数的主要作用是什么？
# a
ocache_alloc_io 为接收到的 msg 分配并初始化一个 ocache_io_t 结构：增加 msg 的引用计数，设置 pri_msg 回调指针，初始化链表（如 ocache_cmd_list）和队列，在 g_ocache_ctrl 中维护 oio 变量，并对整个结构体进行 memset 初始化和指针、队列的初始化。

# q
ocache_split_cmd 完成了哪些核心操作？
# a
ocache_split_cmd 解析 ocache_io 中 msg->data 携带的 req_msg 命令：计算命令数量 cmd_num，分配对应的 ocache_cmd_t 结构，计算每条命令的起始 LBA 位置和长度，根据 req_msg 中 io_hdr 的 type 映射为 ocache cmd 类型（如 OCACHE_READ），并将这些 cmd 插入到 ocache_io->ocache_cmd_list 链表中。

# q
ocache 中的 cmd 链表如何完成入队和出队操作？
# a
使用内核链表宏对称操作：
- 出队：`ocache_cmd = list_first_entry(&ocache_io->ocache_cmd_list, ocache_cmd_t, cmd_ionode);`
- 入队：`list_add_tail(&ocache_cmd->cmd_ionode, &ocache_io->ocache_cmd_list);`
通过 cmd_ionode 成员将 ocache_cmd_t 串入双链表。

# q
ocache 支持的 I/O 类型有哪些？在代码中如何体现？
# a
支持读、写、精简读、精简写、umap 命令。在 msg 结构中通过 type 字段（如 OCACHE_READ_MSG）和 io_hdr.io_type（如 IO_READ）标识具体 I/O 类型，解析后对应的 ocache_cmd->cmd_type 被设置为 OCACHE_READ 等值。

