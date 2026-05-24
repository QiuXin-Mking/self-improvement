# q
树莓派的默认登录账号和密码是什么？
# a
账号：`pi`  
密码：`raspberry`

# q
如何在 Windows 命令提示符中查找树莓派的 IP 地址？
# a
在 cmd 中执行 `ping raspberrypi` 或 `ping raspberrypi.local`，从返回结果中可看到树莓派的 IP 地址（可能是 IPv4 如 `192.168.95.153` 或 IPv6 如 `fe80::6f78:f108:2291:c985%9`）。

# q
如何通过 SSH 开启树莓派的 VNC 服务？
# a
使用 putty 等工具 SSH 登录树莓派后，执行以下命令：
```bash
sudo raspi-config
```
在菜单中选择 `5 Interfacing Options` 进行 VNC 配置。

