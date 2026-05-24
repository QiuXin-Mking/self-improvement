# q
算法竞赛模板中定义的快速读入函数 `inline int read()` 是如何实现的？
# a
```cpp
inline int read()
{
	int x=0,y=1;char a=getchar();
	while ( a>'9' || a<'0'){if ( a=='-')y=-1;a=getchar();}
	while ( a>='0' && a<='9' ){	x=10*x+a-'0'; a=getchar();}
	return x*y;
}
```
其原理是：先跳过非数字字符并处理负号（y记录正负），再连续读取数字字符并累加到x（x = 10*x + 当前数字），最后返回带符号的整数。这比cin/scanf更快。

# q
模板中 `#define inf 0x3f3f3f3f` 的含义是什么？为什么选择这个值？
# a
`inf` 是算法竞赛中常用的无穷大常量，值为 `0x3f3f3f3f`（十进制约 1,061,109,567）。它的优势有：①足够大（10^9 数量级），通常数据不会达到；② `memset(a, 0x3f, sizeof(a))` 可将数组每个元素初始化为该值，且 `inf + inf` 不会溢出；③方便进行最短路等算法中的无穷大比较。

# q
宏 `#define rep(i,aa,bb)` 和 `#define rrep(i,aa,bb)` 的作用是什么？
# a
```cpp
#define rep(i,aa,bb) for(register int i=aa;i<=bb;i++)
#define rrep(i,aa,bb) for(register int i=aa;i>=bb;i--)
```
`rep` 用于正向遍历（i从aa递增到bb），`rrep` 用于反向遍历（i从bb递减到aa）。其中 `register` 关键字提示编译器将循环变量放入寄存器以优化速度。

# q
模板中的 `#define lowbit(x) x&(-x)` 有什么作用？
# a
`lowbit(x)` 返回 `x` 的二进制表示中最低位的 1 所对应的值。常用于树状数组（Fenwick Tree）操作，例如：
- lowbit(12) = 4 （12 的二进制 1100，最低位1在第3位，值为4）。

# q
`freopen("1.txt","r",stdin);` 和 `std::ios::sync_with_stdio(false);` 在算法竞赛中的作用是什么？
# a
- `freopen("1.txt","r",stdin);` 将标准输入重定向到文件 `1.txt`，方便本地测试时从文件读入数据，提交时注释掉即可。
- `std::ios::sync_with_stdio(false);` 关闭 C++ 输入输出流与 C 标准流的同步，可大幅提升 `cin/cout` 的速度，但此时不能混用 `cin` 与 `scanf`。

