# q
如何将静态网页文件（index.html, styles.css, script.js）部署到 Nginx 默认文档根目录？
# a
使用以下命令创建目录并复制文件：
```bash
mkdir -p /var/www/html/
sudo cp index.html styles.css script.js /var/www/html/
```

# q
如何启动 Nginx 并设置开机自启？
# a
执行：
```bash
sudo systemctl start nginx
sudo systemctl enable nginx
```

# q
如何配置防火墙以允许 Nginx 的 HTTP/HTTPS 流量？
# a
使用 UFW 允许 Nginx Full 配置文件：
```bash
sudo ufw allow 'Nginx Full'
```
此命令会同时开放 80 和 443 端口。

