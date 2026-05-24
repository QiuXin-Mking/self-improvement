# q
如何将静态网页文件部署到Nginx的默认Web根目录？
# a
使用以下命令，先创建目录，再复制文件：
```bash
mkdir -p /var/www/html/
sudo cp index.html styles.css script.js /var/www/html/
```

# q
如何启动Nginx并设置开机自启？
# a
```bash
sudo systemctl start nginx
sudo systemctl enable nginx
```

# q
使用UFW防火墙时，如何允许Nginx的HTTP和HTTPS流量？
# a
```bash
sudo ufw allow 'Nginx Full'
```

# q
如何使用mosquitto_pub通过MQTT发布一条包含温度与湿度的JSON消息？
# a
```bash
mosquitto_pub -h localhost -t template -m '{"temperature": 25.5, "humidity": 60}'
```

