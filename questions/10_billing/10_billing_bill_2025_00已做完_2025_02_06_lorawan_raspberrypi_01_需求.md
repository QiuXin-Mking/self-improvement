# q
什么是LoRa和LoRaWAN？它们之间有什么关系？
# a
LoRa（Long Range）是由Semtech公司研发的一种低功耗、远距离的局域网无线调制技术，属于物理层。  
LoRaWAN（Long Range Wide Area Network）是基于LoRa技术构建的低功耗广域通信协议，由LoRa联盟维护，定义了LoRa网络的物理层和MAC层。LoRa是物理层技术，LoRaWAN是在其上构建的协议。

# q
树莓派通过LoRaWAN连接到Home Assistant的整体架构是怎样的？
# a
架构路径为：树莓派 -> LoRaWAN模块 -> LoRaWAN网关 -> 服务器（运行Home Assistant） -> 手机App。

# q
如何使用Python在树莓派上控制GPIO 17引脚上的LED闪烁？
# a
以下代码使用`RPi.GPIO`库实现LED每隔1秒亮灭一次：
```python
import RPi.GPIO as GPIO
import time

GPIO.setmode(GPIO.BCM)
led_pin = 17
GPIO.setup(led_pin, GPIO.OUT)

try:
    while True:
        GPIO.output(led_pin, GPIO.HIGH)
        print("LED is ON")
        time.sleep(1)
        GPIO.output(led_pin, GPIO.LOW)
        print("LED is OFF")
        time.sleep(1)
except KeyboardInterrupt:
    GPIO.cleanup()
    print("Program stopped and GPIO cleaned up")
```

# q
文档中提到的Home Assistant实例的访问地址和登录凭证是什么？
# a
访问地址：`homeassistant-qx.com`  
用户名：`qiuxin`  
密码：`qiuxin@MK@159`

