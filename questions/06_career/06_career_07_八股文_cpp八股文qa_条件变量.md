# q
C++ 中条件变量的标准库拼写方式是什么？
# a
```cpp
std::condition_variable
```

# q
C++ 条件变量提供的等待与通知原语有哪些？
# a
等待原语：
- `wait`
- `wait_for`
- `wait_until`

通知原语：
- `notify_one`
- `notify_all`

# q
`std::condition_variable::wait` 如何工作，锁如何处理？
# a
调用 `wait` 时，线程进入等待状态的同时**原子地释放锁**；当线程被唤醒并返回时，**原子地重新获取锁**。

# q
`std::condition_variable::wait` 的 `lock` 参数应该使用什么类型？
# a
必须使用 `std::unique_lock<std::mutex>`，例如：
```cpp
std::unique_lock<std::mutex> lock(mtx);
```

# q
如何使用带谓词的 `wait` 避免虚假唤醒？写出等待队列非空的条件。
# a
```cpp
cond_var.wait(lock, [&]{ return !q.empty(); });
```

