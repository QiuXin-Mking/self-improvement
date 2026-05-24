# q
scnprintf 是什么？它的主要作用是什么？
# a
scnprintf 是 Linux 内核中的安全字符串格式化函数，用于将格式化数据写入指定缓冲区，防止越界，并返回实际写入的字节数（不包含结尾的 `\0`）。它通常用于内核代码、驱动开发和 `/proc`、`/sys` 等接口的数据输出格式化。

# q
scnprintf 与 snprintf 的核心区别是什么？
# a
两者返回值行为不同：
- snprintf 返回“本来要写的长度”，即使超出缓冲区也会返回完整格式化后的字符数。
- scnprintf 返回“实际写入缓冲区的字节数”，最多写入 `size-1` 个字符，确保以 `\0` 结尾，杜绝缓冲区溢出。

# q
如何在 Linux 内核模块中正确使用 scnprintf？
# a
```c
#include <linux/kernel.h>

char buf[32];
int len = scnprintf(buf, sizeof(buf), "temp:%.2f,humi:%.2f", 22.3, 44.5);
// buf 内容为 "temp:22.30,humi:44.50"，len 为实际写入长度（≤31）
```
使用时需包含 `<linux/kernel.h>`，缓冲区和长度参数应匹配，返回值可用来判断实际写入量。

# q
scnprintf 适用于哪些编程环境？在用户空间可以使用吗？
# a
scnprintf 是内核态接口，适用于 Linux 内核模块、驱动开发和部分 RTOS 环境。用户空间没有 scnprintf，应使用标准 C 库的 snprintf 进行安全的字符串格式化。

