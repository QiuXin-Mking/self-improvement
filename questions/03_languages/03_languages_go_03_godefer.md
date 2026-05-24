# q
Go语言中`defer`关键字的作用和执行时机是什么？
# a
`defer`用于延迟函数或方法的执行，使其在包含`defer`语句的函数返回之前执行，无论函数是正常返回还是因错误返回。常用于在函数退出时进行清理工作。

# q
`defer`有哪些常见用途？
# a
1. **资源释放**：例如关闭文件、关闭数据库连接、解锁互斥体等，确保资源在函数退出前被释放。
2. **跟踪代码执行**：在函数进入或退出时记录日志，便于调试和追踪执行流程。
3. **处理恐慌（Panic）**：配合`recover`函数在`defer`中恢复恐慌，防止程序崩溃。

# q
如何在Go中使用`defer`配合`recover`来捕获并处理恐慌？
# a
在`defer`调用的匿名函数中调用`recover()`，如果程序发生了`panic`，`recover`会捕获到恐慌的值并返回，否则返回`nil`。示例：
```go
defer func() {
    if r := recover(); r != nil {
        fmt.Println("Recovered from panic:", r)
    }
}()
```

