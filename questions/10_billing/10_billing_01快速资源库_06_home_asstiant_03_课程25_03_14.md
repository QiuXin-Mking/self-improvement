# q
如何在CentOS系统上安装Mosquitto MQTT服务器并进行简单的订阅测试？
# a
安装命令：
```bash
yum install -y mosquitto mosquitto-clients
```
配置文件位置：`/etc/mosquitto/mosquitto.conf`。启动并设置开机自启：
```bash
sudo systemctl start mosquitto
sudo systemctl enable mosquitto
systemctl status mosquitto
```
测试本地订阅：
```bash
mosquitto_sub -h localhost -t test/topic
```

# q
在Home Assistant中如何通过HACS安装Xiaomi Miot Auto集成？
# a
在HACS界面中搜索并下载“xiaomi miot auto”。安装完成后，在Home Assistant设置中重启容器（或通过后台重启容器），使集成生效。

