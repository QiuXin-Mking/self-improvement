# q
计算下列代码的输出结果
```cpp
#include <iostream>
#include <memory>

class MyClass {
public:
    MyClass()  { std::cout << "MyClass constructed\n"; }
    ~MyClass() { std::cout << "MyClass destroyed\n"; }
    void greet() { std::cout << "Hello from MyClass!\n"; }
};

void useShared(std::shared_ptr<MyClass> ptr) {
    std::cout << "useShared received shared_ptr, use_count = " << ptr.use_count() << std::endl;
    ptr->greet();
}

int main() {
    std::shared_ptr<MyClass> p1 = std::make_shared<MyClass>();
    std::cout << "p1 use_count = " << p1.use_count() << std::endl;
    p1->greet();

    {
        std::shared_ptr<MyClass> p2 = p1;
        std::cout << "Copy p1 to p2, p1 use_count = " << p1.use_count() << std::endl;

        useShared(p2);
    } // p2 离开作用域

    std::cout << "After p2 out of scope, p1 use_count = " << p1.use_count() << std::endl;
    return 0;
}
```
# a
```
MyClass constructed
p1 use_count = 1
Hello from MyClass!
Copy p1 to p2, p1 use_count = 2
useShared received shared_ptr, use_count = 3
Hello from MyClass!
After p2 out of scope, p1 use_count = 1
MyClass destroyed
```

# q
为什么推荐用 `std::make_shared<MyClass>()` 而不是直接 `std::shared_ptr<MyClass>(new MyClass)`？
# a
1. 效率高：`make_shared<T>` 一次性分配一块内存，同时容纳对象和引用计数控制块；而 `shared_ptr<T>(new T)` 需要两次分配，分别分配对象和控制块，增加内存碎片和分配开销。
2. 防止内存泄漏：若使用 `std::shared_ptr<T>(new T(...))`，在 `new T` 构造抛异常时，内存尚未交给 `shared_ptr` 管理，存在泄漏风险；`make_shared` 内部保证内存和对象要么都构造成功，要么不留残余，更安全。
3. 可读性强，代码更简洁。

