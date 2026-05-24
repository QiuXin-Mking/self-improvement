# q
树莓派默认的登录用户名和密码是什么？
# a
账号：pi  
密码：raspberry

# q
如何在树莓派上通过命令行启用 VNC 服务？
# a
使用 SSH 工具（如 Putty）连接到树莓派后，执行以下命令：
```bash
sudo raspi-config
```
然后选择 “5 Interfacing Options” 进行外设配置。

# q
在 Windows 命令行中如何通过主机名 Ping 树莓派？
# a
使用命令：
```cmd
ping raspberrypi.local
```
系统会通过 mDNS 解析并返回树莓派的 IPv6 或 IPv4 地址。

