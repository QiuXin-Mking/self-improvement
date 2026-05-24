# q
vector push_back 什么时候会触发扩容？
# a
当 size() == capacity() 时，插入新元素会触发扩容。

# q
vector push_back 扩容的时间复杂度是多少？
# a
每次扩容需拷贝所有元素，时间复杂度为 O(n)。频繁扩容时，总时间复杂度可能退化为 O(n²)。

# q
vector push_back 扩容会造成什么内存问题？
# a
扩容后未使用的容量（如容量为 100，但仅使用 50）造成内存浪费。

