# q
在Ubuntu上使用cppcoro库编写并编译C++协程程序需要满足哪些编译器条件？
# a
需要支持C++20的编译器，例如g++ 10及以上版本，编译时需指定`-std=c++20`选项。推荐使用Ubuntu 22.04或更新系统。

# q
给出编译使用cppcoro的C++程序的完整命令，包含链接库和路径指定。
# a
```
g++ -std=c++20 -I/usr/local/include -L/usr/local/lib my_fibonacci.cpp -lcppcoro -pthread -o my_fibonacci
```
其中`-I/usr/local/include`指定cppcoro头文件路径，`-L/usr/local/lib`指定库文件搜索路径，`-lcppcoro`链接cppcoro库，`-pthread`链接线程库。

# q
运行cppcoro程序时如果报错`error while loading shared libraries: libcppcoro.so`，应如何解决？
# a
设置动态链接库搜索路径，将cppcoro的安装目录添加到`LD_LIBRARY_PATH`环境变量：
```
export LD_LIBRARY_PATH=/usr/local/lib:$LD_LIBRARY_PATH
```
再运行程序。也可将该命令写入`~/.bashrc`以永久生效。

