# q
在 C++ 的 `std::vector` 中，添加元素的正确方法是什么？
# a
应使用 `push_back()` 或 `emplace_back()`，而不是 Python 风格的 `append()`。代码中 `res.append(...)` 和 `ans.append(res)` 均为错误，需改为 `res.push_back(...)` 和 `ans.push_back(res)`。

# q
三数之和排序后，对第一个数 `fir` 去重时，这段代码的 `while` 循环存在什么问题？
# a
```cpp
while (fir > 0 && nums[fir] == nums[fir-1]) { fir++; }
```
此写法会让 `fir` 在外层 `for` 的自增之外继续增长，可能导致 `fir` 越过数组边界（例如当大量重复集中在末尾时），并且跳过一些本应处理的 `fir` 值。正确的去重应该用 `if` 判断并 `continue`：
```cpp
if (fir > 0 && nums[fir] == nums[fir-1]) continue;
```

# q
在内层双指针遍历 (`sec` 和 `thr`) 时，代码中 `while (nums[sec] + nums[thr] > target)` 可能导致什么问题？
# a
`thr` 会一直递减，但没有检查 `thr` 是否小于 `sec`，可能在内层 `for` 还未结束时 `thr` 已经小于 `sec`，导致后续访问越界或逻辑错误。应该在递减前或每次使用前检查 `sec < thr`，并在循环中添加该条件。

