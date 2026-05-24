# q
闭区间二分查找的 Python 标准模板是什么？
# a
```python
def binary_search_closed_interval(arr, target):
    left = 0
    right = len(arr) - 1          # 闭区间 [left, right]
    while left <= right:          # 注意条件是 <=
        mid = left + (right - left) // 2
        if arr[mid] == target:
            return mid
        elif arr[mid] < target:
            left = mid + 1
        else:
            right = mid - 1
    return -1
```

# q
`std::lower_bound` 的功能是什么？
# a
在已排序的范围内查找第一个**不小于**（即大于或等于）目标值的元素的位置。

# q
如何使用闭区间写法实现 `lower_bound`（返回第一个 >= target 的位置，即使不存在也返回插入点）？
# a
```python
def lower_bound(nums: List[int], target: int) -> int:
    left = 0
    right = len(nums) - 1          # 闭区间 [0, n-1]
    while left <= right:
        mid = left + (right - left) // 2
        if nums[mid] < target:     # 小于才右移，保证 left 最终停在 >= target 的首位
            left = mid + 1
        else:
            right = mid - 1
    return left                    # left 维持第一个 >= target 的特性
```

# q
如何使用左闭右开区间 `[0, n)` 实现 `lower_bound`？
# a
```python
def lower_bound(nums: List[int], target: int) -> int:
    left = 0
    right = len(nums)              # 右开区间 [0, n)
    while left < right:            # 条件为 left < right
        mid = left + (right - left) // 2
        if nums[mid] < target:
            left = mid + 1
        else:
            right = mid            # 保持右开，mid 已满足条件但区间不含 mid
    return left                    # left/right 重合，指向第一个 >= target 的位置
```

# q
C++ 中如何使用模板实现闭区间版本的 `lower_bound`？
# a
```cpp
template<typename T>
int lower_bound(const std::vector<T>& nums, T target) {
    int left = 0;
    int right = nums.size() - 1;           // 闭区间 [left, right]
    while (left <= right) {
        int mid = left + (right - left) / 2;
        if (nums[mid] < target) {
            left = mid + 1;
        } else {
            right = mid - 1;
        }
    }
    return left;                           // 第一个 >= target 的位置
}
```

