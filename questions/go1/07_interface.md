# q
Go 的接口和 C++ 的虚函数有什么区别？

# a
本质区别：Go 是**隐式实现**（鸭子类型），C++ 是**显式继承**。

| | Go interface | C++ 虚函数 |
|---|---|---|
| 实现方式 | 隐式：方法匹配即实现 | 显式：必须继承/声明 |
| 耦合 | 低：实现者不依赖接口定义 | 高：实现者必须知道接口 |
| 定义方式 | 接口和实现分离 | 类声明中指定继承 |
| 类型安全 | 编译期检查方法签名 | 编译期检查继承关系 |
| 依赖方向 | 接口→实现 | 实现→接口 |

Go 的设计哲学：由使用者定义接口，而不是实现者。比如 io.Reader 只关心"能 Read 就行"，不关心你是谁。

# q
```go
var f *os.File = nil
var w io.Writer = f
// w == nil 吗？
```

# a
w != nil。这是 Go 面试最高频的陷阱。

interface 底层结构 = (type, data) 两部分。只有两个都是 nil，接口才 == nil。

这里：
- type = *os.File（不是 nil！）
- data = nil（f 是 nil 指针）

所以 w 的 type 非 nil，w != nil。但调用 w.Write() 会 panic（data 是 nil）。

```go
func returnsError() error {
    var p *MyError = nil
    return p  // 返回的 error != nil！！！
}
// 正确写法
func returnsError() error {
    if somethingWrong {
        return &MyError{}
    }
    return nil  // 真 nil
}
```

# q
类型断言 v.(type) 和 v, ok := v.(type) 的区别？

# a
不带 ok 的断言失败会 panic，带 ok 的断言失败返回零值和 false。

```go
var x interface{} = 42
s := x.(string)       // panic: interface conversion
s, ok := x.(string)   // ok=false, s=""，不 panic
```

类型断言只能用于 interface 类型。具体类型不能断言。

# q
type switch 怎么用？和类型断言有什么区别？

# a
type switch 可以一次判断多种类型：

```go
switch v := x.(type) {
case int:
    // v 是 int
case string:
    // v 是 string
case nil:
    // 空接口
default:
    // 未知类型
}
```

优势：比连续的类型断言清晰，而且可以同时匹配多个类型：`case int, float64:`。

# q
```go
type Animal interface {
    Speak() string
}
type Dog struct{}
func (d Dog) Speak() string { return "woof" }

var a Animal = Dog{}
var b Animal = &Dog{}
// a 和 b 都能赋值吗？
```

# a
都可以。

Dog 实现了 Speak()（值接收者），那 *Dog 和 Dog 都算实现了接口。

反过来：如果只有 *Dog 实现了 Speak()（指针接收者），那 Dog 不能赋值给接口。

关键规则：
- 值接收者方法 → 值和指针都能实现接口
- 指针接收者方法 → 只有指针能实现接口

# q
```go
type Writer interface {
    Write([]byte) (int, error)
}

// 这段代码中，哪些类型实现了 Writer？
// 1. *os.File
// 2. os.File  
// 3. *bytes.Buffer
```

# a
1 和 3，os.File 本身不实现 Writer。

`*os.File` 有 Write 方法（指针接收者），所以指针实现了 Writer。`os.File` 值类型没有继承指针接收者的方法，所以不实现。

`*bytes.Buffer` 有 Write 方法（指针接收者），指针实现了 Writer。

实际上 os.File 结构体的方法几乎都是指针接收者（涉及修改和锁），所以一般都是传指针。

# q
空接口 interface{}（any）和 void* 的区别？

# a
| | interface{} / any | void* (C/C++) |
|---|---|---|
| 类型安全 | 有类型信息，运行时检查 | 无类型信息 |
| 解引用 | 类型断言后使用 | 强制转换 |
| 转换失败 | 断言可以返回 ok | 未定义行为 |
| GC | 参与 GC | 不管 GC |

Go 的 interface{} 携带类型信息 + 数据指针，类型断言是安全的。C 的 void* 只是裸指针，转换错误是 UB。

# q
接口的底层结构是怎样的？

# a
```go
// 有方法的接口（iface）
type iface struct {
    tab  *itab  // 类型信息 + 方法表
    data unsafe.Pointer
}

// 空接口（eface）
type eface struct {
    _type *_type  // 只存类型信息
    data  unsafe.Pointer
}
```

itab 中缓存了接口方法集和具体类型方法集的映射，第一次赋值时生成并缓存（hash 查找），后续直接复用。

# q
用接口的好处有哪些？在 Go 中什么时候该用接口？

# a
Go 中接口的优势：
1. **解耦**：调用方定义接口，不依赖实现方
2. **可测试**：轻松 mock 替换
3. **组合**：小接口组合成大接口（io.Reader + io.Writer = io.ReadWriter）

什么时候用接口：
- 作为函数参数，让调用更灵活
- 多态：多种实现统一处理
- mock/stub 测试

什么时候不用：
- 只有一种实现 → 不需要抽象
- 过度设计 → Go 接口应该小而精（通常 1-3 个方法）
- 接口定义在实现方 → 违反 Go 惯例

Go 哲学：**接口应该在使用方定义，而不是实现方**。接受接口，返回结构体。

# q
```go
type Foo struct {
    A int
}
type Bar struct {
    A int
}
var x interface{} = Foo{A: 1}
y, ok := x.(Bar) // 能断言成功吗？
```

# a
不能。ok = false。

即使 Foo 和 Bar 字段完全一样，也是不同类型。Go 的类型系统是**名义类型**（除了 channel、slice、map 等复合类型通过元素类型判断）。

结构体之间即使字段相同也不能互转。需要的话自己写转换函数。
