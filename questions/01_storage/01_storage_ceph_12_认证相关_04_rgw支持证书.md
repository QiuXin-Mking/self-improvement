# q
如何为Ceph RGW生成包含完整证书链的PEM文件？
# a
使用 `cat` 命令按顺序合并服务器证书和中间证书、根证书，例如：
```bash
cat server_certificate.crt intermediate_certificate_1.crt intermediate_certificate_2.crt root_certificate.crt > rgw_fullchain.pem
```

# q
如何配置Ceph RGW以启用SSL/TLS支持？
# a
在Ceph配置文件的 `[client.rgw.<rgw_instance_name>]` 部分设置：
```
rgw frontends = beast ssl_port=443 ssl_certificate=/etc/ceph/rgw_fullchain.pem ssl_private_key=/etc/ceph/rgw.key
```
其中 `ssl_certificate` 指向完整的证书链文件，`ssl_private_key` 指向对应的私钥文件。

# q
Ceph RGW的 `ssl_certificate` 参数应包含哪些内容？为什么？
# a
`ssl_certificate` 应包含完整的证书链，即服务器证书、中间证书和根证书的拼接。这样客户端在握手时可以验证整个信任链，确保RGW身份合法，避免信任链断裂导致的连接失败。

