# q
cpp 的条件变量拼写方式是什么？
# a
std::condition_variable

# q
cpp 的条件变量提供哪两类条件原语？具体有哪些函数？
# a
等待：wait、wait_for、wait_until
通知：notify_one、notify_all

# q
cpp 中 std::condition_variable 的 wait 函数的工作原理是怎样的？
# a
进入等待状态时释放锁，结束等待时获取锁，两个操作都是原子的。

# q
wait 函数中需要传入什么类型的 lock？如何定义并使用带谓词的 wait 来避免伪唤醒？
# a
使用 `std::unique_lock<std::mutex> lock(mtx);`
然后 `cond_var.wait(lock, []{ return condition; });` 带谓词的 wait 可防止伪唤醒。

# q
如果等待队列非空，如何书写 cond_var.wait(lock, ...) 的谓词？
# a
`cond_var.wait(lock, [&]{ return !q.empty(); });`

