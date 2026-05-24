# q
C++ 中 `unordered_map` 是什么？如何声明和初始化一个 `unordered_map<int, std::string>` 对象？
# a
`unordered_map` 是 C++ 标准库中的散列表（哈希表），用于存储键值对，提供快速的插入、查找和删除操作。可通过以下方式声明和初始化：
```cpp
std::unordered_map<int, std::string> myMap;                     // 默认构造
std::unordered_map<int, std::string> mp = { {1,"one"},{2,"two"}}; // 初始化列表
std::unordered_map<int, std::string> mp(10);                    // 预分配10个桶
std::unordered_map<int, std::string> mp2 = mp;                  // 拷贝构造
```

# q
C++ 的 `unordered_map` 如何插入或更新元素？`insert` 和 `emplace` 有什么区别？
# a
- 使用 `[]` 运算符：`mp["qx"] = 1;` —— 键存在时更新值，不存在时插入新键值对。
- `insert` 方法：`mp.insert({"1", 1});` —— 键存在时不覆盖，插入失败。
- `emplace` 方法：`myMap.emplace("Charlie", 95);` —— 直接在容器内存中构造元素，避免临时对象拷贝。

# q
C++ `unordered_map` 中如何查找元素？`find`、`count` 和 `[]` 访问有何区别？
# a
- `auto it = myMap.find("Charlie");` 返回迭代器，未找到则返回 `end()`。
- `if (myMap.count("David"))` 返回 0 或 1，适合仅判断键是否存在。
- `myMap["Alice"]` 直接访问值，键不存在时会插入一个默认构造的值；`myMap.at("Bob")` 在键不存在时抛出异常。

# q
在 Python 解决字母异位词分组问题中，`collections.defaultdict(list)` 如何工作？为什么要用它而不是普通字典？
# a
`defaultdict(list)` 会为不存在的键自动创建一个空列表，避免手动检查键是否存在。例如：`mp["aet"].append("eat")` 可以在键不存在时直接添加，无需 `if key not in mp: mp[key] = []` 的判断，使代码更简洁。

# q
如何对字符串排序来作为字母异位词的键？请写出 Python 的典型代码片段。
# a
```python
key = "".join(sorted(st))
```
- `sorted(st)` 将字符串 `st` 转化为排序后的字符列表，如 `st = "eat"` → `['a', 'e', 't']`。
- `"".join(...)` 将字符列表连接成字符串 `"aet"`，作为分组的键。

