# q
new和malloc的区别是什么？
# a
1. new 调用构造函数，malloc只分配内存。
2. new失败时候抛出异常，malloc失败时候返回NULL。
3. new对应释放用del ,malloc释放用free。del调用析构函数，free释放内存片。

# q
被free回收的内存是立即返还给操作系统吗？为什么
# a
被free回收的内存会首先被ptmalloc使用双链表保存起来，当用户下一次申请内存的时候，会尝试从这些内存中寻找合适的返回。
这样就避免了频繁的系统调用，占用过多的系统资源。

