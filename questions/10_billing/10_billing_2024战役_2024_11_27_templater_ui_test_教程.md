# q
部署一个静态HTML网站的基本Nginx步骤是什么？
# a
1. 创建Web根目录：`mkdir -p /var/www/html/`
2. 复制静态文件：`sudo cp index.html styles.css script.js /var/www/html/`
3. 启动Nginx并设置开机自启：
   ```bash
   sudo systemctl start nginx
   sudo systemctl enable nginx
   ```

# q
如何允许Nginx通过UFW防火墙？
# a
执行 `sudo ufw allow 'Nginx Full'`，这会开放HTTP(80)和HTTPS(443)端口。

# q
如何配置Nginx将请求代理到本地5000端口的应用？
# a
在 `/etc/nginx/sites-available/custom_site` 配置文件中添加 `proxy_pass http://127.0.0.1:5000;` 指令，然后启用该站点并重载Nginx。

# q
如何使用mosquitto_pub发送一条包含温度和湿度的MQTT消息？
# a
使用命令：
```bash
mosquitto_pub -h localhost -t template -m '{"temperature": 25.5, "humidity": 60}'
```
- `-h` 指定MQTT broker主机
- `-t` 指定主题（template）
- `-m` 指定消息内容（JSON格式）

