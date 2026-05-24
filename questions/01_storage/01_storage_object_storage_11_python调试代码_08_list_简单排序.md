# q
在Python中，如何对一个列表同时进行去重和排序？
# a
可以使用 `set` 去重后再用 `sorted()` 排序。典型代码：
```python
numbers = [5, 2, 9, 1, 5, 6, 1, 2]
unique_numbers = list(set(numbers))      # 去重，顺序随机
sorted_unique_numbers = sorted(unique_numbers)  # 升序排序
# 结果：[1, 2, 5, 6, 9]
```

# q
为什么使用 `set` 去重后需要再调用 `sorted()`，而不能直接依赖 `set` 的顺序？
# a
`set` 是无序集合，不保留元素的插入顺序，且其内部顺序不可预测。因此去重后需要显式调用 `sorted()` 或 `list.sort()` 进行排序，才能得到有序结果。如果希望保留原始顺序并去重，需使用 `dict.fromkeys()` 等方法。

# q
`sorted()` 函数默认按什么规则排序？如何实现降序排序？
# a
`sorted()` 默认按升序排序（从小到大）。通过设置 `reverse=True` 参数可以实现降序排序。例如：
```python
sorted_unique_numbers_desc = sorted(unique_numbers, reverse=True)
# 结果：[9, 6, 5, 2, 1]
```

