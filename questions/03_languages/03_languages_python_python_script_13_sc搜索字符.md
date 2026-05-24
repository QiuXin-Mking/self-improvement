# q
如何使用 Python 的正则表达式模块 `re` 检测文本中是否包含某个单词？
# a
导入 `re` 模块，定义模式字符串（如 `r'inconsistency'`），然后调用 `re.search(pattern, text)`。若找到匹配则返回一个 `Match` 对象（判断为真），否则返回 `None`（判断为假）。示例代码：
```python
import re
pattern = r'inconsistency'
if re.search(pattern, text):
    # 存在匹配
```

# q
题目中的 `search_inconsistency` 函数实现了什么功能？
# a
接收一个字符串参数 `text`，对其进行正则搜索，判断是否存在子串 `inconsistency`。若存在则返回字符串 `"Found 'inconsistency' in the text."`，否则返回 `"Did not find 'inconsistency' in the text."`。

