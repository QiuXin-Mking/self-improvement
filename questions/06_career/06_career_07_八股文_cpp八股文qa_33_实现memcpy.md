# q
如何实现基本的 memcpy 函数？
# a
```c
void* memcpy(void* dest, const void* src, size_t n)
{
    // 将指针转换为unsigned char*，以字节为单位处理
    unsigned char* pdest = (unsigned char*)dest;
    const unsigned char* psrc = (const unsigned char*)src;

    // 一字节一字节地复制
    while(n--)
        *pdest++ = *psrc++;

    return dest;
}
```

# q
如何实现能处理内存重叠的 safe_memcpy？
# a
```c
void* safe_memcpy(void* dest, const void* src, size_t n) {
    unsigned char* d = (unsigned char*)dest;
    const unsigned char* s = (const unsigned char*)src;
    if (d == s || n == 0) return dest;
    if (d > s && d < s + n) { // 有重叠，且dest在src之后
        for (size_t i = n; i != 0; --i)
            d[i - 1] = s[i - 1];
    } else { // 无重叠或dest在src前面
        for (size_t i = 0; i < n; ++i)
            d[i] = s[i];
    }
    return dest;
}
```

