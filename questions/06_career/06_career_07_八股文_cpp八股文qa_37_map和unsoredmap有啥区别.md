# q
map 的底层实现通常是什么数据结构？
# a
红黑树（有序平衡二叉树）

# q
unordered_map 的底层实现通常是什么数据结构？
# a
哈希表

# q
map 和 unordered_map 的查找/插入/删除时间复杂度分别是多少？
# a
map：O(logN)；unordered_map：平均 O(1)，最坏 O(N)

# q
map 和 unordered_map 分别对 key 类型有什么要求？
# a
map 的 key 需要支持 `operator<` 比较；unordered_map 的 key 需要支持哈希函数和相等判断（通常通过自定义哈希器和 `==`）

# q
unordered_map 的迭代器在什么情况下可能全部失效？
# a
当进行可能触发 rehash 的删除操作时，所有迭代器都可能失效

