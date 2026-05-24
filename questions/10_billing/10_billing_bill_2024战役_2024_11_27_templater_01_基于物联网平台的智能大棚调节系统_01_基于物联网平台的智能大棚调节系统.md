# q
基于物联网平台的智能大棚调节系统能够实现哪些自动化调节功能？
# a
获取环境温度，低于阈值自动加热；获取环境湿度，低于阈值自动加水；获取环境亮度，超过阈值自动开灯；并将参数显示在Web UI和OLED上。

# q
该系统使用了哪些核心硬件器件？
# a
ESP32-S3开发板（ESP32S3 DEV MODULE）、IIC OLED、DHT11模块、继电器。

# q
MQTT公有云服务的端口和主题分别是什么？
# a
MQTT端口为1883，主题为`template1`。

# q
如何启动该系统的Web UI服务？
# a
```bash
cd /var/www/my_project
source /root/cpp_project/python_project/venv/bin/activate
python3 test1.py
```

# q
设备端连接MQTT Broker的服务器地址是什么？
# a
`121.41.87.58`

