# q
如何通过C语言从 `/dev/urandom` 获取随机数？
# a
使用以下步骤：
1. 用 `open("/dev/urandom", O_RDONLY)` 打开随机数设备。
2. 用 `read()` 读取指定字节数，存入变量（如 `unsigned int`）。
3. 检查返回值和错误码，使用 `perror` 输出错误信息。
4. 使用完毕后 `close()` 文件描述符。

示例代码：
```c
int randomData = open("/dev/urandom", O_RDONLY);
if (randomData < 0) {
    perror("open");
    return 1;
}

if (read(randomData, &myRandomNumber, sizeof(myRandomNumber)) < 0) {
    perror("read");
    return 1;
}

close(randomData);
printf("Random number: %u\n", myRandomNumber);
```

# q
`/dev/urandom` 与 `/dev/random` 的主要区别是什么？
# a
从这段代码本身无法直接体现区别，但 `/dev/urandom` 是非阻塞的，即使系统熵池不足也会返回伪随机数；`/dev/random` 在熵池不足时会阻塞。本示例使用 `/dev/urandom` 确保不会阻塞。

# q
示例中如何保证读取到正确大小的随机数据？
# a
通过 `read(randomData, &myRandomNumber, sizeof(myRandomNumber))` 读取。它指定了读取的字节数等于变量 `myRandomNumber` 的字节大小（对于 `unsigned int` 通常是 4 字节）。如果读取成功且返回值等于 `sizeof(myRandomNumber)`，则获得了完整的数据。但代码仅检查了 `< 0` 的错误情况，未严格判断读取字节数是否等于请求数，严格情况下应检查返回值是否等于 `sizeof(myRandomNumber)`。

