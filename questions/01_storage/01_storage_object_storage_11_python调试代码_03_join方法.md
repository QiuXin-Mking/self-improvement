# q
threading.Thread 的 join() 方法是什么？
# a
join() 方法用于等待一个线程完成其执行。调用线程会阻塞（暂停执行），直到被 join() 的线程执行完毕。

# q
join() 方法是否会终止目标线程？
# a
不会。join() 只是让调用线程等待目标线程自然结束，不会强制终止目标线程。

# q
使用 join() 时，调用线程的行为是什么？
# a
调用线程在 join() 处被阻塞，暂停执行后续代码，直到被等待的线程完成后才继续执行。例如：

```python
t = threading.Thread(target=worker)
t.start()
t.join()          # 主线程在此阻塞，直到 worker 完成
print("主线程结束")  # 在 worker 完成后才执行这一行
```

