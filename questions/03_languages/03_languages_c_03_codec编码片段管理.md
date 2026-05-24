# q
这段代码如何判断一个 block 是否有脏数据？
# a
它通过检查脏位图 `dirty_btmp` 的前两个 `uint64_t` 值是否同时为 0 来判断：如果都为 0，则 `pre_dirty_blk` 设为 `false`，否则设为 `true`。条件表达式为 `(0 == *((uint64_t *)blk_info->dirty_btmp)) && (0 == *((uint64_t *)blk_info->dirty_btmp + 1))`。

# q
为什么代码中要检查 `dirty_btmp` 的前两个 `uint64_t`？
# a
因为 `dirty_btmp` 被强制转换为 `uint64_t *` 指针，以 64 位为单位访问位图。检查前两个 64 位无符号整数意味着脏状态由至少 128 位的位图表示，可以覆盖更大范围的 block；只有这两个整数都为 0 时才认为整个位图没有脏数据。

