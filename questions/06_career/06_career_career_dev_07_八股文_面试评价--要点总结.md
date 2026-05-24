# q
new和malloc的区别是什么？
# a
- `new` 调用构造函数，`malloc` 仅分配内存。
- `new[]` 需对应 `delete[]`，`malloc` 用 `free` 释放。
- `new` 失败抛异常，`malloc` 返回 `NULL`。

# q
C++中vector的扩容机制是怎样的？
# a
默认按2倍扩容（部分实现1.5倍），扩容时申请新内存、拷贝旧数据、释放旧内存，时间复杂度分摊O(1)。

# q
零拷贝技术如何减少数据拷贝？
# a
通过 `mmap` 将文件映射到用户空间，绕过内核缓冲区；`sendfile` 直接在内核完成数据传输。

# q
如何用gdb调试多进程程序？
# a
使用 `set follow-fork-mode child` 跟踪子进程，`set detach-on-fork off` 保持父子进程调试。

