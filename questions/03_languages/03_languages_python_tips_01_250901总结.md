# q
Python 文件头部声明 `# -*- coding: utf-8 -*-` 的作用是什么？
# a
指定源文件的字符编码为 UTF‑8，确保解释器能正确解析文件中的中文字符等非 ASCII 内容。

# q
如何使用 `time` 模块统计 Python 脚本的执行时间，并按秒和分钟格式化输出？
# a
在脚本开头导入 `time` 并记录起始时间，结束时计算耗时并格式化打印。示例代码：

```python
import time

start_time = time.time()           # 记录起始时间
# ... 待计时代码 ...
end_time = time.time()             # 记录结束时间
duration = end_time - start_time
print(f"总耗时：{duration:.2f} 秒，约 {duration/60:.2f} 分钟")
```

