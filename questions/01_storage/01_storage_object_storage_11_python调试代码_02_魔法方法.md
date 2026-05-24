# q
Python中`__init__`方法的作用是什么？
# a
`__init__` 是一个特殊方法（魔法方法），在创建类的新实例时自动调用，用于初始化对象的属性。例如：
```python
class MyClass:
    def __init__(self, value):
        self.value = value
obj = MyClass(10)
print(obj.value)  # 10
```

# q
`__init__` 和 `init` 在Python中有何区别？
# a
- `__init__`：双下划线方法，对象实例化时自动调用，用于初始化。
- `init`：普通方法，无特殊含义，需要显式定义并手动调用。如果像`__init__`那样直接传参给类创建实例，会因`__init__`未定义而抛出`TypeError`。为避免混淆，通常不会将普通方法命名为`init`。

