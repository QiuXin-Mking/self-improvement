# q
iovec是什么？它的结构体包含哪些成员？
# a
iovec是I/O向量（I/O vector），配合readv和writev使用，可实现一次系统调用处理多个缓冲区的数据。其结构体定义在`<sys/uio.h>`，包含两个成员：`void *iov_base`（指向缓冲区起始地址）和`size_t iov_len`（缓冲区字节大小）。使用时通常以`struct iovec *iov`指针形式传递，并给出元素个数count。

# q
writev函数的核心用法是什么？如何用iovec实现聚集写？
# a
writev用于将多个分散缓冲区的数据一次性写入文件描述符（聚集写）。使用时先创建一个iovec数组，为每个元素的`iov_base`指向各自的缓冲区，`iov_len`设为对应长度，然后调用`writev(fd, iov, count)`，内核会按顺序将所有缓冲区的数据写入文件。示例中，三个字符串通过一个writev调用原子性地写入文件。

# q
readv函数的核心用法是什么？如何用iovec实现散布读？
# a
readv用于将文件描述符中的数据一次性分散读取到多个用户缓冲区（散布读）。先创建多个缓冲区并定义iovec数组，将`iov_base`指向各缓冲区，`iov_len`设为缓冲区大小，调用`readv(fd, iov, count)`后，内核按顺序将数据填入这些缓冲区。示例中，一次readv调用把文件内容分别读入foo、bar、baz三个字符数组。

