# q
如何在类的内部定义一个名为 f 的方法，要求入参 fname 的类型强制约定为 str？
# a
```python
def f(self, fname: str):
```

# q
如何在类的外部定义一个名为 f 的函数，要求入参 fname 的类型强制约定为 str？
# a
```python
def f(fname: str):
```

