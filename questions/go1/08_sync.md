# q
sync.Mutex 和 sync.RWMutex 的区别？什么时候用哪个？

# a

| | sync.Mutex | sync.RWMutex |
|---|---|---|
| 锁粒度 | 读写互斥 | 读不互斥，写互斥 |
| 方法 | Lock/Unlock | RLock/RUnlock + Lock/Unlock |
| 开销 | 较低 | 稍高（多了读者计数） |
| 适合 | 写多或读写均衡 | 读多写少 |

RWMutex 允许多个 reader 同时持有读锁，但写锁排斥所有读写。

# q
Mutex 可以重复 Lock 吗？

# a
不可以。同一个 goroutine 重复 Lock 同一个 Mutex → panic: sync: fatal lock。

Go 的 Mutex 不是可重入锁。需要可重入的场景自己实现或用其他方式设计（比如把锁范围缩小）。

```go
mu.Lock()
// ...
mu.Lock() // panic!
```

# q
sync.WaitGroup 的基本用法？忘了 Done 会怎样？忘了 Add 会怎样？

# a
```go
var wg sync.WaitGroup
for i := 0; i < 5; i++ {
    wg.Add(1)
    go func() {
        defer wg.Done()
        // do work
    }()
}
wg.Wait() // 等所有 goroutine 完成
```

- 忘了 Done → Wait 永远阻塞，goroutine 泄漏
- 忘了 Add（只跑 goroutine 没 Add）→ Wait 可能提前返回，不等 goroutine
- Add 在 goroutine 内部而非外部 → 有 race risk（Add 没执行就被 Wait 了）
- 重复 Done（减到负数）→ panic: negative WaitGroup counter

# q
sync.Once 干什么用的？怎么实现单例最安全？

# a
sync.Once 保证函数只执行一次，即使多个 goroutine 同时调用。

```go
var (
    once     sync.Once
    instance *Singleton
)

func GetInstance() *Singleton {
    once.Do(func() {
        instance = &Singleton{}
    })
    return instance
}
```

内部实现用原子操作 + 互斥锁双重检查。比单纯用 Mutex 高效（只有第一次用锁）。

# q
sync.Cond 是干什么用的？什么场景需要它？

# a
Cond 是条件变量，让 goroutine 等待某个条件成立后再继续执行。

```go
var mu sync.Mutex
cond := sync.NewCond(&mu)
ready := false

// 等待方
go func() {
    mu.Lock()
    for !ready {       // 循环检查条件！
        cond.Wait()    // Wait 会释放锁，被广播后重新获取锁
    }
    // 条件成立，继续
    mu.Unlock()
}()

// 通知方
mu.Lock()
ready = true
mu.Unlock()
cond.Broadcast() // 或者 cond.Signal()（唤醒一个）
```

关键点：Wait() 必须在 Lock 后调用，Wait 内部会 unlock 再等待，被唤醒后重新 lock。

# q
sync.Map 的 LoadOrStore 是什么？

# a
如果 key 存在就返回已有 value，不存在就存储新 value 并返回。

```go
v, loaded := m.LoadOrStore("key", "value")
// loaded=true: v 是已有值
// loaded=false: v 是刚存储的 value
```

类似 Redis 的 SETNX + GET 的原子版本。比手动 Load → 检查 → Store 安全（没有竞态条件）。

# q
atomic 包的原子操作和 Mutex 锁有什么区别？

# a
| | atomic | Mutex |
|---|---|---|
| 粒度 | 单次操作（加、赋值、CAS） | 一整段代码 |
| 开销 | 极低（CPU 指令级） | 有锁竞争开销 |
| 复杂度 | 低 | 高 |
| 适用 | 简单计数器、标志位 | 复杂数据结构保护 |

atomic 只适合简单的"读-改-写"操作。需要保护多行代码的场景必须用 Mutex。

```go
var count int64
atomic.AddInt64(&count, 1) // 原子递增
v := atomic.LoadInt64(&count) // 原子读取
```

# q
```go
var mu sync.Mutex
var done bool
mu.Lock()
done = true
mu.Unlock()
// done 对其他 goroutine 可见吗？
```

# a
可见。Mutex 的 Unlock 和后续的 Lock 之间有 happens-before 关系。

Unlock 之前的所有修改，在后续 Lock 之后对所有 goroutine 可见。

但这里有个问题：如果不通过 Lock/Unlock 来读 done（直接读），即使 done 被 Lock 保护写入，读操作仍然可能看到旧值（data race）。要么读也用 Lock 保护，要么用 atomic。

# q
```go
type Cache struct {
    mu sync.RWMutex
    m  map[string]string
}
// Get 和 Set 分别怎么写？
```

# a
```go
func (c *Cache) Get(key string) (string, bool) {
    c.mu.RLock()
    v, ok := c.m[key]
    c.mu.RUnlock()
    return v, ok
}

func (c *Cache) Set(key, value string) {
    c.mu.Lock()
    c.m[key] = value
    c.mu.Unlock()
}
```

读用 RLock（允许多个 reader 并发），写用 Lock（互斥）。注意不要把 RUnlock 写成 Unlock（反之亦然）。

# q
怎么排查代码中的数据竞争？

# a
用 Go 的 race detector：

```bash
go test -race ./...
go run -race main.go
go build -race -o myapp
```

原理：编译期插桩，运行时记录每次内存访问的 goroutine ID 和时间戳，检测到冲突就打印详细堆栈。

代价：CPU 慢 5-10 倍，内存多 5-10 倍。仅开发/测试用，不要上生产。
