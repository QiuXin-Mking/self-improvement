# q
如何创建 Nginx 默认的静态文件目录并复制网站文件？
# a
```bash
mkdir -p /var/www/html/
sudo cp index.html styles.css script.js /var/www/html/
```

# q
启动 Nginx 服务并配置开机自启的命令是什么？
# a
```bash
sudo systemctl start nginx
sudo systemctl enable nginx
```

# q
如何配置防火墙以允许 Nginx 流量？
# a
```bash
sudo ufw allow 'Nginx Full'
```

# q
自定义 Nginx 站点配置文件的存放路径是什么？
# a
```
/etc/nginx/sites-available/custom_site
```

