# q
使用 time.Ticker 实现心跳时，为什么必须调用 `defer ticker.Stop()`？
# a
`ticker.C` 是一个 `chan time.Time`，每到间隔时间就会收到当前时间。如果忘记调用 `ticker.Stop()`，内部定时器不会被释放，会导致 goroutine 泄漏。

# q
`go func() { ... }()` 中的括号 `()` 代表什么，加与不加 `go` 有什么区别？
# a
后面的括号 `()` 表示立即调用该匿名函数。加了 `go` 会用新 goroutine 异步并发执行；不加 `go` 就是在当前 goroutine 同步执行。

# q
下面代码有什么问题？如何修正？
```go
for i := 0; i < 3; i++ {
    go func() {
        fmt.Println(i)
    }()
}
```
# a
闭包直接捕获循环变量 `i`，所有 goroutine 共享同一个变量，循环结束时 `i` 可能已经变成 3，导致输出不可预测（常见结果是 3,3,3）。修正方法是将 `i` 作为参数传入匿名函数：
```go
for i := 0; i < 3; i++ {
    go func(n int) {
        fmt.Println(n)
    }(i)
}
```

# q
`defer func()` 通常在什么位置声明？为什么？
# a
一般在资源获取后立即声明，即“越早声明越好”。例如打开文件后立刻 `defer func() { f.Close() }()`。这样无论后面函数如何 return 或 panic，都能保证资源被释放。

