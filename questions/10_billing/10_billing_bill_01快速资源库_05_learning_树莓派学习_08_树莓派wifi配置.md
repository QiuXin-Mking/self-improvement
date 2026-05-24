# q
如何在树莓派首次启动时自动启用 SSH 服务？
# a
在烧录好系统的 SD 卡 boot 分区根目录中新建一个名为 `ssh` 的空文件（无扩展名），树莓派启动后会自动开启 SSH。

# q
如何让树莓派在首次启动时自动连接 WiFi 网络？
# a
在 boot 分区根目录新建 `wpa_supplicant.conf` 文件，写入以下配置内容并保存：
```conf
country=CN
ctrl_interface=DIR=/var/run/wpa_supplicant GROUP=netdev
update_config=1
network={
    ssid="你的WiFi名称"
    psk="你的WiFi密码"
    priority=10
}
```
树莓派启动后会自动将文件移动到 `/etc/wpa_supplicant/` 并连接网络。

# q
树莓派已经启动后，还可以通过哪个路径修改 WiFi 配置？
# a
可以在系统中直接编辑 `/etc/wpa_supplicant/wpa_supplicant.conf` 文件来修改 WiFi 连接信息。

