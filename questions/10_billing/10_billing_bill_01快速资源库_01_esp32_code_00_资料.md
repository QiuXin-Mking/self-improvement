# q
哪些ESP32 Arduino项目使用了Blinker系统并通过串口2控制外设？
# a
13_节水系统_串口2_按键控制.ino、14_晾衣架系统_串口2_按键控制.ino、07_oled_dht11_blinker.ino、15_光伏加热系统.ino 均使用 Blinker 系统，通过 app 点击控制串口2发送命令字。

# q
07_oled_dht11_blinker.ino 项目在 Blinker 平台上实现了哪些数据交互功能？
# a
该项目支持温湿度上报和滑块上报，同时通过 app 点击控制串口2发送命令字。

# q
用于学习 ESP32 与 DHT11 传感器的官方文档和在线仿真链接分别是什么？
# a
文档：https://docs.geeksman.com/esp32/Arduino/29.esp32-arduino-dht11.html  
仿真：https://wokwi.com/projects/322577683855704658

