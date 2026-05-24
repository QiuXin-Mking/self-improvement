# q
在实现最长连续序列算法时，常见的拼写错误和API误用有哪些？
# a
典型错误包括：
1. 将 `unordered_set` 拼写为 `unodered_set`，导致编译错误；
2. 误用 vector 的 `append` 方法向集合插入元素，正确操作应为 `insert`；
3. 在正向遍历查找连续序列时，在不该使用 `else { continue; }` 的地方插入了该语句，导致循环被意外中断，造成漏判。

# q
如何从代码片段中定位“向 unordered_set 插入元素失败”的问题？
# a
检查插入操作的调用方式：
- 如果代码中出现了 `set_a.append(num)`，说明误用了 `vector::append`，unordered_set 没有 `append` 方法，应改为 `set_a.insert(num)`。
- 确认容器是否正确声明为 `unordered_set`，注意拼写不能是 `unodered_set`。

# q
解决“最长连续序列算法结果错误”的标准排查流程是什么？
# a
1. 检查集合构建部分：
   - 容器声明拼写是否正确（`unordered_set` 而非 `unodered_set`）；
   - 插入元素是否使用 `insert()` 而非其他不存在的成员函数；
2. 检查核心循环逻辑：
   - 确认只对“没有前驱的元素”启动内层 while 统计，避免多余 `else { continue; }` 打断循环；
   - 确认内层 while 正确递减并累计长度，最后用 `max` 更新全局答案；
3. 编译运行，用边界用例（空数组、单元素、无连续、全部连续）验证结果。

