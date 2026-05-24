# q
手机热点的WiFi名称和密码是什么？
# a
- 名称：`qx`
- 密码：`123456789`

# q
如何通过SSH登录树莓派？
# a
使用命令 `ssh pi@192.168.236.85`，密码为 `raspberry`。

# q
树莓派上VNC登录的凭据是什么？
# a
用户名 `pi`，密码 `raspberry`。

# q
ESP32向树莓派输出的JSON数据格式是什么？
# a
```json
{
    "t":23.45,
    "m":[1, 2, 3]
}
```

# q
基于树莓派、ESP32和传感器的系统连接关系是怎样的？
# a
```
mlx90640 --(iic)--> esp32 --(uart)--> raspberrypi --> show
mlx90916 --(iic)--> esp32 --(uart)--> raspberrypi --> show
csi --> raspberrypi --> show
```

