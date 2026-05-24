# q
使用OpenSSL生成RSA私钥的命令是什么？`-aes256`参数的作用是什么？
# a
```bash
openssl genpkey -algorithm RSA -out my_private_key.key -aes256
```
`-algorithm RSA` 指定使用RSA算法，`-aes256` 使用AES256算法加密私钥文件，以增强安全性。输出文件为 `my_private_key.key`。

# q
如何使用OpenSSL从证书签名请求(CSR)生成自签名证书？
# a
```bash
openssl x509 -req -days 365 -in my_csr.csr -signkey my_private_key.key -out my_certificate.crt
```
- `-req` 表示输入是一个CSR文件
- `-days 365` 设置证书有效期为365天
- `-in` 指定CSR文件路径
- `-signkey` 指定用于签名的私钥
- `-out` 指定输出的证书文件

# q
在Ceph RGW中如何配置自签名SSL证书？
# a
1. 将私钥和证书文件放置到 `/etc/ceph/` 目录下；
2. 在 `ceph.conf` 中添加以下配置：
```ini
[client.rgw.rgw.myinstance]
rgw frontends = beast ssl_port=443 ssl_certificate=/etc/ceph/my_certificate.crt ssl_private_key=/etc/ceph/my_private_key.key
```
3. 重启RGW服务：
```bash
sudo systemctl restart ceph-radosgw@rgw.myinstance
```

