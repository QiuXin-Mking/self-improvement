# q
ZeroDivisionError 是什么 / 如何理解它
# a
`ZeroDivisionError` 是 Python 的内建异常类，表示除法操作中除数为零时引发的错误。当执行 `10 / 0` 等除数为零的运算时，Python 会抛出该异常。

# q
如何在使用除法时捕获 ZeroDivisionError
# a
使用 `try...except` 语句捕获：

```python
try:
    result = 10 / 0
except ZeroDivisionError:
    print("除以零错误发生！")
```

`ZeroDivisionError` 属于标准库异常，无需额外导入。

