# q
mdlog_replay_log 在处理日志时，解析和加载的单位分别是多少？
# a
以 4K 大小解析日志，以 1M 大小加载响应。

# q
在 mdlog 分析中，如何判断日志是否连续？
# a
检查前面的 osd id 和后面的 osd id 是否连续，前面的 segment id 和后面的 segment id 是否一致，并且整个 block_num 长度中的 block idx 中没有 INVALID_BLK_ID。

# q
blist_to_iovec 的作用是什么？
# a
将 blist 转换为 iovec 结构，所有数据都是从 blist 返回的，新起一个 iovec。

# q
osd_mdlog_buf_add_update_log 函数的关键实体和参数有哪些？
# a
入参是 oio，出参是 mdlog_buf（即 mdlog 实体），关键参数包括 mdlog_buf_len，oio->mdlog_buf。

# q
在 mdlog 回放时，如何判断日志是否到了尾部以及如何处理不完整的头部？
# a
通过判断本次是否为最后一次数据，或本次与下一次不连续来确认是否到达尾部。若头部不完整，会将数据放入临时 item 区，并且 finish_len 表示本次取到的不定长日志中已完成 block 的长度。

