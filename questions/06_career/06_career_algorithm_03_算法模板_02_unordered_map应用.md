# q
`std::unordered_map` 的 `count` 函数的作用是什么？如何用它判断键是否存在？
# a
`count` 返回特定键在 `unordered_map` 中出现的次数（`0` 或 `1`，因为键唯一）。可以用 `if (myMap.count(key))` 判断键是否存在。

# q
`std::unordered_map` 的 `empty` 函数返回什么？什么时候返回 `true`？
# a
`empty()` 检查容器是否不含任何元素，若为空返回 `true`，否则返回 `false`。

# q
如何用列表初始化的方式创建包含初始键值对的 `std::unordered_map`？
# a
使用花括号列表初始化，例如：
```cpp
unordered_map<char, char> pairs = {
    {')', '('},
    {']', '['},
    {'}', '{'}
};
```

