# q
Go 怎么写单元测试？有哪几种测试类型？

# a
Go 有三种测试类型：

1. **单元测试**（`_test.go`, TestXxx 函数）
2. **基准测试**（BenchmarkXxx 函数）
3. **示例测试**（ExampleXxx 函数，产物可作为文档）

```go
// 单元测试
func TestAdd(t *testing.T) {
    got := Add(1, 2)
    want := 3
    if got != want {
        t.Errorf("Add(1, 2) = %d; want %d", got, want)
    }
}

// 基准测试
func BenchmarkAdd(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Add(1, 2)
    }
}
```

运行：`go test ./...` / `go test -bench=.` / `go test -coverprofile=coverage.out`

# q
什么是 table-driven tests？为什么推荐？

# a
把多个测试用例放在一个表格里，循环跑。

```go
func TestAdd(t *testing.T) {
    tests := []struct {
        name string
        a, b int
        want int
    }{
        {"positive", 1, 2, 3},
        {"zero", 0, 0, 0},
        {"negative", -1, -2, -3},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := Add(tt.a, tt.b)
            if got != tt.want {
                t.Errorf("got %d, want %d", got, tt.want)
            }
        })
    }
}
```

好处：增删用例只需改一行，每个 case 独立运行（t.Run），失败时一眼看哪个 case 挂了。

# q
```go
func TestSomething(t *testing.T) {
    // 1. t.Fatal 和 t.Error 有什么区别？
    // 2. t.Run 有什么好处？
    // 3. t.Parallel 干什么的？
}
```

# a
1. t.Fatal → 立即终止当前测试；t.Error → 标记失败但继续执行
2. t.Run → 子测试，每个子测试独立命名，可单独运行 `go test -run TestSomething/子测试名`
3. t.Parallel → 标记测试可以并行运行（和其他带 Parallel 的测试一起跑）

```go
func TestFoo(t *testing.T) {
    t.Parallel() // 并行跑
    // ...
}
```

# q
怎么 mock / stub 外部依赖？

# a
Go 靠接口实现可测试性。依赖注入接口，测试时传 mock。

```go
// 定义接口
type Store interface {
    Get(key string) (string, error)
}

// 真实实现
type RedisStore struct { client *redis.Client }

// Mock
type MockStore struct {
    Data map[string]string
}
func (m *MockStore) Get(key string) (string, error) {
    v, ok := m.Data[key]
    if !ok {
        return "", ErrNotFound
    }
    return v, nil
}

// 使用接口
func GetUser(s Store, id string) (*User, error) {
    data, err := s.Get(id)
    // ...
}

// 测试
mock := &MockStore{Data: map[string]string{"1": `{"name":"张三"}`}}
user, err := GetUser(mock, "1")
```

社区 mock 工具：gomock、mockery（自动生成 mock）。

# q
怎么测试 HTTP handler？

# a
用 httptest 包，不需要启动真实服务器。

```go
func TestHandler(t *testing.T) {
    req := httptest.NewRequest("GET", "/hello?name=张三", nil)
    w := httptest.NewRecorder()

    handler(w, req)

    resp := w.Result()
    body, _ := io.ReadAll(resp.Body)

    if resp.StatusCode != http.StatusOK {
        t.Errorf("status %d", resp.StatusCode)
    }
}
```

# q
基准测试怎么跑？输出怎么解读？

# a
```bash
go test -bench=. -benchmem -benchtime=3s
```

输出：
```
BenchmarkAdd-8    1000000000    0.25 ns/op    0 B/op    0 allocs/op
```

- `-8`：用了 8 个 CPU
- `1000000000`：跑了 10 亿次
- `0.25 ns/op`：每次操作耗时 0.25 纳秒
- `0 B/op`：每次操作 0 字节内存分配
- `0 allocs/op`：每次操作 0 次堆分配

# q
怎么生成和查看代码覆盖率？

# a
```bash
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out  # 浏览器可视化
go tool cover -func=coverage.out  # 每个函数的覆盖率
```

但覆盖率不是万能的——100% 覆盖不等于没 bug，只是每行代码都被跑过一次。

# q
怎么测试并发代码？

# a
```go
func TestConcurrent(t *testing.T) {
    var wg sync.WaitGroup
    for i := 0; i < 100; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            // 操作共享资源
        }()
    }
    wg.Wait()
    // 验证结果
}
```

更重要的是用 race detector：`go test -race ./...`

# q
TestMain 是干什么的？

# a
TestMain 是测试的主入口，可以在所有测试前后做 setup/teardown。

```go
func TestMain(m *testing.M) {
    // setup（启动 DB、创建临时目录等）
    setup()
    
    code := m.Run() // 运行所有测试
    
    // teardown（清理资源）
    teardown()
    
    os.Exit(code)
}
```

一个包最多一个 TestMain。
