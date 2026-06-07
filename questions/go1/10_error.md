# q
Go 的 error 和 C++ 的 exception 核心区别是什么？

# a
| | Go error | C++ exception |
|---|---|---|
| 机制 | 返回值 | throw/catch |
| 开销 | 极低 | 有栈展开开销 |
| 显式性 | 必须显式处理 | 可能被遗漏 |
| 控制流 | 正常返回 | 跳转到 catch |
| 建议 | 可预期的错误 | 严重异常 |

Go 哲学：错误就是值，不是"异常"。不要用 panic 模拟 exception。

# q
```go
f, err := os.Open("file.txt")
if err != nil {
    // 怎么判断是什么错误？
}
```

# a
```go
// 1. sentinel error 比较
if errors.Is(err, os.ErrNotExist) {
    // 文件不存在
}

// 2. 类型断言（检查错误类型）
var pathError *os.PathError
if errors.As(err, &pathError) {
    // 是路径相关的错误
}

// 3. 简单字符串比较（不推荐）
if strings.Contains(err.Error(), "not found") {
    // 脆弱，错误信息可能变化
}
```

`errors.Is` 和 `errors.As` 是 Go 1.13+ 引入的，能正确处理错误链（wrapped error）。

# q
`errors.Is` 和 `errors.As` 有什么区别？

# a
`errors.Is`: 判断错误链中是否包含某个特定错误（值比较）。

`errors.As`: 判断错误链中是否有某类型的错误，有则提取出来（类型匹配）。

```go
// errors.Is: 值比较
if errors.Is(err, io.EOF) { ... }

// errors.As: 类型匹配
var netErr *net.OpError
if errors.As(err, &netErr) {
    fmt.Println(netErr.Op) // 提取具体字段
}
```

# q
怎么包装一个错误并添加更多上下文？

# a
```go
// Go 1.13+: fmt.Errorf + %w
err := fmt.Errorf("read config: %w", ioErr)

// 或第三方库
err := errors.Wrap(ioErr, "read config")

// 检测
errors.Is(err, ioErr)   // true，能追溯到根因
errors.As(err, &target) // true
```

关键：%w（不是 %v）才能让 errors.Is/As 穿透包装层。

# q
panic 和 error 怎么选择？

# a
用 error 的场景（大多数情况）：
- 网络超时、文件不存在、格式错误
- 调用方可以处理的错误

用 panic 的场景（极少数）：
- 程序 bug（数组越界、nil pointer dereference）
- 不可恢复的状态（配置文件缺失且无默认值）
- init() 函数中的初始化失败

简单规则：**库函数永远返回 error，不 panic**。只有顶层 main 或 init 可以考虑 panic。

# q
```go
if err != nil {
    return err
}
```
为什么 Go 代码里到处都是这个？不烦人吗？

# a
确实啰嗦，但 Go 社区认为这是优点：

1. 错误处理就在代码路径上，一看就知道错误去哪了
2. 没有隐藏的控制流（不像 C++ exception 会跳到哪里不知道）
3. 强迫程序员思考每个错误怎么处理

现在 Go 2 的 error handling 提案也在讨论简化方案，但 Go 1.x 就是这个风格。接受了就好了。

# q
自定义 error 类型怎么写？

# a
```go
// 简单版
var ErrNotFound = errors.New("not found")

// 带信息版
type ValidationError struct {
    Field string
    Value interface{}
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("invalid %s: %v", e.Field, e.Value)
}

// 使用
func validate() error {
    return &ValidationError{Field: "age", Value: -1}
}
```

sentinel error 用 errors.New / fmt.Errorf，复杂错误自定义 struct 实现 Error() 接口。

# q
defer + recover 怎么捕获 panic？

# a
```go
func safeCall() (err error) {
    defer func() {
        if r := recover(); r != nil {
            err = fmt.Errorf("panic: %v", r)
        }
    }()
    // 可能 panic 的代码
    doSomething()
    return nil
}
```

注意：recover 必须**直接**在 defer 函数体中调用。嵌套调用无效。

# q
什么时候在 defer 里 recover 还不够？

# a
1. 跨 goroutine 的 panic：每个 goroutine 必须有自己的 recover
2. 不可恢复的崩溃：如内存耗尽、栈溢出
3. 并发不安全导致的 race panic：说明代码逻辑有 bug，应该修复而非 recover
4. recover 不能阻止程序退出——如果 panic 是因为所有 goroutine 都休眠了（deadlock），recover 也捕获不了

recover 是最后的手段，不是 error handling 的替代品。

# q
```go
err := doSomething()
// 如果不检查 err，Go 编译器报错吗？
```

# a
不报错。Go 不会强制检查 error 返回值。

但可以用 linter 检查：
- `errcheck` 工具：检测未检查的 error
- golangci-lint 中的 `errcheck` 规则
- IDE 也能提示（Go 插件会标黄）

工程化建议：ci 里跑 golangci-lint 并开启 errcheck。
