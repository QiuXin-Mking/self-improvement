# q
如何发送一个带有 JSON 数据的 POST 请求？
# a
```bash
curl -X POST \
     -H "Content-Type: application/json" \
     -d '{"name":"John", "age":30}' \
     https://www.example.com/api/users
```

# q
如何下载文件并指定保存的文件名？
# a
```bash
curl -o myfile.zip https://www.example.com/files/myfile.zip
```

# q
如何仅获取 HTTP 响应状态码？
# a
```bash
curl -s -o /dev/null -w "%{http_code}\n" https://www.example.com
```

# q
如何忽略 SSL 证书验证进行请求？
# a
```bash
curl -k https://untrusted-root.badssl.com/
```

# q
如何为 curl 请求设置最大超时时间？
# a
```bash
curl --max-time 30 https://www.example.com
```

