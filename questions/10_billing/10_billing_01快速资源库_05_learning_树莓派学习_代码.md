# q
在树莓派上使用 Python 进行串口通信，如何初始化串口连接？
# a
```py
import serial
ser = serial.Serial('/dev/serial0', 9600)  # 指定串口设备为 /dev/serial0，波特率 9600
```
根据实际硬件连接，可以修改串口号和波特率。

# q
如何从串口读取一行数据并发送响应？
# a
```py
while True:
    if ser.in_waiting > 0:               # 检查接收缓冲区是否有数据
        data = ser.readline().decode('utf-8')  # 读取一行并解码为字符串
        print("Received command:", data)
        ser.write(data_to_send.encode())      # 发送响应数据（需先定义 data_to_send）
```
注意循环体内需要正确的缩进，且发送前应将字符串编码为字节。

