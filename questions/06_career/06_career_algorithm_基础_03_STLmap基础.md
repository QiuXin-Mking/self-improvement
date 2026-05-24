# q
STL map 的嵌套定义时需要注意什么语法细节？
# a
嵌套定义时最后两个 `>` 之间必须留有空格，例如 `map<string, map<string, long> >`，否则编译器可能将其解析为右移运算符 `>>`。

# q
map 如何使用下标运算符访问或插入元素？
# a
map 支持 `operator[]`，形式为 `map[key]`。如果 key 存在则返回对应值的引用；如果 key 不存在，则会插入一个使用默认值构造的元素并返回其引用。

# q
如何用迭代器遍历 map 中的所有键值对？
# a
```cpp
map<int, string>::iterator iter;
for (iter = mapStudent.begin(); iter != mapStudent.end(); ++iter) {
    cout << iter->first << " " << iter->second << endl;
}
```
其中 `iter->first` 是键，`iter->second` 是值。

