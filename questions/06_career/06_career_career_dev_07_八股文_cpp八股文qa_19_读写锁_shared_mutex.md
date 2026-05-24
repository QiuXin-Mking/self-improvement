# q
std::mutex的特点有哪些？
# a
同一时刻仅允许一个线程持有锁；只有加锁者能解锁，不可重入、不可重复加锁。

# q
std::shared_mutex与std::mutex的主要区别是什么？
# a
std::shared_mutex允许多个线程并发读，但写操作需要独占锁；std::mutex同一时刻只允许一个线程持有锁。

# q
std::recursive_mutex的使用场景是什么？
# a
需要同一线程重复加锁时使用std::recursive_mutex，无递归需求时优先使用普通mutex以降低开销。

# q
std::lock_guard如何使用？有何特点？
# a
基于RAII机制，构造时对互斥量加锁，析构时自动解锁。只能在同一线程内构造与析构，禁止拷贝。典型用法是将临界区限制在一个作用域块内，保证异常安全且不会忘记解锁。头文件为<mutex>。代码示例：
```cpp
#include <mutex>
std::mutex mtx;
int shared = 0;

void foo() {
    {
        std::lock_guard<std::mutex> lk(mtx);
        ++shared;  // 临界区
    } // 自动解锁
    // 耗时操作放在此处，不持锁
}
```

