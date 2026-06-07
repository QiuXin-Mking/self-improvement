# q
json.Marshal 和 json.Unmarshal 的基本用法？

# a
```go
type User struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

// 序列化
u := User{Name: "张三", Age: 30}
data, err := json.Marshal(u)
// data = {"name":"张三","age":30}

// 反序列化
var u2 User
err = json.Unmarshal(data, &u2)
```

Unmarshal 需要传指针。如果不传指针，只会填充副本，毫无效果。

# q
omitempty 是什么？有什么坑？

# a
`json:"name,omitempty"` — 零值时字段不出现在 JSON 中。

```go
type User struct {
    Name  string `json:"name,omitempty"`
    Age   int    `json:"age,omitempty"`
    Admin bool   `json:"admin,omitempty"`
}

u := User{Name: "", Age: 0, Admin: false}
json.Marshal(u) // → {} 全都没有！
```

坑：bool false、int 0、空字符串都算 empty。如果你的业务需要区分"没设置"和"值为 0"，就不能用 omitempty。解决：用指针类型 `*bool` — nil 才 omitempty，false 会输出。

# q
怎么反序列化未知结构的 JSON（动态字段）？

# a
```go
// 方式1：map
var m map[string]interface{}
json.Unmarshal(data, &m)

// 方式2：json.RawMessage 延迟解析
type Event struct {
    Type string          `json:"type"`
    Data json.RawMessage `json:"data"` // 先不解析
}
// 根据 Event.Type 决定怎么解析 Data
```

# q
怎么自定义 JSON 序列化/反序列化？

# a
实现 json.Marshaler 和 json.Unmarshaler 接口。

```go
type UserID int

func (u UserID) MarshalJSON() ([]byte, error) {
    return json.Marshal(strconv.Itoa(int(u)))
}

func (u *UserID) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    n, _ := strconv.Atoi(s)
    *u = UserID(n)
    return nil
}
```

注意：MarshalJSON 用值接收者就能触发，UnmarshalJSON 必须是指针接收者（要修改）。

# q
json.NewDecoder 和 json.NewEncoder 什么时候用？

# a
流式处理时用 Decoder/Encoder，避免一次性把全部数据读进内存。

```go
// Reader 流式输入（http.Body、文件等）
decoder := json.NewDecoder(r.Body)
var user User
decoder.Decode(&user)

// Writer 流式输出
encoder := json.NewEncoder(w)
encoder.Encode(user)
```

一次性处理小 JSON 用 Marshal/Unmarshal。处理流（HTTP 请求体、日志流等）用 Decoder/Encoder。

# q
只用 json.Unmarshal 解析部分字段怎么处理？

# a
```go
type User struct {
    Name string `json:"name"`
    // Age 不定义
}

// 未知字段会被忽略，不报错
var u User
json.Unmarshal([]byte(`{"name":"张三","age":30,"city":"深圳"}`), &u)
// u = {Name:"张三"}，age 和 city 被忽略
```

如果要用 DisallowUnknownFields（未知字段报错），用 Decoder：
```go
d := json.NewDecoder(r.Body)
d.DisallowUnknownFields()
d.Decode(&u) // 未知字段会报错
```

# q
json 里的数字反序列化成什么类型？

# a
不指定类型时，json.Unmarshal 到 interface{} 中，数字默认是 float64。

这是经典的坑：
```go
var m map[string]interface{}
json.Unmarshal([]byte(`{"count": 100}`), &m)
count := m["count"].(int) // panic! 实际上是 float64
count := int(m["count"].(float64)) // 正确
```

如果用 json.NewDecoder 并设置 UseNumber，数字会是 json.Number 类型（字符串表示），可以自己转。

# q
```go
type Result struct {
    Items []string `json:"items"`
}
// items 为 null 和 items 为 [] 时，反序列化后分别是什么？
```

# a
- `"items": null` → Items = nil
- `"items": []` → Items = []string{}（空切片，非 nil）

但在 JSON 输出时，nil slice 输出 null，空 slice 输出 []。这个差异有时很重要（API 设计中 null vs [] 语义不同）。

# q
能解析嵌套 JSON 到扁平 struct 吗？

# a
不能直接解析。但可以通过自定义 Unmarshaler 或中间 struct 实现。

```go
// JSON: {"user": {"name": "张三"}}
// 想解析成: type Data struct { Name string }

// 方案1：中间 struct
type Data struct { Name string }
type wrapper struct { User Data }
var w wrapper
json.Unmarshal(data, &w)
name := w.User.Name

// 方案2：自定义 UnmarshalJSON
```

Go 的 JSON 库映射是严格按 struct 层级来的，不支持路径表达式嵌套映射。
