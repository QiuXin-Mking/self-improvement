# q
C++中new和malloc的区别是什么？
# a
- new会调用构造函数，malloc仅分配内存。
- new[]需对应delete[]，malloc用free释放。
- new分配失败抛出异常，malloc返回NULL。

# q
C++的vector扩容机制是怎样的？
# a
默认按2倍扩容（部分实现如MSVC采用1.5倍），扩容时申请新内存、拷贝旧数据、释放旧内存，时间复杂度均摊O(1)。

# q
C++智能指针shared_ptr的循环引用问题如何产生和解决？
# a
当两个shared_ptr互相引用时会导致引用计数永不为0，造成内存泄漏。解决方法是将其中一个改用weak_ptr，weak_ptr不增加引用计数且能检测对象是否有效。

# q
如何使用gdb调试多进程程序？
# a
使用`set follow-fork-mode child`跟踪子进程，`set detach-on-fork off`保持父子进程同时调试，结合`bt`和`frame`查看调用栈。

# q
零拷贝技术如何减少数据拷贝？
# a
通过`mmap`将文件映射到用户空间，绕过内核缓冲区拷贝；或使用`sendfile`系统调用直接在内核完成数据传输，避免用户态与内核态之间的多次拷贝。

