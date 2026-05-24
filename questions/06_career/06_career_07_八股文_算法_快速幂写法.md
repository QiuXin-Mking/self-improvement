# q
带取模的快速幂算法的完整实现代码是什么？
# a
```cpp
const int MOD = 1e9 + 7;

long long mod_pow(long long base, long long exp) {
    long long result = 1;
    while (exp > 0) {
        if (exp % 2 == 1) {
            result = (result * base) % MOD;
        }
        base = (base * base) % MOD;
        exp /= 2;
    }
    return result;
}
```

# q
快速幂实现中使用的模数 `MOD` 的值是多少？
# a
`1e9 + 7`，即 `const int MOD = 1e9 + 7;`

# q
快速幂算法循环内如何判断当前指数最低位是否为 1？
# a
通过 `if (exp % 2 == 1)` 判断，等价于检查指数二进制最低位。

