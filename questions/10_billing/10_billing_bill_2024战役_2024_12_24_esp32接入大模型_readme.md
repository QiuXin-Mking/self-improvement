# q
ESP32 S3 接入大模型的音频硬件方案由哪些组件组成？
# a
- ESP32 S3 主控  
- MAX98357 音频放大器（用于增益）  
- 麦克风  
- 喇叭

# q
MAX98357 与 ESP32 的典型接线是怎样的？
# a
``` 
VIN  → 5V
GND → GND
DIN  → GPIO 25
BCLK → GPIO 27
LRC  → GPIO 26
```

# q
小智固件的获取地址是什么？本文推荐的稳定版本是什么？
# a
小智固件 GitHub 发布地址：https://github.com/78/xiaozhi-esp32/releases  
推荐版本：v0.9.2

# q
小智 AI 后台配置面板的访问地址是什么？
# a
https://xiaozhi.me/

# q
接入豆包大模型可以参考哪些资料？
# a
- https://blog.csdn.net/wxl781227/article/details/137789379  
- https://blog.csdn.net/vor234/article/details/140620752  
- 豆包大模型官方产品页：https://www.volcengine.com/product/doubao

