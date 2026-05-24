# q
Python中的修饰符（装饰器）是什么？
# a
修饰符是Python中一种向已有代码中添加功能的特性，本质上是**高阶函数**。它接受一个函数作为参数，并返回一个包装后的新函数（或可调用对象），从而在不修改原函数代码的情况下扩展其行为。使用`@decorator_name`语法，放在被修饰函数定义的前一行。

# q
如何理解Python中的高阶函数？
# a
高阶函数是指**把其他函数当作参数**或**返回函数**的函数。在Python中，一切皆是对象，函数也能赋值给变量、作为参数传递。例如：
```python
def add(x): return x + 1
def test(func, x): return func(x)
test(add, 2)  # 结果为3
```

# q
使用修饰符时，如何处理被装饰函数的参数？
# a
在修饰符的内部嵌套函数中使用`*args, **kwargs`接收任意参数，并在调用原函数时透传这些参数。示例：
```python
def smart_divide(func):
    def inner(a, b):
        if b == 0:
            print("Whoops! cannot divide")
            return
        return func(a, b)
    return inner
```
通用写法：
```python
def decorator(func):
    def inner(*args, **kwargs):
        # 前置操作
        result = func(*args, **kwargs)
        # 后置操作
        return result
    return inner
```

# q
多个修饰符的执行顺序是什么？如何用非修饰符语法等价表示？
# a
多个修饰符从**下往上**依次执行。例如：
```python
@star
@percent
def printer(msg): print(msg)
```
等价于：
```python
printer = star(percent(printer))
```
执行时先由`percent`包装，再由`star`包装，输出效果为star的打印在最外层。

