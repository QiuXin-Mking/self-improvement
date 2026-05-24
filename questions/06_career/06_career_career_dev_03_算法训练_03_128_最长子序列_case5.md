# q
这段最长连续序列代码中存在什么典型的变量使用错误？
# a
代码第3行 `set_a = set(num)` 中 `num` 应为 `nums`（函数参数名），正确写法是 `set_a = set(nums)`。另外第9行 `while (num+1) in set_a:` 中的 `num+1` 应为 `cur+1`，因为需要基于当前数 `cur` 来递增查找连续序列，而不是固定使用最初的 `num`。

# q
如何修正这个最长连续序列解法的关键 bug 使其正确运行？
# a
需修改两处：
1. 将 `set_a = set(num)` 改为 `set_a = set(nums)`。
2. 将 `while (num+1) in set_a:` 改为 `while (cur+1) in set_a:`，同时在循环内将 `cur = cur + 1` 放在 `while` 内部开头或之前。

修正后代码示例：
```python
class Solution:
    def longestConsecutive(self, nums: List[int]) -> int:
        set_a = set(nums)
        ans = 0
        for num in set_a:
            if (num - 1) not in set_a:
                cur = num
                lev = 1
                while (cur + 1) in set_a:
                    lev += 1
                    cur += 1
                ans = max(ans, lev)
        return ans
```

