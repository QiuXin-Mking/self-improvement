# q
树莓派系统镜像通常使用哪个软件烧录？其常见安装路径是什么？
# a
使用 balenaEtcher 烧录。常见安装路径为：`C:\Users\Administrator\AppData\Local\Programs\balena-etcher\balenaEtcher.exe`

# q
远程访问树莓派的两种常用方式及对应工具是什么？
# a
- SSH 连接：使用 PuTTY
- 远程桌面（VNC）：使用 RealVNC Viewer，默认安装路径例如 `C:\Program Files\RealVNC\VNC Viewer\vncviewer.exe`

# q
在树莓派上使用 Python 进行串口通信（读取并回复数据）的典型代码框架是怎样的？
# a
```py
import serial
ser = serial.Serial('/dev/serial0', 9600)  # 根据实际情况修改串口号和波特率

while True:
    if ser.in_waiting > 0:
        data = ser.readline().decode('utf-8')
        print("Received command:", data)
        # 将需要发送的数据编码后发送
        ser.write(data_to_send.encode())
```

