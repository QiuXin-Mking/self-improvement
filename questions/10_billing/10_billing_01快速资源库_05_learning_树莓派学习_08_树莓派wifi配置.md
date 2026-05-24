# q
在树莓派的无头模式（无键盘/显示器）下，如何让系统启动时自动开启SSH服务？
# a
在SD卡的boot分区根目录下新建一个名为`ssh`的空文件（无扩展名）。系统启动时会检测该文件并启用SSH。

# q
如何让树莓派启动后自动连接到指定的WiFi网络？给出boot分区中的配置文件名及其内容结构。
# a
在SD卡的boot分区中新建`wpa_supplicant.conf`文件，写入以下内容（根据实际情况替换ssid和psk）：
```
country=CN
ctrl_interface=DIR=/var/run/wpa_supplicant GROUP=netdev
update_config=1
network={
    ssid="你的WiFi名称"
    psk="你的WiFi密码"
    priority=10
}
```
若需多个网络，可添加多个`network={}`块。

# q
树莓派系统中WiFi配置文件`wpa_supplicant.conf`在运行时的默认路径是什么？
# a
```
/etc/wpa_supplicant/wpa_supplicant.conf
```

