# q
如何在 Python 中通过串口连接树莓派的 `/dev/serial0` 并设置波特率？
# a
使用 `serial` 模块创建串口对象，示例代码如下：
```python
import serial
ser = serial.Serial('/dev/serial0', 9600)
```
其中 `'/dev/serial0'` 是树莓派默认的串口设备路径，`9600` 是波特率，需根据实际设备调整。

# q
在串口循环读取数据时，如何判断缓冲区是否有数据并正确解码？
# a
通过 `ser.in_waiting` 判断是否有待读取的字节，若大于 0 则调用 `ser.readline()` 读取一行并使用 UTF-8 解码：
```python
if ser.in_waiting > 0:
    data = ser.readline().decode('utf-8')
```

# q
如何通过串口发送数据，需要注意什么？
# a
发送数据前需将字符串编码为字节，使用 `encode()` 方法，然后通过 `ser.write()` 写入串口：
```python
ser.write(data_to_send.encode())
```

