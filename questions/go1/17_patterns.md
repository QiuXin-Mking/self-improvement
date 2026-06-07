# q
Go 的标准项目结构是怎样的？

# a
Go 没有官方强制的项目结构，但社区演进出一些常见模式：

```
project/
├── cmd/            # 可执行入口（每个子目录一个 main）
│   ├── server/
│   └── worker/
├── internal/       # 私有包（外部项目不能 import）
│   ├── handler/
│   ├── service/
│   └── repository/
├── pkg/            # 可公开的库包（其他项目可 import）
├── api/            # API 定义（proto、OpenAPI）
├── configs/        # 配置文件
├── migrations/     # 数据库迁移
├── scripts/        # 构建/部署脚本
├── go.mod
└── go.sum
```

但小项目不需要这么复杂。一个 main.go + 几个包就行了。别过度设计。

# q
Option 模式（Functional Options）是什么？

# a
用于构造复杂对象，每个 option 是独立的可选参数。

```go
type Server struct {
    addr    string
    timeout time.Duration
    maxConn int
}

type Option func(*Server)

func WithTimeout(d time.Duration) Option {
    return func(s *Server) { s.timeout = d }
}

func WithMaxConn(n int) Option {
    return func(s *Server) { s.maxConn = n }
}

func NewServer(addr string, opts ...Option) *Server {
    s := &Server{
        addr:    addr,
        timeout: 30 * time.Second,  // 默认值
    }
    for _, opt := range opts {
        opt(s)
    }
    return s
}

// 使用
s := NewServer(":8080", WithTimeout(10*time.Second), WithMaxConn(1000))
```

好处：向后兼容（加新 option 不影响旧调用）、自文档化、零值友好。

# q
Worker Pool 模式怎么实现？

# a
```go
func workerPool(numWorkers int, jobs <-chan Job, results chan<- Result) {
    var wg sync.WaitGroup
    for i := 0; i < numWorkers; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for job := range jobs { // channel 关闭时自动退出
                results <- process(job)
            }
        }()
    }

    // 等所有 worker 完成
    go func() {
        wg.Wait()
        close(results)
    }()
}
```

关键点：
1. 用 channel 分发 job（jobs chan）和收集结果（results chan）
2. range channel 在 channel 关闭时自动结束
3. 发送方关闭 channel 通知 worker 结束

# q
Pipeline 模式怎么实现？

# a
```go
// 每个阶段接收 channel，返回 channel
func gen(nums ...int) <-chan int {
    out := make(chan int)
    go func() {
        for _, n := range nums {
            out <- n
        }
        close(out)
    }()
    return out
}

func sq(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        for n := range in {
            out <- n * n
        }
        close(out)
    }()
    return out
}

// 组合
out := sq(sq(gen(1, 2, 3))) // 1 → 1 → 1; 2 → 4 → 16; 3 → 9 → 81
```

每个阶段独立的 goroutine 负责处理，通过 channel 连接。类似 Unix 管道。

# q
Fan-out / Fan-in 模式怎么实现？

# a
Fan-out：一个输入 channel 分发给多个 worker（通过启动多个 goroutine 读同一个 channel）。

Fan-in：多个 channel 的输出合并到一个 channel。

```go
// Fan-in
func merge(chs ...<-chan int) <-chan int {
    var wg sync.WaitGroup
    out := make(chan int)
    for _, ch := range chs {
        wg.Add(1)
        go func(c <-chan int) {
            defer wg.Done()
            for v := range c {
                out <- v
            }
        }(ch)
    }
    go func() {
        wg.Wait()
        close(out)
    }()
    return out
}
```

Fan-out/Fan-in 常用于并行处理 CPU 密集型任务或并发 IO 操作。

# q
Graceful Shutdown 的标准写法？

# a
```go
func main() {
    // 启动服务
    srv := startServer()
    
    // 监听信号
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    
    select {
    case <-quit:
        log.Println("Shutting down...")
    case <-ctx.Done():
        // 其他退出条件
    }
    
    // 优雅关闭
    shutdownCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()
    
    if err := srv.Shutdown(shutdownCtx); err != nil {
        log.Fatal("Shutdown error:", err)
    }
}
```

关闭顺序通常是：先停接收新请求 → 等现有请求处理完 → 关闭数据库连接等资源。

# q
让各个 goroutine 之间通信，channel 和共享变量 + Mutex 怎么选？

# a
Go 名言："Don't communicate by sharing memory; share memory by communicating."

| 场景 | 选 channel | 选 Mutex |
|------|-----------|---------|
| 传递数据（生产者→消费者） | ✓ | ✗（需要手动管理） |
| 保护共享状态 | ✗（可以但别扭） | ✓ |
| 等待/通知 | ✓ | ✗（要用 Cond） |
| 简单计数器/标志 | ✗（atomic 更好） | ✓ |

实际中经常混用：channel 传任务给 worker，worker 内部用 Mutex 保护本地状态。

# q
怎么组织 package 才合理？按功能拆还是按层级拆？

# a
Go 推荐**按功能拆分**（domain-driven），而不是按层级（MVC）。

```go
// 按层级（不推荐）
project/
├── handlers/
├── services/
├── repositories/
└── models/

// 按功能（推荐）
project/
├── user/      // 用户相关都在这里
│   ├── handler.go
│   ├── service.go
│   └── repo.go
├── order/
│   ├── handler.go
│   └── service.go
└── payment/
```

按功能拆的好处：
- 一个功能的所有代码在一起，修改不跨包
- 明确的所有权和边界
- 新人好理解（按业务概念导航）

# q
什么时候用 sync.Once？什么场景必须用它？

# a
```
sync.Once 保证函数只执行一次，即使并发调用。

必须用 Once 的场景：
1. 单例初始化（第一次用到时才初始化）
2. 数据库连接池、配置加载（只初始化一次）
3. 资源的懒加载 + 并发安全

```go
var (
    db   *sql.DB
    once sync.Once
)

func GetDB() *sql.DB {
    once.Do(func() {
        db, _ = sql.Open("postgres", connStr)
    })
    return db
}
```

如果不用 Once 用 Mutex，需要双重检查锁（DCL），容易写错。Once 就是正确的 DCL 实现。
