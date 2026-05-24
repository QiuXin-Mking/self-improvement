# q
如何获取系统内存页面大小？
# a
使用 `getconf PAGESIZE` 命令，通常在大多数现代系统上返回 **4096** 字节（即 **4KB**）。

# q
什么是 Hugepagesize，它有什么作用？
# a
Hugepagesize 是 Linux 内核特性 HugePages 所使用的大页面大小，通常为 **2MB**（2048KB）。使用更大的页面可以减小页表大小、降低 TLB 失效频率，从而提升需要大量内存的大型应用程序（如数据库）的性能。

# q
如何查看当前系统的 Hugepagesize 值？
# a
执行以下命令，读取 `/proc/meminfo` 文件中对应的行：
```
grep -i Hugepagesize /proc/meminfo
```

