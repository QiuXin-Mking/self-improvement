# q
vector<int> a 的大小如何获取？
# a
调用 `a.size()` 成员函数。

# q
vector<int>& &&a = b; 中的 && 是什么含义？
# a
`&&` 表示右值引用（Rvalue Reference），C++11 引入，用于实现移动语义和完美转发。通过右值引用可以直接“窃取”临时对象的资源（如内存指针），避免深拷贝，提升性能。

# q
vector<int> a[2]; 和 vector<int> a(2); 有什么区别？
# a
`vector<int> a[2];` 定义了一个包含 2 个 `vector<int>` 的数组，`a[0]` 和 `a[1]` 是两个独立的 vector，可以通过 `push_back` 等操作各自添加元素。  
`vector<int> a(2);` 定义了一个 `vector<int>`，其初始包含 2 个元素（默认值为 0），可以通过 `a[0]`、`a[1]` 直接访问和修改这些元素。

# q
vector<int> a(2); 会不会自动初始化为 0？
# a
会，`vector<int> a(2);` 会创建包含 2 个元素的 vector，每个元素的值都会被初始化为 0。

# q
函数参数用 `string &str` 和 `string str` 有什么区别？
# a
使用引用 `string &str` 可以避免字符串拷贝，提升性能，并允许函数修改原字符串；但要求传入左值。  
使用值传递 `string str` 会产生一次拷贝，对原参数无影响，但可以接受右值（临时对象），结合移动语义也可能减少拷贝开销。

