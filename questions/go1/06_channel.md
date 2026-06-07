# q
channel 有几种类型？分别有什么特点？

# a
三种：无缓冲 channel、有缓冲 channel、nil channel。

| 类型 | 声明 | 发送行为 | 接收行为 |
|------|------|---------|---------|
| 无缓冲 | make(chan int) | 阻塞直到有人接收 | 阻塞直到有人发送 |
| 有缓冲 | make(chan int, 5) | 缓冲满才阻塞 | 缓冲空才阻塞 |
| nil | var ch chan int | 永远阻塞 | 永远阻塞 |

nil channel 在 select 中很有用：把 case 的 channel 设为 nil 可以禁用该 case，避免 CPU 空转。

# q
```go
ch := make(chan int, 2)
ch <- 1
ch <- 2
// 不读，直接 close(ch) 会发生什么？
```

# a
可以正常 close。关闭 channel 后，缓冲中的数据仍然可以读取。读完后继续读返回零值和 false。

关闭规则（重要）：
- 只能发送方关闭（接收方关闭会 panic）
- 不能重复关闭（panic）
- 不能向已关闭 channel 发送（panic）
- 关闭 nil channel（panic）

# q
```go
ch := make(chan int, 3)
ch <- 1; ch <- 2; ch <- 3
close(ch)
for v := range ch {
    fmt.Println(v)
}
// 输出什么？会死循环吗？
```

# a
输出 1 2 3，然后正常退出。range channel 在 channel 关闭后自动结束。

如果没 close，range 会永远阻塞等新数据，导致 goroutine 泄漏。

# q
从已关闭的 channel 读会发生什么？

# a
能读到缓冲中剩余的数据。清空后继续读返回零值 + false。

```go
ch := make(chan int, 2)
ch <- 1
ch <- 2
close(ch)

v1, ok1 := <-ch // 1, true
v2, ok2 := <-ch // 2, true
v3, ok3 := <-ch // 0, false
```

# q
```go
ch := make(chan int)
close(ch)
ch <- 1
// 发生什么？
```

# a
panic: send on closed channel。

向已关闭的 channel 发数据是 Go 的运行时 panic。这个没法 recover 恢复后继续用（channel 已经被关了）。

# q
```go
func main() {
    ch := make(chan int)
    <-ch
}
// 发生什么？
```

# a
死锁。所有 goroutine 都休眠了，没有 goroutine 能向 ch 发数据。

运行时检测到这种情况会 panic: all goroutines are asleep - deadlock!

注意：这种检测只在没有其他活跃 goroutine 时才触发。如果还有别的 goroutine 在跑，就不会报 deadlock。

# q
select 多个 case 就绪时，选哪个？

# a
**随机选择**。伪随机，均匀分布。不能用 select 的随机性来实现公平性——不要依赖 case 的选中顺序。

```go
select {
case <-ch1:
    // ...
case <-ch2:
    // ...
}
// 两者都就绪时，等概率选一个
```

# q
select 里的 default 分支有什么用？

# a
default 让 select 变成非阻塞的——如果所有 case 都不就绪，立刻执行 default，不会阻塞等待。

```go
select {
case msg := <-ch:
    // 有消息
default:
    // 没消息，不阻塞，继续往下走
}
```

经典用法：
1. 非阻塞 channel 操作（try-send / try-receive）
2. 定时任务中的轮询
3. 防止 goroutine 泄漏（避免永久阻塞在 select 上）

# q
```go
select {
case ch <- 1:
    fmt.Println("sent")
default:
    fmt.Println("full")
}
// ch 什么时候输出 sent？什么时候输出 full？
```

# a
有无缓冲 channel：
- 无缓冲 ch → 永远输出 full（没人接收）
- 有缓冲 ch → 缓冲不满输出 sent，满了输出 full

这就是 try-send 模式。

# q
channel 的底层实现是怎样的？

# a
channel 底层是一个共享内存的环形缓冲区 + mutex。

```go
type hchan struct {
    qcount   uint          // 缓冲区中元素个数
    dataqsiz uint          // 缓冲区大小
    buf      unsafe.Pointer // 环形缓冲区指针
    sendx    uint          // 发送索引
    recvx    uint          // 接收索引
    sendq    waitq         // 等待发送的 goroutine 队列
    recvq    waitq         // 等待接收的 goroutine 队列
    lock     mutex         // 互斥锁
}
```

发送过程：
1. 如果有等待接收的 G → 直接把数据拷贝给接收方（不经过缓冲区，快！）
2. 如果缓冲区有空位 → 放入缓冲区
3. 否则 → 当前 G 入 sendq 等待队列，挂起

接收过程对称。所以其实 channel 是用锁实现的——"不要用共享内存通信"是理念，底层还是锁。

# q
什么是 channel 的"happens-before"保证？

# a
channel 操作提供了 happens-before 保证：

1. 发送 happens-before 对应的接收完成
2. 关闭 channel happens-before 从该 channel 读到零值
3. 无缓冲 channel：接收 happens-before 对应的发送完成

这意味着：发送前的修改，接收方一定可见。channel 不仅是通信机制，也是同步机制。

# q
单向 channel 有什么用？怎么用？

# a
单向 channel 用于函数签名，让编译器帮你检查 channel 的读写方向。

```go
func producer(out chan<- int) {  // 只能发
    out <- 1
}

func consumer(in <-chan int) {   // 只能收
    v := <-in
}

func main() {
    ch := make(chan int)         // 双向
    go producer(ch)
    consumer(ch)
}
```

双向 channel 可以隐式转换为单向，但反过来不行。
