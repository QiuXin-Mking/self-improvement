# q
如何安装Mosquitto MQTT服务并测试订阅？
# a
使用命令 `yum install -y mosquitto mosquitto-clients` 安装，启动并设置开机自启：
```
sudo systemctl start mosquitto
sudo systemctl enable mosquitto
```
测试订阅：`mosquitto_sub -h localhost -t test/topic`

# q
如何查看Mosquitto的配置文件？
# a
使用 `cat /etc/mosquitto/mosquitto.conf` 查看配置文件。

# q
在Home Assistant中通过HACS安装xiami miot auto后需要做什么？
# a
在HACS中下载 `xiaomi miot auto` 后，需要在设置中重启Home Assistant容器。

