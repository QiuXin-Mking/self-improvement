# q
var m map[string]int 声明后能直接 m["a"]=1 吗？

# a
不行，会 panic: assignment to entry in nil map。

nil map 可以读（返回零值），但不能写。必须用 make 或字面量初始化：

```go
m = make(map[string]int)
m["a"] = 1  // OK
```

nil map 的 len 也是 0。

# q
```go
m := map[string]int{"a": 1}
v := m["b"]       // 不存在，返回什么？
v2, ok := m["b"]  // ok 是什么？
```

# a
v = 0（int 的零值），ok = false。

Go 的 map 访问不存在 key 不会 panic，返回零值。用 comma ok 模式判断 key 是否存在。

# q
map 遍历顺序是怎样的？

# a
顺序是**随机的**，每次运行都可能不同。这是 Go 故意设计的，防止程序依赖遍历顺序。

如果需要有序遍历：先收集 keys → sort → 再按 key 遍历。

```go
keys := make([]string, 0, len(m))
for k := range m {
    keys = append(keys, k)
}
sort.Strings(keys)
for _, k := range keys {
    fmt.Println(k, m[k])
}
```

# q
两个 goroutine 同时写 map 会怎样？

# a
会 panic: concurrent map writes。map 不是并发安全的。

Go 内置的 race detector 能检测到：`go run -race`。

三种解决方案：
1. sync.RWMutex 保护（读多写少场景）
2. sync.Mutex 保护（读写均衡场景）
3. sync.Map（读多写少、key 不固定的场景，内部有缓存优化）

# q
```go
type User struct {
    Name string
}
m := map[string]User{"u1": {Name: "张三"}}
m["u1"].Name = "李四"  // 能编译吗？
```

# a
不能编译。map 中的 value 不可寻址，不能直接修改 value 的字段。

原因：map 扩容时元素地址会变，Go 禁止取 map value 的地址。

解决方案：
```go
// 方案1：整个替换
m["u1"] = User{Name: "李四"}

// 方案2：用指针 value
m2 := map[string]*User{"u1": {Name: "张三"}}
m2["u1"].Name = "李四" // OK
```

# q
往 map 里循环写入会怎样？要注意什么？

# a
在遍历 map 的过程中新增 key，新增的 key 可能出现在本次遍历中，也可能不会——行为是未定义的。

不要在遍历 map 的同时写 map（包括增删），除非你很清楚自己在做什么。

# q
```go
m := map[int]string{1: "a", 2: "b", 3: "c"}
for k := range m {
    if k == 2 {
        delete(m, k)  // 安全吗？
    }
}
```

# a
安全。遍历中删除已经遍历过或正在遍历的 key 是安全的。但删除未遍历的 key，该 key 可能不会再被遍历到。

Go 规范允许遍历中安全删除，这是一种惯用写法（比如清理 map）。

# q
sync.Map 和 map + sync.RWMutex 怎么选？

# a

| | map + RWMutex | sync.Map |
|---|---|---|
| 适用场景 | 读写都有 | 读多写少，key 集合稳定 |
| 原理 | 读写锁保护整个 map | 两个 map + 原子操作 + 延迟拷贝 |
| 类型安全 | 泛型，编译期检查 | interface{}，需类型断言 |
| 性能 | 写多时锁竞争明显 | 读几乎无锁，写少时性能好 |

基本规则：能用 map + RWMutex 就先用它，除非性能 profile 证明有瓶颈。

# q
map 的扩容触发条件和扩容过程是怎样的？

# a
map 存储结构是 bucket 数组，每个 bucket 存 8 个 key-value。

扩容条件（二选一）：
1. 负载因子 > 6.5（平均每个 bucket 超过 6.5 个元素）→ 增量扩容，bucket 数量翻倍
2. 溢出 bucket 太多（溢出 bucket 数 ≥ 普通 bucket 数）→ 等量扩容，重新排列

扩容是**渐进式的**（和 Redis 的一次性 rehash 不同），每次读写操作迁移 1-2 个 bucket，避免一次性迁移的性能抖动。

# q
什么类型能做 map 的 key？

# a
可以用 `==` 比较的类型都能做 key：int、string、指针、channel、interface、struct（所有字段可比较）、数组。

**不能做 key**：slice、map、function（不可比较）。

struct 做 key 时，所有字段都必须可比较：
```go
type Key struct {
    A int
    B string
}
m := map[Key]int{} // OK

type BadKey struct {
    A int
    B []int  // 含 slice，不能做 key
}
// map[BadKey]int{} // 编译错误
```
