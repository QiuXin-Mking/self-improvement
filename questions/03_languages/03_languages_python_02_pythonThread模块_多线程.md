# q
Python中使用threading模块创建并启动线程的基本步骤是什么？
# a
创建`threading.Thread`对象，指定目标函数`target`和参数`args`（或`kwargs`），然后调用`start()`方法启动线程。
```python
import threading
t = threading.Thread(target=func, args=(arg1, arg2))
t.start()
```

# q
什么是守护线程（daemon thread）？如何设置？它与非守护线程有何区别？
# a
守护线程是在主线程结束时自动退出的线程。设置方式：在`start()`之前调用`t.setDaemon(True)`（或`t.daemon = True`）。  
- 非守护线程：主线程结束后仍会继续运行，直到自身任务完成。  
- 守护线程：主线程结束时立即终止，无需等待其执行完毕。

# q
使用`args`参数向线程传递单个参数时，为什么会导致`TypeError`？如何正确传递？
# a
因为`args`期望接收一个可迭代对象（如元组），直接传递单个值（如整数）会被视为不可迭代。正确做法是在值后加一个逗号构成单元素元组：
```python
# 错误
threading.Thread(target=func, args=(value))
# 正确
threading.Thread(target=func, args=(value,))
```

# q
`join()`方法的作用是什么？在案例1中，为什么需要对所有线程调用`join()`？
# a
`join()`使调用线程（通常是主线程）阻塞，直到被调用线程执行完毕。  
在案例1中，主线程需要等待所有子线程（对各个节点执行操作）全部完成，才能准确打印“所有服务操作完成”的信息，因此对每个线程调用`join()`来同步。

