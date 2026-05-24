# q
在使用memset时出现"passing argument 2 of ‘memset’ makes integer from pointer without a cast"错误，原因是什么？如何修正？
# a
`memset`的第二个参数应该是`int`类型的值（用于填充内存的字节值），而代码中错误地传入了指针`NULL`。应将`NULL`改为`0`（以零填充），并注意第三个参数应是对应数组的总字节数（如`sizeof(skiplist_t*) * 32`而非`sizeof(skiplist_t)`）。修正示例：
```c
skiplist_t *update_list[32] = {NULL};
memset(update_list, 0, sizeof(update_list));
```

# q
C语言编译时出现"implicit declaration of function"错误是什么原因？如何解决？
# a
该错误是因为在调用函数之前没有提供函数的声明或定义，编译器隐式地假定该函数返回`int`类型，导致类型不匹配告警。解决方法是在调用前添加函数声明（或包含对应的头文件），例如：
```c
extern void osd_repair_seg_ctx_free(osd_repair_seg_ctx_t* ctx);
```

