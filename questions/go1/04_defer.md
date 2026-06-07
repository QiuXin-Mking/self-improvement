# q
多个 defer 的执行顺序是什么？

# a
后进先出（LIFO），像一个 defer 栈。

```go
defer fmt.Println("1") // 第三个执行
defer fmt.Println("2") // 第二个执行
defer fmt.Println("3") // 第一个执行
// 输出: 3 2 1
```

# q
```go
func demo1() {
    x := 1
    defer fmt.Println(x)
    x = 2
}
// 输出什么？
```

# a
输出 1。

defer 的参数在 defer 声明时求值，不是执行时。声明时 x=1，所以 defer 记住了 1。

但如果传的是闭包（函数体），则不同。

# q
```go
func demo2() {
    x := 1
    defer func() { fmt.Println(x) }()
    x = 2
}
// 输出什么？
```

# a
输出 2。

和上一题不同：这里 defer 的是一个闭包，闭包内部引用的是变量 x（不是值拷贝）。执行时 x 已经变成 2 了。

面试核心区分：
- defer 直接传参 → 声明时求值
- defer 闭包引用 → 执行时取值

# q
```go
func f1() int {
    x := 5
    defer func() { x++ }()
    return x
}
// 返回什么？
```

# a
返回 5。

执行过程：
1. return x → 把 x 的值 5 赋给返回值变量（匿名）
2. defer 执行，x++（x 变成 6，但返回值已经是 5 了）
3. 函数返回 5

defer 可以修改命名返回值，但改不了匿名返回的"已赋值结果"。

# q
```go
func f2() (x int) {
    defer func() { x++ }()
    return 5
}
// 返回什么？
```

# a
返回 6。

有命名返回值 x 时：
1. return 5 → x = 5
2. defer 执行，x++ → x = 6
3. 函数返回 x，即 6

命名返回值让 defer 可以"看到并修改"返回值。这是 defer 最常考的面试题。

# q
```go
func f3() (x int) {
    defer func(x int) { x++ }(x)
    return 5
}
// 返回什么？和 f2 有什么区别？
```

# a
返回 5。

和 f2 的关键区别：defer 的闭包有参数 x（传值），参数在声明时求值——就是 x 当前的值 0（命名返回值的零值）。闭包内部 x++ 改的是闭包自己的参数副本，不影响外层返回值。

三个区分：
- `defer func() { x++ }()` — 闭包引用外层 x，执行时 x=5，x++ 后 x=6，返回 6
- `defer func(x int) { x++ }(x)` — 参数 x 求值时 x=0（声明时还没 return），x++ 改的是副本，返回 5

# q
defer 在什么时机执行？return 之前还是之后？

# a
defer 在 return 之后、函数真正返回调用方之前执行。

更精确的步骤：
1. 给返回值赋值
2. 执行 defer（LIFO 顺序）
3. 返回到调用方（RET 指令）

# q
```go
func demo3() {
    defer fmt.Println("A")
    panic("oops")
    defer fmt.Println("B") // 永远不会执行
}
// 输出什么？
```

# a
输出 A，然后 panic: oops。

panic 发生时，已经注册的 defer 仍然会执行（包括 recover），但 panic 之后的 defer 不会注册。

defer 的注册时机很重要——先 defer 的才能在 panic 时被调用。

# q
recover 放在什么位置才有效？

# a
recover 必须直接在 defer 函数内部调用才有效。

```go
// 有效
defer func() {
    if r := recover(); r != nil {
        fmt.Println("recovered:", r)
    }
}()

// 无效：recover 不在 defer 的直接函数里
defer func() {
    func() {
        recover() // 不生效！
    }()
}()

// 无效：recover 不在 defer 里
func main() {
    recover() // 永远是 nil
    panic("x")
}
```

recover 捕获的是当前 goroutine 的 panic，跨 goroutine 无效。

# q
defer 对性能有影响吗？什么时候不应该用 defer？

# a
有微弱影响（纳秒级），主要开销在函数调用和参数准备。绝大多数场景 negligible。

不适用 defer 的场景：
1. 循环体内高频操作（如循环内打开文件并 defer Close()）——应手动关闭
2. 性能极度敏感的热路径

```go
// 不好：循环内 defer 造成堆积
for _, f := range files {
    f := os.Open(f)
    defer f.Close() // 所有文件在循环结束后才关闭
}

// 好：用完即关
for _, f := range files {
    f := os.Open(f)
    process(f)
    f.Close()
}
```
