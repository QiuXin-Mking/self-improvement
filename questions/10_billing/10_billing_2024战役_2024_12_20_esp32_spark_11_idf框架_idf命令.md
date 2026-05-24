# q
如何设置 ESP-IDF 工程的目标芯片为 ESP32-S3？
# a
```bash
idf.py set-target esp32s3
```
如果遇到 “Refusing to automatically delete files” 错误，请手动删除 build 文件夹后再重新执行该命令。

# q
如何启动 ESP-IDF 工程的图形化配置界面？
# a
```bash
idf.py menuconfig
```
启动后可在界面中搜索并选择组件（如 BMI270），按 `/` 搜索，选择后保存退出。

# q
如何编译 ESP-IDF 工程？
# a
```bash
idf.py build
```

# q
如何将编译好的固件烧录到 ESP32 设备（指定 COM22 端口）？
# a
```bash
idf.py -p COM22 flash
```

# q
如何擦除 ESP32 的整个闪存？
# a
```bash
idf.py erase-flash
```

