# q
在 Debian/Ubuntu 系统上如何安装 minicom 串口通信工具？
# a
```bash
sudo apt-get install minicom
```

# q
如何使用 minicom 连接指定的 USB 转串口设备（例如 /dev/ttyUSB1）？
# a
```bash
minicom -D /dev/ttyUSB1
```
`-D` 参数指定要打开的串行设备文件。

