# q
如何用 Python 正则表达式提取多行文本每行的第一个单词？
# a
使用模式 `r'^\s*([\w]+)\s+'` 并结合 `re.findall` 及 `re.MULTILINE` 标志。示例：
```python
import re
data = """sdk   8:160  0  10G  0 disk
sdl   8:176  0  10G  0 disk"""
pattern = r'^\s*([\w]+)\s+'
matches = re.findall(pattern, data, re.MULTILINE)
```
`re.MULTILINE` 使 `^` 匹配每行开头，模式会捕获第一个单词（连续的单词字符）并忽略前导空格。

# q
正则表达式 `^\s*([\w]+)\s+` 中各个部分的含义是什么？
# a
- `^`：锚点，匹配行首（配合 `re.MULTILINE` 时为每行开头）。
- `\s*`：匹配零个或多个空白字符（处理前导空格）。
- `([\w]+)`：捕获组，匹配一个或多个单词字符（`\w` 表示字母、数字、下划线）。
- `\s+`：匹配一个或多个空白字符，分隔第一个单词与后续内容。

# q
`re.findall(pattern, data, re.MULTILINE)` 中 `re.MULTILINE` 的作用是什么？
# a
`re.MULTILINE`（或 `re.M`）让正则中的 `^` 和 `$` 分别匹配每一行的开始和结束，而不是整个字符串的开始和结束。这样 `^\s*([\w]+)\s+` 就能逐行匹配并提取该行的第一个单词。

# q
正则表达式中的 `[\w]` 匹配哪些字符？
# a
`[\w]` 匹配单词字符（word character），包括：
- 字母（大小写）
- 数字（0‑9）
- 下划线 `_`
在 Python 中默认等同于 `[a-zA-Z0-9_]`。

