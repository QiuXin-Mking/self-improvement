# q
RAII原理是啥
# a
RAII（Resource Acquisition Is Initialization，资源获取即初始化）是 C++ 中管理资源的核心编程范式，其本质是通过对象的生命周期自动控制资源的获取与释放，从而避免资源泄漏和异常安全问题。

​资源获取即初始化​：在对象的构造函数中获取资源（如内存、文件句柄、锁等）。

​资源释放即析构​：在对象的析构函数中释放资源。

​作用域控制​：对象作为局部变量存在时，其生命周期由作用域决定，离开作用域时自动调用析构函数。

# q
RAII在C++中有哪些典型应用？
# a
- 智能指针：std::unique_ptr、std::shared_ptr 管理动态内存。
- 锁管理：std::lock_guard、std::unique_lock 自动加锁/解锁。
- 文件流：std::ifstream、std::ofstream 自动关闭文件。

