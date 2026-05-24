# q
如何在基于 Debian 的系统（如树莓派）上安装 minicom？
# a
```bash
sudo apt-get install minicom
```

# q
如何使用 minicom 通过 USB 转串口设备 /dev/ttyUSB1 建立串行通信？
# a
```bash
minicom -D /dev/ttyUSB1
```

