# q
在C语言中，声明 `osd_mdlog_mgt_t *src_omdlog = NULL, dst_omdlog == NULL;` 时，`dst_omdlog` 的实际类型是什么？为什么容易出错？
# a
`dst_omdlog` 的类型是 `osd_mdlog_mgt_t`，而不是指针。因为 `*` 只作用于紧随其后的变量名 `src_omdlog`，不会延续到逗号声明的下一个变量。这种写法容易让人误以为 `dst_omdlog` 也是指针，属于典型的声明陷阱。

# q
如何安全地编写多个指针变量的声明以避免类似 `osd_mdlog_mgt_t *src_omdlog = NULL,dst_omdlog == NULL;` 的陷阱？
# a
每个指针变量单独使用一行声明并立即初始化，例如：
```c
osd_mdlog_mgt_t *src_omdlog = NULL;
osd_mdlog_mgt_t *dst_omdlog = NULL;
```
避免将指针和非指针或不同初始化的变量混在同一行声明。

