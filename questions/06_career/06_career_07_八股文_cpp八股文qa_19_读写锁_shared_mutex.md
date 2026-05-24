# q
std::shared_mutex 的读写并发规则是什么？
# a
std::shared_mutex 允许多个线程同时持有共享锁（读锁），但独占锁（写锁）只能由一个线程持有，且写锁与其他任何锁互斥。作为对比，std::mutex 同一时刻仅允许一个线程持有锁。

# q
std::recursive_mutex 适用于什么场景？
# a
当需要同一线程对同一个互斥量多次加锁（递归加锁）时，应使用 std::recursive_mutex。如果无递归需求，优先使用普通 std::mutex 以降低开销。

# q
如何使用 std::lock_guard 管理互斥锁？
# a
std::lock_guard 基于 RAII​ 机制，构造时对互斥量加锁，析构时自动解锁；只能在同一线程内构造与析构，禁止拷贝。头文件为 `<mutex>`，典型用法是将临界区代码包裹在一个作用域块内，确保异常安全且不会忘记解锁。示例：
```cpp
#include <mutex>
std::mutex mtx;
int shared = 0;

void foo() {
    {
        std::lock_guard<std::mutex> lk(mtx);
        ++shared;  // 临界区
    } // 离开作用域自动解锁
    // 耗时操作放在不持锁的位置
}
```

