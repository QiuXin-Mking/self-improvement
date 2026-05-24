# q
如何在 Python 中获取当前时间的 Unix 时间戳？
# a
使用 `time.time()` 函数，返回自 Epoch（1970年1月1日 UTC）以来的秒数（浮点数）。
```python
import time
current_timestamp = time.time()
```

# q
如何将 Unix 时间戳转换为本地可读的日期时间字符串？
# a
先用 `time.localtime(timestamp)` 将时间戳转换为本地时间的 `struct_time` 对象，再用 `time.strftime(format, struct_time)` 格式化为字符串。
```python
formatted_time = time.strftime("%Y-%m-%d %H:%M:%S", time.localtime(current_timestamp))
```

# q
两个代码片段在获取时间戳的方式上有何区别？哪种更简洁？
# a
- 第一个片段先调用 `time.time()` 将结果存入变量，再传给 `time.localtime`。
- 第二个片段直接在 `time.localtime` 的参数中调用 `time.time()`，省略中间变量，代码更简洁。  
两个片段在功能上完全等效。

# q
`time.strftime` 中 `"%Y-%m-%d %H:%M:%S"` 各部分代表什么？
# a
- `%Y`：四位数的年份（如 2023）
- `%m`：两位数的月份（01-12）
- `%d`：两位数的日期（01-31）
- `%H`：24 小时制的小时（00-23）
- `%M`：分钟（00-59）
- `%S`：秒（00-59）  
生成的格式示例：`2023-10-05 14:30:25`

