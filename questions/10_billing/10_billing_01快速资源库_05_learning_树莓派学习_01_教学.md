# q
树莓派镜像烧录软件 balenaEtcher 的默认安装路径是什么？
# a
```
C:\Users\Administrator\AppData\Local\Programs\balena-etcher\balenaEtcher.exe
```

# q
用于树莓派远程 SSH 连接的 PuTTY 工具在资料盘中的哪个目录？
# a
```
E:\018_树莓派\树莓派入门套件A资料盘 2021-10-11\2.开发环境\系统镜像
```

# q
1200 bit/s 的传输速率换算成 M/s 是多少？
# a
计算过程：
1200 bit/s ÷ 8 = 150 Byte/s
150 ÷ 1024 ≈ 0.1465 KB/s
0.1465 ÷ 1024 ≈ 0.000143 M/s
忽略部分计算误差，最终约为 0.000143 M/s

# q
ASCII 码的取值范围对应的十六进制范围是多少？
# a
0x00 - 0xFF（即 0 到 255）

# q
使用 Python 的 serial 库从串口读取一行数据并发送回复的代码示例是什么？
# a
```py
import serial
ser = serial.Serial('/dev/serial0', 9600) # 根据实际情况修改串口号和波特率
while True:
    if ser.in_waiting > 0:
        data = ser.readline().decode('utf-8') 
        print("Received command:", data)
        # 发送数据
        ser.write(data_to_send.encode()) 
```

