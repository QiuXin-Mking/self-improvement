# q
用 var s []int 声明后，s 是什么？能直接 s[0]=1 吗？

# a
s 是 nil slice（len=0, cap=0），不能直接 s[0]=1 会 panic: index out of range。需要用 append 或 make 初始化。

slice 底层结构是 ptr + len + cap。nil slice 的 ptr 为空。append nil slice 完全合法，Go 会自动分配底层数组。

# q
```go
x := []int{1, 2, 3, 4, 5}
y := x[1:3]  // [2, 3]
y[0] = 100
// x 现在是什么？
```

# a
x = [1, 100, 3, 4, 5]

y 和 x 共享底层数组。y 只是 x 的一个"窗口"，修改 y[0] 就是修改底层数组的第 1 个位置。

切片的本质是数组的引用视图，不是拷贝。这点和 C++ 的 vector 完全不同——C++ 的 vector 之间是独立的数据，而 Go 切片共享底层。

# q
```go
z := make([]int, 2, 3)
z[0], z[1] = 10, 20
w := append(z, 30)
w[0] = 999
// z 变成什么？
```

# a
z = [999, 20]

z 的 cap=3，append 一个元素后 cap 还够，所以 w 和 z 共享底层数组。w[0] 就是 z[0]。

关键判断：append 后 len 是否超过 cap？超过 → 扩容分配新数组；不超过 → 共享原数组。

# q
```go
u := make([]int, 2) // len=2, cap=2
u[0], u[1] = 10, 20
v := append(u, 30)  // cap 不够
v[0] = 999
// u 变了吗？
```

# a
u = [10, 20]，不变。

cap 不够所以 append 扩容了，v 指向全新的底层数组，和 u 完全分离。

Go 1.18+ 扩容规则：新容量 < 256 → 翻倍； >= 256 → 约 1.25 倍。

# q
```go
arr := [4]int{10, 20, 30, 40} // 这是数组！
s := arr[:2]                   // len=2 cap=4
s = append(s, 50)              // 会怎样？
s = append(s, 60)
s = append(s, 70) // cap 不够了，扩容
// arr 最终是什么？
```

# a
arr = [10, 20, 50, 60]

两次 append(50, 60) 都在 cap=4 范围内，直接写到了 arr 的底层数组。第三次 append(70) 超过 cap，s 扩容指向新数组，不再影响 arr。

这是最容易出 bug 的场景：对数组切片后 append，不小心改了原数组。

# q
```go
s := make([]int, 0)
p := s
s = append(s, 1, 2, 3)
// p 呢？
```

# a
p = []，还是空的。

p 持有的是旧的 slice header（ptr=nil, len=0, cap=0）。append 后 s 的 header 变了（新的 ptr、len、cap），但 p 不受影响。

Go 的赋值是值拷贝，slice header 的 3 个字段（ptr, len, cap）被复制了一份给 p。

# q
```go
func modify(s []int) {
    s = append(s, 100)
}

func main() {
    s := make([]int, 1, 2)
    s[0] = 1
    modify(s)
    // s 是什么？len 变了吗？
}
```

# a
s = [1]，len 还是 1，没变。

slice 传参传的是 header 拷贝。modify 内部的 append 改变了内部变量 s 的 len，但 main 里的 s 的 len 还是 1。

但是！如果 cap 够，底层数组的第 2 个位置已经被写入了 100（只是 main 里的 s 看不到，因为 len=1）。如果把 main 的 s 重新切片 s[:2] 就能看到 100。

如果 cap 不够扩容了，那就连底层数组都不影响。

# q
```go
s := []int{1, 2, 3, 4, 5}
s = s[:3]
// len? cap?
```

# a
s = [1, 2, 3], len=3, cap=5

缩小 len 不会改变 cap。底层数组还是 5 个元素，只是"看不见"后两个了。重新 s = s[:5] 又能看到后面两个。

# q
nil slice 和 empty slice 的区别？

# a

| | nil slice | empty slice |
|---|---|---|
| 声明 | var s []int | s := []int{} 或 make([]int, 0) |
| ptr | nil | 有底层数组指针（非 nil） |
| len/cap | 0/0 | 0/0 |
| == nil | true | false |
| JSON | null | [] |
| append | 合法 | 合法 |

大多数场景下行为一致，但 JSON 序列化结果不同。用 len(s)==0 判断是否为空，别用 s==nil。

# q
for range slice 里怎么修改元素？直接改 value 变量行吗？

# a
不行。for range 里的 value 是元素拷贝，改它不影响原切片。

```go
// 错误
for _, v := range s {
    v = 100 // 白改
}

// 正确
for i := range s {
    s[i] = 100
}
```
