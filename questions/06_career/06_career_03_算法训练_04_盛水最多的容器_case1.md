# q
该算法在处理空输入时存在什么典型的潜在缺陷？
# a
当输入 `height` 为空时，`size = height.size() - 1` 会导致无符号整数下溢（`size_t` 类型的 0 减 1 会变成极大值），导致后续 `r = size` 使用错误索引，访问越界。
核心代码片段：
```cpp
int size = height.size() - 1; 
r = size;                    // 若 height 为空，size 溢出为 SIZE_MAX
```

