# q
如何实现基本的 memcpy 函数？
# a
```c
void* memcpy(void* dest, const void* src, size_t n)
{
    unsigned char* pdest = (unsigned char*)dest;
    const unsigned char* psrc = (const unsigned char*)src;
    while(n--)
        *pdest++ = *psrc++;
    return dest;
}
```

# q
如何实现考虑内存重叠的 safe_memcpy 函数？
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

# q
在 safe_memcpy 中，当 dest 在 src 之后且内存重叠时，采用什么复制顺序避免覆盖源数据？
# a
采用从后向前的顺序复制（递减循环：for (i = n; i != 0; --i) d[i-1] = s[i-1]），确保先复制尾部再复制头部，避免源数据在复制过程中被覆盖。

