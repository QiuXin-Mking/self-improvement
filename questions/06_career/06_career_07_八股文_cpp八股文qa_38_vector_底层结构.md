# q
std::vector 的扩容因子通常是多少？
# a
扩容时新容量通常是当前容量的 2 倍（GCC libstdc++）或 1.5 倍（MSVC 等），具体由实现决定。示例：
```cpp
size_t new_capacity = old_size == 0 ? 1 : old_size * 2;
```

# q
为什么 std::vector 在增长元素时可能导致迭代器失效？
# a
因为扩容会重新分配一块更大的连续内存，把原有元素移动过去，然后释放原内存。原迭代器指向的地址空间被销毁，继续使用会变成悬空指针。

# q
std::vector 的随机访问和中间插入/删除的时间复杂度分别是多少？
# a
- 随机访问（operator[]、at()）：O(1)
- 中间插入/删除：O(n)（需要移动后续元素）

# q
在 GCC libstdc++ 中，std::vector 的核心私有成员变量有哪些？
# a
三个指针：
- `_start`：指向已用空间的起始位置
- `_finish`：指向已用空间的尾后位置（`_start + size()`）
- `_end_of_storage`：指向存储容量的尾后位置（`_start + capacity()`）

