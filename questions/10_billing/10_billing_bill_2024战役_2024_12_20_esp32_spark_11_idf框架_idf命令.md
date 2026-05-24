# q
如何为 ESP32-S3 项目设置芯片目标？
# a
在工程目录下运行：
```bash
idf.py set-target esp32s3
```
如果出现 “doesn't seem to be a CMake build directory” 错误，需先手动删除 `build` 目录，然后重新执行该命令。

# q
如何编译 ESP-IDF 工程？
# a
```bash
idf.py build
```

# q
如何将固件烧录到 ESP32 设备？
# a
使用 `-p` 指定串口号，例如：
```bash
idf.py -p COM22 flash
```

# q
如何启动串口监视器查看设备输出？
# a
```bash
idf.py monitor
```

# q
如何擦除 ESP32 闪存？
# a
```bash
idf.py erase-flash
```

