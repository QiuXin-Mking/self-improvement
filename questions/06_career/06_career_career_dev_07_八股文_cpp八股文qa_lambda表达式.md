# q
lambda 表达式 怎么写
写个func 函数，求 a+b;
# a
auto func = [](int a, int b) { return a + b; };
std::cout << func(2, 3) << std::endl; // 输出5

# q
定义一个最简单的 Lambda，不接受参数，也不返回值
输出: Hello, World!
函数名称 greet

1 定义
2 调用
# a
auto greet = []() {
    std::cout << "Hello, World!" << std::endl;
};

greet(); // 输出: Hello, World!

# q
定义一个接受两个参数并返回它们之和的 
输出: 返回两数字之和
函数名称 add
# a
auto add = [](int a, int b) -> int {
    return a + b;
};

# q
Lambda 表达式是 C++11 引入的一种创建匿名函数对象的简洁方式。 
抽象语法表达式
# a
[捕获列表](参数列表) -> 返回类型 {
    // 函数体
}

# q
介绍如下语法的
[x]
[&x]
[=]
[&]
[=,&x]
[&,x]
# a
值捕获​         [x]  在 Lambda 创建时，将 x的副本捕获进来。之后外部 x的变化不会影响 Lambda 内的 x。  
引用捕获​       [&x]  捕获 x的引用。Lambda 内对 x的修改会直接影响外部的 x。  
隐式值捕获​     [=]  以值捕获的方式捕获所有外部变量。  
隐式引用捕获​   [&]  以引用捕获的方式捕获所有外部变量。  
混合捕获​       [=, &x]  大部分变量值捕获，但 x是引用捕获。  
                [&, x]  大部分变量引用捕获，但 x是值捕获。  

不捕获​  []  不捕获任何外部变量。

