# q
Go 1.18 泛型的基本语法？

# a
```go
// 泛型函数
func Max[T int | float64](a, b T) T {
    if a > b {
        return a
    }
    return b
}

// 调用（类型推断）
m := Max(1, 2)           // T = int
n := Max(1.5, 2.5)       // T = float64
// 显式指定
p := Max[int](1, 2)
```

关键语法：
- `[T constraint]`：类型参数列表，用方括号
- `int | float64`：类型约束，用 | 表示"或"
- 调用时类型可以推断，也可以显式指定

# q
泛型的类型约束（constraint）有哪些？

# a
```go
// 1. any → 任意类型（interface{} 的别名）
func Print[T any](v T) { ... }

// 2. comparable → 可用 ==/!= 比较的类型
func Contains[T comparable](s []T, v T) bool { ... }

// 3. 自定义约束（接口）
type Number interface {
    int | int64 | float64
}
func Sum[T Number](s []T) T { ... }

// 4. 组合约束
type Stringer interface {
    ~string | fmt.Stringer  // ~string 包含底层类型是 string 的自定义类型
}
```

Go 1.18+ 内置 constraints 包（1.21 移到标准库）：constraints.Ordered、constraints.Integer 等。

# q
泛型 struct 怎么写？

# a
```go
type Stack[T any] struct {
    items []T
}

func (s *Stack[T]) Push(v T) {
    s.items = append(s.items, v)
}

func (s *Stack[T]) Pop() (T, bool) {
    if len(s.items) == 0 {
        var zero T  // 零值
        return zero, false
    }
    v := s.items[len(s.items)-1]
    s.items = s.items[:len(s.items)-1]
    return v, true
}
```

注意：方法定义时也需要写 `[T any]`，但方法的类型参数来自类型定义，不能另加。

# q
Go 泛型有什么局限性？

# a
1. **不能做类型断言/switch**：`switch v.(type)` 对泛型类型参数无效
2. **不支持方法上的泛型参数**：方法不能用和类型不同的类型参数
3. **不能访问字段**：即使约束了 struct，也不能 `v.Field`
4. **没有操作符重载**：不能 `a + b`（除非 a、b 是受约束的 Ordered 类型）
5. **编译后代码膨胀**：类似 C++ 模板，不同类型生成不同代码
6. **没有特化**：不能对特定类型写特殊实现

Go 泛型是偏保守的设计，不像 C++ 模板那样图灵完备。

# q
`~` 符号在泛型约束中代表什么？

# a
`~T` 表示底层类型是 T 的所有类型（包括 T 本身和以 T 为底层类型的自定义类型）。

```go
type MyInt int

// 没有 ~：只接受 int
func Add1[T int](a, b T) T { return a + b }

// 有 ~：接受 int 和 MyInt 等底层是 int 的类型
func Add2[T ~int](a, b T) T { return a + b }

Add1(1, 2)       // OK
Add1(MyInt(1), MyInt(2)) // 编译错误
Add2(1, 2)       // OK
Add2(MyInt(1), MyInt(2)) // OK
```

`~` 让泛型约束更实用（能接受基于基础类型的自定义类型）。

# q
泛型 slice/map 操作怎么实现？

# a
```go
// map keys
func Keys[K comparable, V any](m map[K]V) []K {
    keys := make([]K, 0, len(m))
    for k := range m {
        keys = append(keys, k)
    }
    return keys
}

// filter
func Filter[T any](s []T, fn func(T) bool) []T {
    var result []T
    for _, v := range s {
        if fn(v) {
            result = append(result, v)
        }
    }
    return result
}
```

这些是泛型最常见的用途——以前每个类型都要写一遍的工具函数。

# q
什么时候应该用泛型？什么时候不应该？

# a
**应该用**：
- 容器类型（Stack、Set、Queue）
- slice/map 工具函数（Filter、Map、Contains）
- 多种数字类型共用算法

**不应该用**：
- 只有一两种类型 → 写两份代码更清晰
- 只是为了"好看" → 泛型增加了复杂度
- 接口能解决的 → 先考虑接口（但两者解决不同问题）
- 过度抽象 → 简单重复好过错误抽象

Go 社区的共识：泛型是工具，不是装饰。保持简单优先。
