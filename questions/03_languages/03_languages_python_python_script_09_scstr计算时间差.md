# q
如何使用 Python 的 datetime 模块计算两个日期时间字符串的时间差？
# a
使用 `datetime.strptime` 方法按照格式字符串解析两个日期字符串为 `datetime` 对象，然后直接相减，得到 `timedelta` 对象，即时间差。
```python
from datetime import datetime

datetime_obj1 = datetime.strptime("2023-09-06 11:12:31.763", "%Y-%m-%d %H:%M:%S.%f")
datetime_obj2 = datetime.strptime("2023-09-06 11:12:31.765", "%Y-%m-%d %H:%M:%S.%f")
time_difference = datetime_obj2 - datetime_obj1
```

# q
`strptime` 格式字符串 `"%Y-%m-%d %H:%M:%S.%f"` 中各占位符的含义是什么？
# a
- `%Y`：四位年份（如 2023）
- `%m`：两位月份（01-12）
- `%d`：两位日期（01-31）
- `%H`：24 小时制小时（00-23）
- `%M`：分钟（00-59）
- `%S`：秒（00-59）
- `%f`：微秒（000000-999999）

# q
`timedelta` 对象在默认打印时展示哪些信息？
# a
默认输出以 `days, seconds, microseconds` 格式显示，分别代表天数、秒数和微秒数。例如时间差可能显示为 `0:00:00.002000`（0 天 0 秒 2000 微秒）。

