# q
在Apache反向代理配置中，如何禁用正向代理并允许所有来源访问？
# a
在VirtualHost内设置 `ProxyRequests Off` 禁用正向代理，只保留反向代理功能。通过 `<Proxy *>` 块配合 `Order deny,allow` 和 `Allow from all` 指令允许所有来源访问。
```apache
ProxyRequests Off
<Proxy *>
    Order deny,allow
    Allow from all
</Proxy>
```

# q
ProxyPass和ProxyPassReverse指令在反向代理中各自的作用是什么？
# a
`ProxyPass /app http://backend.example.com:8080/app` 将客户端请求的 `/app` 路径转发到后端服务器的对应URL。  
`ProxyPassReverse /app http://backend.example.com:8080/app` 修改后端响应头（如Location、Content-Location），将后端地址映射回代理地址，确保重定向和资源引用正确。
```apache
ProxyPass /app http://backend.example.com:8080/app
ProxyPassReverse /app http://backend.example.com:8080/app
```

# q
在Apache反向代理配置中，如何为代理单独指定错误日志和访问日志？
# a
在VirtualHost内使用 `ErrorLog` 和 `CustomLog` 指令指定独立的日志文件。示例中将错误日志写入 `/var/log/httpd/proxy-error.log`，访问日志以 combined 格式写入 `/var/log/httpd/proxy-access.log`。
```apache
ErrorLog /var/log/httpd/proxy-error.log
CustomLog /var/log/httpd/proxy-access.log combined
```

