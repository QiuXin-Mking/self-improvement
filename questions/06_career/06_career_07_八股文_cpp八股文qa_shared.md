# q
计算下面的打印
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
    // 创建一个 shared_ptr，指向动态分配的 MyClass 实例
    std::shared_ptr<MyClass> p1 = std::make_shared<MyClass>();
    std::cout << "p1 use_count = " << p1.use_count() << std::endl;
    p1->greet();

    {
        // 拷贝p1，p2和p1均指向同一对象
        std::shared_ptr<MyClass> p2 = p1;
        std::cout << "Copy p1 to p2, p1 use_count = " << p1.use_count() << std::endl;

        useShared(p2); // 作为参数传递
    } // p2 离开作用域，引用计数减1

    std::cout << "After p2 out of scope, p1 use_count = " << p1.use_count() << std::endl;

    // p1 离开作用域，引用计数减为0，自动释放内存
    return 0;
}
# a
MyClass constructed
p1 use_count = 1
Hello from MyClass!
Copy p1 to p2, p1 use_count = 2
useShared received shared_ptr, use_count = 3
Hello from MyClass!
After p2 out of scope, p1 use_count = 1
MyClass destroyed

# q
为什么推荐用std::make_shared<MyClass>()而不是直接 std::shared_ptr<MyClass>(new MyClass) 呢？
# a
1.效率高
    make_shared<T>会一次性只分配一块内存，既包含你的对象（T），也包含引用计数控制块。
    直接用shared_ptr<T>(new T)则会执行两次内存分配：一次分配你的对象，一次分配控制块，这会增加内存碎片和分配开销。
2.防止内存泄露
    用std::shared_ptr<T>(new T(…))时，如果new T构造过程中抛异常，没来的及交给shared_ptr管理，有泄漏风险。
    make_shared内部能确保内存和对象要么都构造成功，要么都不留残渣，更安全。
3.可读性强

