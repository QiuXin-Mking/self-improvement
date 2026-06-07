# q
context.Context 是干什么的？

# a
context 主要做三件事：
1. **超时控制**：设置请求的 deadline，超时自动取消
2. **取消传播**：父 context 取消，所有子 context 自动取消
3. **值传递**：在调用链中传递 request-scoped 数据（traceID、user 等）

四个创建方法：
- context.Background() — 根 context
- context.WithCancel() — 可取消
- context.WithTimeout() — 超时取消
- context.WithValue() — 携带值

# q
```go
ctx, cancel := context.WithTimeout(context.Background(), time.Second)
cancel()
// 谁负责调 cancel？
```

# a
调用方负责调 cancel()。即使 WithTimeout 会在一秒后自动取消，也应该在函数结束时手动 cancel，释放 timer 资源。

标准用法：
```go
ctx, cancel := context.WithTimeout(ctx, time.Second)
defer cancel()
// ...
```

手动 cancel 能让子 goroutine 提前退出，不用等 timeout。用 defer 保证一定调用。

# q
context 应该在函数签名中放在什么位置？

# a
永远是第一个参数，命名 ctx。

```go
func Process(ctx context.Context, data string) error {
    // ...
}
```

不要放在 struct 字段中（除非是 http.Request 这种标准库的已有设计）。context 应该像水流一样在调用链中传递。

# q
context 中的 value 怎么存、怎么取？

# a
```go
// 存（创建新 context）
type keyType string
const TraceIDKey keyType = "traceID"
ctx = context.WithValue(ctx, TraceIDKey, "abc123")

// 取
traceID, ok := ctx.Value(TraceIDKey).(string)
if !ok {
    // traceID 不存在或类型不对
}
```

注意点：
1. key 应该是自定义类型（不是 string），避免不同包 key 冲突
2. value 应该只放 request-scoped 数据（traceID、用户信息），不要放业务参数
3. 不要用 context 传过多数据，会变成"大杂烩"

# q
HTTP 请求中怎么正确使用 context？

# a
```go
func handler(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context()           // 从 request 获取
    ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
    defer cancel()

    result, err := doWork(ctx)   // 传给下游
    if err != nil {
        if errors.Is(err, context.DeadlineExceeded) {
            http.Error(w, "timeout", http.StatusGatewayTimeout)
            return
        }
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    // ...
}
```

http.Request 自带 context（客户端断开连接时自动取消）。下游操作（数据库查询、RPC 调用）应该检查 ctx.Done()，及时响应取消。

# q
```go
ctx, cancel := context.WithCancel(parentCtx)
cancel()
// 调了 cancel 后，ctx 的 Done() channel 会怎样？
```

# a
Done() channel 会被关闭。关闭的 channel 可以无限读取，不会阻塞。

```go
select {
case <-ctx.Done():
    // 收到取消信号
    return ctx.Err() // context.Canceled
}
```

# q
如何实现一个 goroutine 的超时控制？

# a
```go
func doWithTimeout(ctx context.Context, fn func() error) error {
    errCh := make(chan error, 1)
    go func() {
        errCh <- fn()
    }()

    select {
    case err := <-errCh:
        return err
    case <-ctx.Done():
        return ctx.Err()
    }
}
```

注意 errCh 必须是缓冲 channel（cap=1），否则 goroutine 可能因没人接收而永远阻塞→泄漏。

# q
context 取消后，正在运行的 goroutine 会自动停止吗？

# a
不会。context 只是一个信号，goroutine 需要主动检查 ctx.Done() 来响应取消。

```go
func worker(ctx context.Context) {
    for {
        select {
        case <-ctx.Done():
            return // 主动退出
        default:
            // 继续工作
        }
    }
}
```

context 是"建议性"的取消，不是强制性的。goroutine 如果不检查 ctx.Done()，就完全不受影响。

# q
多个 context 可以合并吗？比如同时等待两个 context 的取消？

# a
Go 没有内置的 context 合并。但有社区方案或手动实现：

```go
func mergeCtx(ctx1, ctx2 context.Context) context.Context {
    ctx, cancel := context.WithCancel(context.Background())
    go func() {
        select {
        case <-ctx1.Done():
            cancel()
        case <-ctx2.Done():
            cancel()
        case <-ctx.Done():
            // merged ctx 自己被取消了
        }
    }()
    return ctx
}
```

实际使用较少——通常一个 context 就足够。

# q
context 应该在哪些场景使用？

# a
1. HTTP/RPC 请求：透传超时和 traceID
2. 数据库操作：用 context 控制查询超时
3. 并发任务：一个子任务失败，取消其他子任务
4. 长时间运行的任务：支持优雅关闭

不应该用 context 的场景：
- 替代函数参数传数据（应该用参数）
- 存储全局配置（应该用配置结构体）
- 作为 struct 字段（反模式，除了极少例外）
