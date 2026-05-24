# q
C++20 协程是什么？它通过哪三个关键字实现暂停和恢复？
# a
C++20 协程是一种可以在函数执行过程中暂停和恢复的函数结构，用于异步编程和协作式多任务。它通过 `co_await`、`co_yield` 和 `co_return` 三个关键字实现暂停和恢复。

# q
协程的 `promise_type` 需要提供哪些关键方法？各自作用是什么？
# a
通常需要提供以下方法：
- `initial_suspend()`：控制协程开始执行后是否立即挂起，返回 `std::suspend_always` 表示让调用者控制启动。
- `final_suspend() noexcept`：协程结束时挂起，以便资源清理。
- `yield_value(T value)`：处理 `co_yield`，保存产出的值并挂起。
- `return_void()` 或 `return_value()`：处理 `co_return`。
- `unhandled_exception()`：异常处理。
- `get_return_object()`：返回协程关联的返回对象（如 `Generator`）。

# q
在示例 `Generator` 中，`next()` 成员函数如何工作？它如何判断协程是否结束？
# a
`next()` 调用 `coro.resume()` 恢复协程执行，然后返回 `!coro.done()`。`coro.done()` 在协程结束后返回 `true`，因此当协程尚未结束时 `next()` 返回 `true`，允许继续迭代。

# q
协程与线程在调度方式上有何根本区别？
# a
协程是协作式、非抢占的，需要主动通过 `co_await`、`co_yield` 等让出控制权；线程是抢占式调度的，操作系统可以在任意时刻切换线程，程序员无法控制切换时机。

