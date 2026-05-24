# q
在 CentOS 系统中通过 yum 安装 Nginx 的命令是什么？
# a
```bash
yum install nginx
```

# q
如何同时安装 Mosquitto MQTT 服务端和客户端？
# a
```bash
yum install -y mosquitto mosquitto-clients
```

# q
启动 Mosquitto 服务并设置开机自启的命令是什么？
# a
```bash
sudo systemctl start mosquitto
sudo systemctl enable mosquitto
```

# q
使用 mosquitto_sub 工具在本地测试 MQTT 订阅的命令是？
# a
```bash
mosquitto_sub -h localhost -t test/topic
```

# q
通过 HACS 为 Home Assistant 安装 Xiaomi Miot Auto 集成的步骤是？
# a
在 HACS 中搜索并下载 “xiaomi miot auto” 集成，然后在设置中重启 Home Assistant 容器（或后台重启容器）。

