# q
如何为 Linux 上的 frps 编写 systemd 服务单元文件，使其作为后台服务运行？
# a
创建文件 `/etc/systemd/system/frps.service`，内容如下：
```ini
[Unit]
Description = frp server
After = network.target syslog.target
Wants = network.target

[Service]
Type = simple
ExecStart = /root/frp_0.61.0_linux_amd64/frps -c /root/frp_0.61.0_linux_amd64/frps.toml

[Install]
WantedBy = multi-user.target
```
然后执行 `systemctl daemon-reload` 和 `systemctl enable --now frps` 即可启用和启动服务。

# q
在 Windows 上如何将 frpc 配置为登录后自动隐藏运行？
# a
1. 按 `Win + R`，输入 `taskschd.msc` 打开任务计划程序。
2. 创建基本任务，触发设为“当用户登录时”，操作选择“启动程序”。
3. 程序路径填写 `frpc.exe` 的完整路径，参数填写 `-c frpc.toml`。
4. 在任务属性中勾选“隐藏”窗口（或使用“不管用户是否登录都要运行”配合“不存储密码”等配置，通常选“仅当用户登录时运行”并勾选“隐藏”即可）。

