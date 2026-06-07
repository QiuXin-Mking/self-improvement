# q
Go 标准库怎么创建一个简单的 HTTP 服务？

# a
```go
http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("hello world"))
})
http.ListenAndServe(":8080", nil) // nil 表示用 DefaultServeMux
```

实际项目推荐用自己的 mux（避免全局注册污染，方便测试）：
```go
mux := http.NewServeMux()
mux.HandleFunc("/hello", handler)
http.ListenAndServe(":8080", mux)
```

# q
http.ResponseWriter 和 io.Writer 是什么关系？

# a
http.ResponseWriter 包装了 io.Writer，但多了 Header 设置和状态码控制。

注意：调用 Write 之前必须设置好 Header。因为第一次 Write 会隐式发送 Header（状态码 200 + header）。之后再改就不生效了。

```go
func handler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("X-Custom", "value")
    w.WriteHeader(http.StatusOK) // 显式设置状态码
    w.Write([]byte("body"))
}
```

# q
怎么实现 HTTP 中间件（middleware）？

# a
```go
type Middleware func(http.Handler) http.Handler

// 日志中间件
func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        next.ServeHTTP(w, r)
        log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
    })
}

// 使用
mux := http.NewServeMux()
mux.HandleFunc("/hello", handler)
wrapped := loggingMiddleware(mux)
http.ListenAndServe(":8080", wrapped)
```

多个中间件用 `chain(a, b, c)` 或者 `b(a(c(handler)))` 方式组合。

# q
HTTP Handler 和 HandlerFunc 的区别？

# a
Handler 是接口（只有一个方法 ServeHTTP），HandlerFunc 是把普通函数转成 Handler 的类型。

```go
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}

type HandlerFunc func(ResponseWriter, *Request)
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
    f(w, r)
}
// HandlerFunc 实现了 Handler 接口，所以可以作为参数传
```

HandleFunc 是语法糖，内部自动转换。

# q
怎么正确处理 HTTP 请求的 Body 关闭？

# a
```go
resp, err := http.Get(url)
if err != nil {
    return err
}
defer resp.Body.Close() // 必须关闭！

// 读取
body, err := io.ReadAll(resp.Body)
```

服务端不需要手动关 r.Body（框架会关），客户端必须关 resp.Body。

# q
HTTP 服务的优雅关闭怎么做？

# a
```go
srv := &http.Server{Addr: ":8080"}

go func() {
    if err := srv.ListenAndServe(); err != http.ErrServerClosed {
        log.Fatal(err)
    }
}()

// 等信号
quit := make(chan os.Signal, 1)
signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
<-quit

// 优雅关闭（给现有请求一些时间完成）
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()
srv.Shutdown(ctx)
```

关键点：Shutdown 不处理新请求，只等现有请求完成。设置 timeout 防止卡死。

# q
http.Request 怎么获取 URL 参数和 POST body？

# a
```go
func handler(w http.ResponseWriter, r *http.Request) {
    // URL query 参数
    q := r.URL.Query()
    name := q.Get("name")

    // POST form 参数（application/x-www-form-urlencoded）
    r.ParseForm()
    age := r.FormValue("age")

    // JSON body
    var data SomeStruct
    json.NewDecoder(r.Body).Decode(&data)
}
```

# q
http.Client 怎么设置超时？

# a
```go
client := &http.Client{
    Timeout: 10 * time.Second,  // 整个请求的超时
}

// 更细粒度控制
client := &http.Client{
    Transport: &http.Transport{
        DialContext: (&net.Dialer{
            Timeout:   3 * time.Second, // 连接超时
            KeepAlive: 30 * time.Second,
        }).DialContext,
        TLSHandshakeTimeout:   5 * time.Second, // TLS 握手超时
        ResponseHeaderTimeout: 5 * time.Second, // 等待响应头超时
    },
}
```

更好的方式：用 context 控制每个请求的超时。
```go
ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
defer cancel()
req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
resp, err := client.Do(req)
```

# q
怎么处理 HTTP 请求中的大文件上传？

# a
```go
func uploadHandler(w http.ResponseWriter, r *http.Request) {
    r.ParseMultipartForm(32 << 20) // 32MB 内存限制

    file, header, err := r.FormFile("file")
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    defer file.Close()

    // 保存到磁盘
    dst, _ := os.Create("/uploads/" + header.Filename)
    defer dst.Close()
    io.Copy(dst, file)
}
```

ParseMultipartForm 限制用多少内存缓存（超出的存临时文件）。不要无限制接受大文件。
