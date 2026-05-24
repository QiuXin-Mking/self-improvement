# q
go 怎么看coredump
# a
dlv core /opt/macrosan/mdbs/prometheus/bin/mdbs_exporter core_mdbs_exporter_721078

# q
go defer怎么使用
# a
defer是 Go 语言里用来延迟执行某个函数调用的关键字，常用于资源释放、解锁、关闭连接等场景，确保即使发生 panic 也会被执行。

# q
讲出打印
```go
package main

import "fmt"

func main() {
    fmt.Println("start")
    defer fmt.Println("defer 1")
    defer fmt.Println("defer 2")
    fmt.Println("end")
}
```
# a
start
end
defer 2
defer 1
defer语句会立即求值（参数先算好），但函数体延后到当前函数返回前才执行。
多个 defer按 后进先出（LIFO）​ 的顺序执行。

# q
回忆 defer 互斥锁解锁
# a
```go
var mu sync.Mutex

func update() {
    mu.Lock()
    defer mu.Unlock() // 保证锁一定被释放
    // 修改共享数据
}
```

