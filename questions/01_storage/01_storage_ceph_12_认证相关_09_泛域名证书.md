# q
什么是泛域名证书？
# a
泛域名证书是一种SSL证书，它可以保护指定域名下的所有子域名。泛域名证书使用星号（\*）作为通配符来表示子域名。例如，一个保护 `*.example.com` 的泛域名证书可以同时保护 `www.example.com`、`mail.example.com`、`blog.example.com` 等直接子域名。

# q
`*.example.com` 和 `*.*.example.com` 在保护范围上有何区别？
# a
- `*.example.com` 保护 `example.com` 及其所有直接子域名，如 `www.example.com`，但不包括更深层次的子域名，如 `sub.mail.example.com`。
- `*.*.example.com` 保护更深一层的子域名，如 `sub.mail.example.com` 或 `test.blog.example.com`，但此类证书比较少见且不常用，具体需要视具体的证书颁发机构（CA）是否支持这种格式。

