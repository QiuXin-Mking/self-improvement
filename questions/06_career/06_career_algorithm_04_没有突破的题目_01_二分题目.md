# q
在二分查找中，使用 `if (matrix[mid][0] > target) { right = mid - 1; } else { left = mid + 1; }`，`while (left <= right)`，为何 `right` 可能变为 -1 并导致报错？
# a
当 target 小于所有元素时，最后一次迭代 `left == right`（假设初始 `left=0`），mid 等于 left，条件 `matrix[mid][0] > target` 成立，执行 `right = mid - 1`，导致 `right` 变为 -1。循环结束后若以 `right` 作为索引访问数组，就会越界报错。

# q
在二分查找中，使用 `if (matrix[col][mid] < target) { left = mid + 1; } else { right = mid - 1; }`，`while (left <= right)`，为何 `left` 可能超出数组最大下标？
# a
当 target 大于所有元素时，最后一次迭代 `left == right`（假设 `right` 为数组最后一个索引），mid 等于 left，条件 `matrix[col][mid] < target` 成立，执行 `left = mid + 1`，导致 `left` 变为 `right+1`，即数组长度。循环结束后若以 `left` 作为索引访问数组，将导致越界。

