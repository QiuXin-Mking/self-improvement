# q
如何使用 curl 发送带请求体的 POST 请求？
# a
使用 `-X POST` 指定 HTTP 方法，用 `-d` 传递请求体数据。若请求体来自文件，可使用 `@文件名` 如：
```bash
curl -X POST -d @data.xml http://example.com/api/endpoint
```
注意：设置 `Content-Type` 等头部应使用 `-H`，而不是 `-d`。

# q
curl 命令中 `-d` 参数的作用是什么？
# a
`-d`（`--data`）用于指定要发送的请求体内容，支持直接跟字符串或通过 `@文件路径` 从文件读取数据（如 `-d @data.xml`）。它会自动将请求方法设为 POST（除非通过 `-X` 显式指定其他方法）。

