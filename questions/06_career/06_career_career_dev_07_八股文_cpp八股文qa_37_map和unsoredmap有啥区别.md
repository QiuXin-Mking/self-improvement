# q
C++ STL中map和unordered_map的底层实现结构分别是什么？
# a
- **map**：底层实现是**红黑树**（有序平衡二叉树）
- **unordered_map**：底层实现是**哈希表**

# q
map和unordered_map在查找、插入、删除操作上的平均时间复杂度有何差异？
# a
- **map**：查找/插入/删除时间复杂度为 **O(logN)**
- **unordered_map**：查找/插入/删除平均时间复杂度为 **O(1)**，最坏情况 O(N)

# q
对容器的key类型，map和unordered_map分别有什么要求？
# a
- **map**：key类型需要支持 `<` 比较（`operator<`）
- **unordered_map**：key类型需要支持哈希函数和相等判断（通常可自定义哈希器）

# q
map和unordered_map在元素顺序和遍历稳定性上有什么不同？
# a
- **map**：键自动有序，支持正序/逆序遍历，遍历顺序稳定
- **unordered_map**：元素无序，每次遍历的顺序可能不同

# q
map和unordered_map的内存占用和迭代器失效特点是什么？
# a
- **unordered_map** 通常比 map 占用更多内存（因为哈希结构和可能的扩容）
- 两者迭代器在删除当前元素外的操作一般不会失效，但 **unordered_map 删除时可能触发 rehash，导致所有迭代器失效**

