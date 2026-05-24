# q
`SSL_ERROR_RX_RECORD_TOO_LONG` 错误通常表示什么？
# a
表示客户端收到的 SSL/TLS 记录长度超过了最大允许值，常见原因是目标端口上运行的不是 HTTPS 服务（例如是纯 HTTP 服务），或者服务端 SSL 配置错误导致返回了非 TLS 数据。

# q
如何使用 `openssl s_client` 命令验证 HTTPS 服务？
# a
```bash
openssl s_client -connect 127.0.0.1:8085
```
该命令会尝试建立到指定主机和端口的 SSL/TLS 连接，并显示证书链和握手详情，用于诊断服务端 TLS 配置是否正确。

# q
如何使用 `curl` 命令通过客户端证书访问 HTTPS 服务？
# a
```bash
curl -v --cert /path/to/client.crt --key /path/to/client.key https://127.0.0.1:8085
```
`--cert` 指定客户端证书文件，`--key` 指定对应的私钥文件，`-v` 输出详细连接过程，用于调试基于证书的认证。

