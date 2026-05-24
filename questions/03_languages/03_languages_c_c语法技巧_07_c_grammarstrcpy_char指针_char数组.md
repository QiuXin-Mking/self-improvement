# q
在 C 语言中，为什么不能用 `=` 直接将 `char *` 类型的字符串赋给字符数组成员？
# a
当结构体成员是字符数组（如 `char devname[N]`）时，数组名是常量，不能作为赋值左值；即使是字符指针，直接使用 `=` 也只是复制指针地址，不会复制字符串内容，可能导致指针悬挂或意外修改原串。正确做法是使用 `strncpy` 或类似函数将字符串内容复制到数组缓冲区中，并指定目标缓冲区大小防止溢出。

# q
`strncpy(req_msg->devname, devname, sizeof(req_msg->devname))` 和 `req_msg->devname = devname` 的核心区别是什么？
# a
- `strncpy` 会将 `devname` 指向的字符串内容（最多指定长度）复制到 `req_msg->devname` 数组的存储空间中，属于**值复制**。
- `req_msg->devname = devname` 只是将指针 `devname` 的地址赋给 `req_msg->devname`（如果 `req_msg->devname` 是指针），改变的是指向，**不复制字符串数据**。若 `req_msg->devname` 是数组，该赋值会引发编译错误。

