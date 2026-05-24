# q
这些ESP32项目中使用Blinker物联网平台主要实现了哪些控制与上报功能？
# a
通过Blinker App点击控制设备，经**串口2**发送命令字；同时支持通过**滑块**上报传感器数据（如温湿度）到App。

# q
哪个示例文件集成了OLED显示屏和DHT11温湿度传感器，并通过Blinker上报数据？
# a
`07_oled_dht11_blinker.ino` 集成了OLED显示，并通过Blinker App滑块上报温湿度数据。

