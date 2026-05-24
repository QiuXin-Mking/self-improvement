# q
在基于 Ubuntu 的轻量级服务器上如何安装 Mosquitto MQTT Broker 和客户端？
# a
```bash
sudo apt-get update
sudo apt-get install mosquitto mosquitto-clients
```

# q
Mosquitto 的主配置文件通常位于什么路径？
# a
`/etc/mosquitto/mosquitto.conf`

# q
安装后如何启动 Mosquitto 服务并设为开机自启？
# a
```bash
sudo systemctl start mosquitto
sudo systemctl enable mosquitto
```

# q
如何使用命令行工具测试 Mosquitto 的发布与订阅功能？
# a
- 订阅消息：
  ```bash
  mosquitto_sub -h localhost -t test/topic
  ```
- 发布消息：
  ```bash
  mosquitto_pub -h localhost -t test/topic -m "Hello, MQTT"
  ```

