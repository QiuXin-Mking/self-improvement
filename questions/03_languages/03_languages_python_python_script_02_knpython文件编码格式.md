# q
Python文件出现“SyntaxError: Non-ASCII character”错误的原因和解决方法是什么？
# a
原因：文件中包含非ASCII字符（如中文），但没有声明编码格式。解决方法：在文件头部添加编码声明：
```
# -*- coding: utf-8 -*-
```

