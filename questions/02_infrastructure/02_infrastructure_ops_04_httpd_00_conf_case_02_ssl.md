# q
SSL虚拟主机无法响应HTTPS请求的典型根因是什么？
# a
典型根因包括：未配置 `SSLEngine on`；证书文件路径（`SSLCertificateFile` / `SSLCertificateKeyFile`）错误或文件缺失；未设置 CA 证书链（`SSLCACertificateFile`）导致证书信任链不完整；虚拟主机未监听 443 端口（`<VirtualHost *:443>` 缺失或错误）。

# q
如何从日志定位 Apache SSL 虚拟主机的配置错误？
# a
检查该虚拟主机配置中指定的错误日志路径，例如：
```
ErrorLog /var/log/httpd/example.com-ssl-error.log
```
使用 `tail -f` 或直接查看该文件，重点寻找 SSL 握手失败、证书无效、权限拒绝等报错信息。

# q
配置 Apache SSL 虚拟主机的标准流程是什么？
# a
1. 创建监听端口 443 的虚拟主机：`<VirtualHost *:443>`
2. 指定 ServerName 和 DocumentRoot
3. 启用 SSL 引擎：`SSLEngine on`
4. 配置证书路径：
   ```apache
   SSLCertificateFile /etc/ssl/certs/example.com.crt
   SSLCertificateKeyFile /etc/ssl/private/example.com.key
   SSLCACertificateFile /etc/ssl/certs/ca-bundle.crt
   ```
5. 设置目录访问权限，如：
   ```apache
   <Directory /var/www/example.com/public_html>
       Options Indexes FollowSymLinks
       AllowOverride All
       Require all granted
   </Directory>
   ```
6. 指定访问日志和错误日志：
   ```apache
   ErrorLog /var/log/httpd/example.com-ssl-error.log
   CustomLog /var/log/httpd/example.com-ssl-access.log combined
   ```

