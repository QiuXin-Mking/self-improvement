# q
goroutine 和操作系统线程的关系？GMP 模型是什么？

# a
GMP 模型是 Go 调度器的核心：

| | 角色 | 数量 |
|---|---|---|
| G (Goroutine) | 用户态协程 | 成千上万 |
| M (Machine) | OS 线程，真正执行代码 | 默认 CPU 核数 |
| P (Processor) | 逻辑处理器，管理 G 的运行队列 | GOMAXPROCS（默认 CPU 核数）|

关键设计：
- G 在 M 上运行，但通过 P 调度，实现 M:N 调度（M 个 goroutine 映射到 N 个 OS 线程）
- 每个 P 有本地 G 队列（无锁获取 G），全局也有 G 队列
- 当 G 阻塞（系统调用），P 会被转移到其他 M，原来的 M 带着阻塞的 G 等着
- 当 G 发起网络 IO，进入 netpoller（非阻塞 IO），G 被挂起不占 M

和 C++ 线程的区别：goroutine 初始栈只有 2KB（C++ 线程 2MB），创建/切换成本极低。

# q
```go
func main() {
    go fmt.Println("hello")
}
// 能看到输出吗？
```

# a
大概率看不到。main goroutine 结束，整个进程就退出了，不会等待其他 goroutine。

解决方案：
1. sync.WaitGroup
2. channel 同步
3. time.Sleep（不推荐，不确定）

```go
var wg sync.WaitGroup
wg.Add(1)
go func() {
    defer wg.Done()
    fmt.Println("hello")
}()
wg.Wait()
```

# q
```go
func main() {
    for i := 0; i < 10; i++ {
        go func() {
            fmt.Println(i)
        }()
    }
    time.Sleep(time.Second)
}
// 输出什么？
```

# a
大概率输出一些重复的 10，不是 0-9。

经典的闭包捕获循环变量问题（Go 1.22 之前）。所有 goroutine 共享同一个 i，创建 goroutine 的速度远慢于循环递增，所以执行时 i 可能已经是 10 了。

```go
// Go 1.21 及以前：传参
for i := 0; i < 10; i++ {
    go func(n int) { fmt.Println(n) }(i)
}

// Go 1.22+：循环变量自动隔离
for i := 0; i < 10; i++ {
    go func() { fmt.Println(i) }() // i 每次迭代都是新变量
}
```

# q
goroutine 什么时候会被调度（发生切换）？

# a
goroutine 在以下时机让出 CPU：

1. **主动让出**：channel 阻塞、time.Sleep、runtime.Gosched()
2. **函数调用**：编译器在函数入口插入检查点（morestack），可能触发抢占
3. **系统调用**：进入系统调用时 M 可能阻塞，P 会找其他 M 继续运行其他 G
4. **异步抢占**（Go 1.14+）：基于信号的抢占，长时间运行的循环也会被中断

Go 1.14 之前的坑：`for {}` 空循环永远不让出 CPU，导致其他 G 饿死。1.14 以后不会了。

# q
GOMAXPROCS 是什么？设多少合适？

# a
GOMAXPROCS 决定了同时运行 goroutine 的 P 数量（即最多同时使用多少个 CPU 核）。

默认 = CPU 核数。一般不需要改。

调低：限制 Go 程序的 CPU 使用。调高：基本没用，因为 P 已经等于核数了。

```go
runtime.GOMAXPROCS(4) // 最多同时用 4 个核
```

IO 密集型可以适度调高，CPU 密集型保持默认。

# q
一个 goroutine panic 了，其他 goroutine 会怎样？

# a
如果没有 recover，整个程序 crash，所有 goroutine 都会被终止。

panic 只在自己的 goroutine 中传播，不会传染其他 goroutine。但一个 goroutine 没 recover 的 panic 会导致整个进程退出。

```go
go func() {
    defer func() {
        if r := recover(); r != nil {
            // 只捕获这个 goroutine 的 panic
        }
    }()
    panic("A")
}()

go func() {
    panic("B") // 这个 panic 没人捕获，整个程序 crash
}()
```

# q
goroutine 泄漏是什么？怎么排查？

# a
goroutine 一直阻塞（等 channel、等锁、等 IO），永远不会退出，但又没人引用它 → 泄漏。GC 回收不了 goroutine，它会一直占用内存和调度资源。

常见泄漏场景：
- channel 永远没人发/收
- select 所有 case 都阻塞，没有 default 或超时
- 网络请求没有超时

排查：`pprof.Lookup("goroutine").WriteTo(os.Stdout, 1)` 看 goroutine 堆栈，或者用 `go tool pprof http://localhost:6060/debug/pprof/goroutine`。

# q
runtime.Gosched() 干什么用的？

# a
主动让出当前 goroutine 的执行权，把它放回队列末尾，让其他 goroutine 先跑。

```go
go func() {
    for { fmt.Println("A"); runtime.Gosched() }
}()
go func() {
    for { fmt.Println("B"); runtime.Gosched() }
}()
// A B A B A B ... 交替输出
```

一般不需要用，但在忙等循环或自旋锁实现中可能用到。

# q
goroutine 的栈是固定的吗？

# a
不是，goroutine 栈是**动态伸缩**的。初始只有 2KB，按需增长，最大可达 1GB。

C 和 C++ 线程栈通常是固定的（2MB 或 8MB），大量线程会消耗很多内存。Go 的轻量栈是它支持百万 goroutine 的基础。

栈扩容通过栈拷贝（Go 1.4+）：分配更大的栈，把旧栈内容拷贝过去，更新指针。
