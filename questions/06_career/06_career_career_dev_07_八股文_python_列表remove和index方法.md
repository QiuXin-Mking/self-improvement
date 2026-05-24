# q
以下代码的输出结果是什么？
```python
lists = [1, 1, 1, 3, 5]
lists.remove(1)
res = lists.index(1)
print(res)
```
# a
输出结果为 `0`。

原因：
- `lists.remove(1)` 移除列表中**第一个**值为 `1` 的元素，列表变为 `[1, 1, 3, 5]`
- `lists.index(1)` 返回第一个匹配项 `1` 的索引，在 `[1, 1, 3, 5]` 中索引为 `0`
- `print(res)` 输出 `0`

# q
Python 中列表的 `remove()` 方法有哪些特点？
# a
`remove()` 方法的特点：
- **功能**：移除列表中**第一个匹配**指定值的元素，不删除所有匹配项
- **原地修改**：直接修改原列表，返回值是 `None`
- **异常**：若值不存在，抛出 `ValueError`
- **示例**：
  ```python
  lst = [1, 2, 3, 2]
  lst.remove(2)   # lst 变为 [1, 3, 2]
  lst.remove(5)   # ValueError
  ```
- **移除所有匹配项需用循环或列表推导式**（如 `[x for x in lst if x != value]`）

# q
Python 中列表的 `index()` 方法有哪些特点？如何查找所有匹配项的索引？
# a
`index()` 方法的特点：
- **功能**：返回第一个匹配项的索引（从 0 开始），不会返回所有匹配项
- **语法**：`list.index(value[, start[, end]])`，可选参数限定搜索范围
- **异常**：值不存在时抛出 `ValueError`
- **不修改原列表**

查找所有匹配项的索引示例：
```python
lst = [1, 2, 3, 2, 4, 2]
indices = [i for i, x in enumerate(lst) if x == 2]  # [1, 3, 5]
```

# q
如何高效移除列表中所有等于某个值的元素？列举至少两种方法。
# a
推荐使用**列表推导式**（最高效）：
```python
lst = [1, 1, 1, 3, 5]
lst = [x for x in lst if x != 1]   # 结果 [3, 5]
```

其他方法：
- **`filter()` 函数**：`list(filter(lambda x: x != 1, lst))`
- **`while` 循环**：`while 1 in lst: lst.remove(1)`（但时间复杂度 O(n²)）
- **倒序 `del`**：`for i in range(len(lst)-1, -1, -1): if lst[i] == 1: del lst[i]`

其中列表推导式简洁且性能最优。

