# q
树莓派烧录Home Assistant OS镜像时，使用什么软件工具？
# a
使用 imager_1.8.5.exe 烧录器。

# q
本例中使用的 Home Assistant OS 镜像文件是什么？
# a
镜像源文件为 `C:\Users\Administrator\Downloads\haos_rpi4-64-14.2.img.xz`。

# q
烧录完成后，如何通过主机名访问该树莓派设备？
# a
使用命令 `ping -4 homeassistant` 可解析并测试其 IPv4 连通性，设备默认主机名为 `homeassistant`。

