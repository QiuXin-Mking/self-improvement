# q
在树莓派上安装 Home Assistant OS 时，官方推荐的镜像烧录工具是什么？
# a
使用官方提供的烧录器 `imager_1.8.5.exe`，配合镜像文件 `haos_rpi4-64-14.2.img.xz`（位于 `C:\Users\Administrator\Downloads\`）进行烧录。

# q
当前 Home Assistant 设备的网络连接有什么限制？
# a
只能通过有线网络连接，无法通过无线连接。

# q
如何通过浏览器访问 Home Assistant 的 Web 界面？
# a
首先通过命令 `ping -4 homeassistant` 确认主机可达，然后在浏览器中输入 `http://homeassistant.local:8123/` 访问。

