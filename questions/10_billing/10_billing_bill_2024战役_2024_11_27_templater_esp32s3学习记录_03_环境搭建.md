# q
在Arduino IDE中为ESP32-S3添加硬件支持时，需要在Sketchbook目录下创建什么文件夹结构？
# a
在Sketchbook目录（默认为 ```C:\Users\<用户名>\Documents\Arduino```）下，创建 ```hardware\esp32\esp32``` 子目录，并将arduino-esp32的核心文件放入其中。例如：```C:\Users\Administrator\Documents\Arduino\hardware\esp32\esp32```

# q
在Arduino IDE中开发ESP32-S3时，应选择哪个开发板？
# a
在“工具” → “开发板”菜单中，选择 ```ESP32S3 Dev Module```。

# q
arduino-esp32 的官方GitHub仓库地址及其本地存放路径是什么？
# a
官方仓库地址：```https://github.com/espressif/arduino-esp32```  
本地建议克隆到：```C:\Users\Administrator\Documents\Arduino\hardware\esp32\esp32```

# q
ESP32-S3 Arduino示例（测试代码）在本地文件系统中的路径是什么？
# a
```C:\Users\Administrator\Documents\Arduino\hardware\esp32\esp32\tests```

