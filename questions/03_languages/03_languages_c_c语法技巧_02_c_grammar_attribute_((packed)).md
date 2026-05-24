# q
如何取消C语言结构体的字节对齐？
# a
在结构体声明末尾加上 `__attribute__((packed))` 关键字，即可让结构体按照紧凑排列的方式占用内存，取消默认的字节对齐。

# q
请举例说明 `__attribute__((packed))` 对结构体大小的实际影响。
# a
示例结构体定义如下：
```c
struct Student {
    int age;
    char c;
} __attribute__((packed));
```
其中 `int` 通常占4字节，`char` 占1字节，由于取消了字节对齐，整个结构体的大小为5字节（通过 `sizeof` 输出为5），而不是对齐后的8字节。

