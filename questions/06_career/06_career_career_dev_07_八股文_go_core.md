# q
如何使用 dlv 工具分析 Go 程序的 coredump 文件？
# a
```bash
dlv core /opt/macrosan/mdbs/prometheus/bin/mdbs_exporter core_mdbs_exporter_721078
```

# q
defer 关键字的作用是什么？多个 defer 的执行顺序是怎样的？
# a
defer 用于延迟执行函数调用，常用于资源释放、解锁、关闭连接等场景，即使发生 panic 也会执行。defer 语句会立即对参数进行求值，但函数体延迟到外层函数返回前才执行。多个 defer 按后进先出（LIFO）的顺序执行。

# q
请给出 defer 在互斥锁解锁中的典型用法。
# a
```go
var mu sync.Mutex

func update() {
    mu.Lock()
    defer mu.Unlock() // 保证锁一定被释放
    // 修改共享数据
}
```

