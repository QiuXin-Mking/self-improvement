# q
在ceph认证中，客户端通常需要哪些文件？
# a
客户端需要 SSL 证书文件 `client.crt`、私钥文件 `client.key`，以及 CA 证书文件 `ca.crt`（如果使用自签名或自定义 CA）。

# q
在ceph认证中，服务器端通常需要哪些文件？
# a
服务器端需要 `server.crt`（服务器证书）、`server.key`（服务器私钥），以及 `ca-bundle.crt`（CA 证书文件，用于验证服务器证书的 CA 链）。

# q
`ca-bundle.crt` 文件的作用是什么？
# a
`ca-bundle.crt` 是 CA 证书文件，用于验证服务器证书的 CA 链。

# q
`client.crt` 文件是什么？
# a
`client.crt` 是客户端的 SSL 证书文件。

# q
`client.key` 文件是什么？
# a
`client.key` 是客户端的私钥文件。

