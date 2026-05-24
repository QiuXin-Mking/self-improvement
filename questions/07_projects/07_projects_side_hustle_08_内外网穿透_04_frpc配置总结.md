# q
在 frpc 配置中，如何指定一个名为 `rdp-mapping` 的 TCP 代理，将本地 3389 端口映射到远程服务器的指定端口？
# a
在 `[[proxies]]` 块中设置：
- `name = "rdp-mapping"`
- `type = "tcp"`
- `local_ip = "127.0.0.1"`
- `local_port = 3389`
- `remote_port = 你的远程端口号`

同时需在文件顶部设置 `server_addr` 和 `server_port` 指向 frps 服务器。

# q
frpc 配置中 `server_addr` 和 `serverAddr` 有什么区别？
# a
两者功能相同，都用于指定 frps 服务器的 IP 地址，但属于不同版本的配置键名：
- `server_addr` 是 TOML 风格的下划线命名
- `serverAddr` 是驼峰命名风格
需根据 frpc 版本选择对应的键名，不可混用。

# q
在 frpc 默认配置示例中，`localIP` 和 `localPort` 分别设置为多少，含义是什么？
# a
`localIP = "127.0.0.1"`，表示代理本地回环地址；`localPort = 22`，表示代理本地的 SSH 服务端口。

# q
frpc 配置中 `remotePort` 的作用是什么？
# a
`remotePort` 是 frps 服务器上监听的端口，客户端连接服务器该端口时，流量会被转发到 frpc 对应的本地服务（`localIP:localPort`）。

