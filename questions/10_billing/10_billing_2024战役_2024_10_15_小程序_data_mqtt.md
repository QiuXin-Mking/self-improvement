# q
在轻量级服务器上推荐使用哪个MQTT Broker软件，为什么？
# a
推荐使用Eclipse Mosquitto，因为它是一个广泛使用的开源MQTT Broker，轻量级且易于配置，适合资源有限的设备。

# q
在Ubuntu系统上如何安装Mosquitto Broker及客户端工具？
# a
```bash
sudo apt-get update
sudo apt-get install mosquitto mosquitto-clients
```

# q
安装Mosquitto后，其主配置文件通常位于哪个路径？
# a
`/etc/mosquitto/mosquitto.conf`

# q
如何启动Mosquitto服务并设置为开机自启？
# a
```bash
sudo systemctl start mosquitto
sudo systemctl enable mosquitto
```

# q
如何使用命令行测试MQTT Broker的消息发布与订阅功能？
# a
发布消息：
```bash
mosquitto_pub -h localhost -t test/topic -m "Hello, MQTT"
```
订阅消息：
```bash
mosquitto_sub -h localhost -t test/topic
```
在两个终端分别运行，验证消息能否被正确接收。

