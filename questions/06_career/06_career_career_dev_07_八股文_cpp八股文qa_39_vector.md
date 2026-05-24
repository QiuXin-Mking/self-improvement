# q
vector<int>a 的数组大小如何得出？
# a
```cpp
a.size();
```

# q
vector<int>&&a = b; 这里的&&代表什么意思？
# a
`vector<string>&&`中的双 `&&` 符号表示 右值引用（Rvalue Reference），这是 C++11 引入的关键特性，用于实现移动语义和完美转发。  
移动语义：通过右值引用，可以直接“窃取”临时对象的资源（如内存指针），避免深拷贝，提升性能。

# q
vector<int>a[2]; 与 vector<int>a(2); 的区别是什么？
# a
- `vector<int> a[2];` 定义了 2 个 `vector<int>` 对象 `a[0]` 和 `a[1]`，每个都是独立的 vector，可以单独操作，例如：
  ```cpp
  a[0].push_back(10);
  a[1].push_back(20);
  ```
- `vector<int> a(2);` 定义了一个 `vector<int>` 对象 `a`，包含 2 个整型元素，初值均为 0，例如：
  ```cpp
  a[0] = 100;
  a[1] = 200;
  ```
  也可以写成 `vector<int> a = {0, 0};`

# q
vector<int> a(2); 会不会初始化为0？
# a
会初始化，且为 0。等价于 `vector<int> a = {0, 0};`，容量为 2，默认初始化为 0。

