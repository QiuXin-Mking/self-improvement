# q
time33 哈希函数通常也被称为什么？
# a
DJBX33A

# q
time33 哈希算法的核心计算步骤是什么？
# a
初始 hash 为 0，遍历字符串每个字符，执行 `hash = hash * 33 + (unsigned long) str[i]`，最后返回 hash 值（`uint32_t` 类型）

