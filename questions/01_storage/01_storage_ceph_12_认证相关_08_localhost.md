# q
localhost 是什么？
# a
localhost 是一个主机名，通常映射到回环地址 127.0.0.1，用于指向本机。在 Linux 系统中通过 `/etc/hosts` 文件配置此映射。

# q
如何理解 `ping localhost` 能够通？
# a
因为 `/etc/hosts` 中将 `localhost` 解析为 `127.0.0.1`，这是系统预留的回环地址，数据包不会发送到网络，而是在 TCP/IP 协议栈内部直接返回，所以必定能 ping 通。

# q
`/etc/hosts` 文件的作用是什么？
# a
`/etc/hosts` 是本地静态主机名解析文件，用于将主机名映射到 IP 地址，优先级通常高于 DNS 解析。修改此文件不需要重启网络服务即可立即生效。

