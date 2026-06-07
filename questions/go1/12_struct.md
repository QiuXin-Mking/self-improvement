# q
值接收者和指针接收者怎么选？

# a

| | 值接收者 | 指针接收者 |
|---|---|---|
| 能改原值吗 | 不能（操作的是副本） | 能 |
| 调用方 | 值和指针都能调用 | 只有指针能调用 |
| 拷贝开销 | 有（大 struct 开销大） | 无（只拷贝指针） |
| 并发安全 | 天然安全（副本） | 需额外保护 |

选择规则（按优先级）：
1. 方法需要修改接收者 → 指针接收者
2. struct 较大（超过 4 个字段） → 指针接收者，避免拷贝
3. 包含 Mutex 等不能拷贝的字段 → 指针接收者
4. 其他 → 值接收者

一个类型的接收者应该统一：如果一个方法用了指针接收者，其他方法也应该用指针。

# q
```go
type Counter struct {
    n int
}
func (c Counter) Inc() { c.n++ }
func (c *Counter) IncPtr() { c.n++ }

func main() {
    c := Counter{}
    c.Inc()
    fmt.Println(c.n)    // ?
    c.IncPtr()
    fmt.Println(c.n)    // ?
}
```

# a
第一个输出 0，第二个输出 1。

c.Inc() 操作的是 c 的副本，不影响原值。c.IncPtr() 操作的是 c 的指针，修改原值。

# q
结构体嵌入（embedding）和 C++ 继承有什么区别？

# a
Go 的嵌入本质是**组合**，不是继承。没有虚函数、没有多态、没有类型层级。

```go
type Person struct {
    Name string
}
func (p Person) Greet() string {
    return "Hi, I'm " + p.Name
}

type Employee struct {
    Person  // 嵌入
    Company string
}

e := Employee{Person: Person{Name: "张三"}, Company: "ABC"}
fmt.Println(e.Greet()) // 方法提升，可以直接调用
```

虽然用法像继承，但 Employee 不是 Person 的子类型。不能把 Employee 传给接受 Person 的函数。

# q
结构体嵌入中，方法提升的规则是什么？

# a
嵌入类型的导出方法会被"提升"到外层类型，直接通过外层类型调用。

如果外层类型定义了同名方法 → 外层方法覆盖内层（但不是重写，是 shadow）。

```go
func (e Employee) Greet() string {
    return e.Person.Greet() + " from " + e.Company
}
// 调 e.Greet() → Employee 的版本
// 调 e.Person.Greet() → Person 的版本
```

# q
struct tag 是干什么用的？常见有哪些？

# a
struct tag 是字段的元数据，通过反射读取。最常见用途：JSON 序列化、ORM 映射、参数校验。

```go
type User struct {
    ID    int    `json:"id"`           // JSON 字段名为 id
    Name  string `json:"name,omitempty"` // 空值不输出
    Email string `json:"email"`        // 
    pass  string `json:"pass"`         // 不起作用！pass 是非导出字段
}
```

常见 tag：
- `json:"name,omitempty"` — JSON 序列化配置
- `db:"column_name"` — 数据库列名
- `validate:"required,min=1"` — 参数校验
- `yaml:"name"` — YAML 序列化
- `xml:"name,attr"` — XML 序列化

tag 只对导出字段有效（大写开头），内部用反射读取。

# q
```go
type User struct {
    Name string `json:"name"`
}
u := User{Name: "张三"}
data, _ := json.Marshal(u)
// 如果 Name 改成 name（小写），json.Marshal 还能序列化吗？
```

# a
不能。小写字段是非导出的，json.Marshal 通过反射只能看到导出字段。

Go 的可见性是包级别：大写 = 公开，小写 = 私有。reflect 也无法访问非导出字段的值。

# q
```go
// 这两种方式有什么区别？
func (u *User) Update(name string) {
    u.Name = name
}

func (u User) String() string {
    return u.Name
}
```

# a
Update 是指针接收者（能改原值），String 是值接收者（只读）。

指针接收者的方法 set > 值接收者的方法 set：
- *User 能调所有方法（值和指针接收者都算）
- User 只能调值接收者的方法

所以如果 User 只有指针接收者方法，传值会调不到——这也是为什么通常统一用指针接收者。

# q
怎么实现构造函数？Go 有构造函数吗？

# a
Go 没有构造函数语法。惯用 NewXXX 工厂函数：

```go
type Server struct {
    addr string
    port int
}

func NewServer(addr string, port int) *Server {
    s := &Server{
        addr: addr,
        port: port,
    }
    // 验证、初始化...
    return s
}
```

配合 Option 模式（见 17_patterns.md）可以支持可选参数。

# q
空结构体 struct{} 有什么用？

# a
空结构体不占内存（0 字节），用于"信号"场景：

```go
// 当作 set
set := make(map[string]struct{})
set["key"] = struct{}{}
if _, ok := set["key"]; ok {
    // 存在
}

// 当作信号 channel
done := make(chan struct{})
go func() {
    // do work
    close(done)
}()
<-done // 等完成
```

比 bool 省内存，语义更清晰（"我不需要值，只是要个信号"）。

# q
```go
type A struct{ x int }
type B struct{ x int }
a := A{x: 1}
b := B(a) // 能编译吗？
```

# a
如果字段完全一样（类型和名字都相同），能编译。Go 允许同字段结构的结构体之间强制类型转换。

但如果字段有差异（数量不同、类型不同、名字不同、tag 不同），不能转换。只有底层类型相同才可以。

注意：这里的转换是值拷贝，a 和 b 是独立的。
