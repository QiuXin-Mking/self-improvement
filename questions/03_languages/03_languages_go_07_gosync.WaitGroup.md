# q
sync.WaitGroup 是什么？它的核心作用是什么？
# a
`sync.WaitGroup` 是 Go 标准库 `sync` 包提供的并发原语，用于等待一组 goroutine 完成执行。它让一个 goroutine（通常是主 goroutine）可以阻塞，直到所有被跟踪的 goroutine 都完成工作，计数器归零后继续运行。

# q
如何使用 sync.WaitGroup 同步一组 goroutine？涉及哪些关键方法？
# a
基本用法分为三步：
1. 在启动每个 goroutine 之前调用 `wg.Add(1)` 增加计数器。
2. 将 `*sync.WaitGroup` 指针传入 goroutine，在其内部通过 `defer wg.Done()` 确保完成时计数器减 1。
3. 在需要等待的位置调用 `wg.Wait()` 阻塞，直到计数器变为 0。
示例：
```go
var wg sync.WaitGroup
for i := 1; i <= 3; i++ {
    wg.Add(1)
    go worker(i, &wg)
}
wg.Wait()
```

# q
使用 sync.WaitGroup 时有哪些关键注意事项？
# a
- `Add` 必须在启动 goroutine 前调用，且计数要与实际启动的 goroutine 数量一致。
- `Done` 应在 goroutine 内通过 `defer wg.Done()` 调用，避免因提前返回导致计数不匹配。
- `WaitGroup` 必须通过指针传递给 goroutine，否则会复制值导致计数器不同步。
- `Wait` 会阻塞当前 goroutine，直至计数器归零，程序才能继续执行。

