# q
如何使用 cppcoro 库创建协程生成器并输出斐波那契数列？
# a
使用 ```cppcoro::generator<int>``` 作为返回类型，在协程中使用 ```co_yield``` 暂停并产生值。示例：
```cpp
#include <cppcoro/generator.hpp>
cppcoro::generator<int> fibonacci() {
    int a = 0, b = 1;
    while (true) {
        co_yield b;
        auto next = a + b;
        a = b;
        b = next;
    }
}
```
调用时通过范围 for 循环遍历生成器即可逐一获取数值。

# q
在 Ubuntu 上如何安装 cppcoro 库？
# a
```bash
git clone https://github.com/lewissbaker/cppcoro.git
cd cppcoro
mkdir build && cd build
cmake ..
make -j
sudo make install
```
安装完成后头文件和库文件会被安装到系统目录（如 `/usr/local/include` 和 `/usr/local/lib`）。

# q
如何编译使用了 cppcoro 的 C++20 程序？
# a
使用支持 C++20 协程的编译器（g++ 10 以上），命令示例：
```bash
g++ -std=c++20 my_fibonacci.cpp -lcppcoro -pthread -o my_fibonacci
```
如果链接失败，可显式指定头文件与库路径：
```bash
g++ -std=c++20 -I/usr/local/include -L/usr/local/lib my_fibonacci.cpp -lcppcoro -pthread -o my_fibonacci
```

# q
运行 cppcoro 程序时出现找不到共享库的错误应如何处理？
# a
设置动态库搜索路径：
```bash
export LD_LIBRARY_PATH=/usr/local/lib:$LD_LIBRARY_PATH
```
然后重新运行可执行文件。

