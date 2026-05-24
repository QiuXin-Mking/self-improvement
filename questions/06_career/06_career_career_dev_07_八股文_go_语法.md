# q
使用 `time.NewTicker` 后为什么必须调用 `defer ticker.Stop()`？
# a
`ticker.C` 是一个 `chan time.Time`，定时触发时会发送当前时间。如果不 `defer ticker.Stop()`，内部定时器不会被释放，会导致 goroutine 泄漏。

```go
ticker := time.NewTicker(1 * time.Second)
defer ticker.Stop()
for {
    select {
    case t := <-ticker.C:
        fmt.Println("Tick at", t.Format("15:04:05"))
    case <-time.After(5 * time.Second):
        fmt.Println("5 秒后退出")
        return
    }
}
```

# q
`go func() { ... }()` 中的 `go` 关键字和末尾括号分别代表什么？
# a
`go` 关键字表示启动一个新 goroutine 来异步执行该匿名函数，末尾 `()` 表示立即调用这个匿名函数。如果不加 `go`，就是在当前 goroutine 同步执行；加了 `go`，就是并发执行。

```go
go func() {
    fmt.Println("子 goroutine 执行")
}()
```

# q
在 `for` 循环中直接用 goroutine 打印循环变量 `i` 会有什么问题？如何修复？
# a
直接使用 `i` 会形成闭包捕获共享变量，循环结束时所有 goroutine 看到的都是最终值（例如全部输出 `3`）。修复方法是通过函数参数传递当前值的副本。

```go
// 错误示范
for i := 0; i < 3; i++ {
    go func() {
        fmt.Println(i) // 危险：结果可能是 3,3,3
    }()
}

// 正确写法
for i := 0; i < 3; i++ {
    go func(n int) {
        fmt.Println(n) // 每个 goroutine 拿到当时的副本
    }(i)
}
```

# q
`defer func()` 应该在什么位置声明？
# a
应该在资源刚获取或函数一进入时立即声明，越早越好。这样无论函数中途如何 `return` 或发生 `panic`，`defer` 中的清理逻辑都能保证执行。

```go
f, err := os.Open(path)
if err != nil {
    return nil, err
}
defer func() {
    f.Close()
    fmt.Println("文件已关闭")
}()
// 后续业务逻辑...
```

