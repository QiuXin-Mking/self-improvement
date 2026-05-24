# q
Python 解析 .conf 配置文件的核心流程是什么？
# a
逐行读取文件，忽略空行和 `#` 开头的注释行；遇到 `[SectionName]` 形式的行则创建一个新的配置段（嵌套字典）；其余行按 `key = value` 格式用 `split("=", 1)` 分割键值，去除首尾空格后存入当前段字典。最终返回一个嵌套字典结构，通过 `config[section][key]` 访问参数。

# q
代码中如何识别和提取配置段（Section）？
# a
检查行是否以 `[` 开头并以 `]` 结尾，若是则提取中间内容作为段名：
```python
if line.startswith("[") and line.endswith("]"):
    current_section = line[1:-1]
    config[current_section] = {}
```

# q
解析键值对时，如何处理包含等号的值？
# a
使用 `line.split("=", 1)` 限制分割次数为 1，防止值中的 `=` 被错误拆分：
```python
key, value = line.split("=", 1)
key = key.strip()
value = value.strip()
```

