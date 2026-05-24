# q
如何在 CentOS 上安装 Mosquitto MQTT 服务端和客户端？
# a
```bash
yum install -y mosquitto mosquitto-clients
```

# q
如何启动 Mosquitto 并设置开机自启？
# a
```bash
sudo systemctl start mosquitto
sudo systemctl enable mosquitto
```

# q
如何用命令行测试本地 MQTT 订阅功能？
# a
```bash
mosquitto_sub -h localhost -t test/topic
```

