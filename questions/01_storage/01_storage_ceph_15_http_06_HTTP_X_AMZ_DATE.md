# q
HTTP_X_AMZ_DATE 是什么？
# a
`HTTP_X_AMZ_DATE` 是一个自定义 HTTP 头部字段，用于 Amazon Web Services (AWS) 的请求签名过程，表示请求发送的日期和时间，用于身份验证和完整性保护。

# q
HTTP_X_AMZ_DATE 的日期时间格式是什么？
# a
遵循 ISO 8601 格式：`YYYYMMDDTHHMMSSZ`，例如 `20230915T121212Z`。

# q
HTTP_X_AMZ_DATE 在 AWS 签名过程中是如何被使用的？
# a
客户端将 `HTTP_X_AMZ_DATE` 加入请求头，AWS 服务端会结合该头部及其他请求参数计算预期签名，并与请求中提供的签名比对。签名匹配则请求有效，否则拒绝。

# q
HTTP_X_AMZ_DATE 是 HTTP 协议标准的一部分吗？
# a
不是，它是 AWS 自定义头部，仅用于与 AWS 服务交互的请求中，不在 HTTP 标准定义内。

