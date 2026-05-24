# q
如何通过SSH隧道从本地计算机访问远程服务器的Clawdbot仪表盘？
# a
在本地终端执行以下命令建立SSH隧道：
```bash
ssh -N -L 18789:127.0.0.1:18789 root@<服务器IP>
```
然后在浏览器中打开 `http://localhost:18789/` 即可访问。

# q
SSH隧道命令执行后没有任何输出或无法连接，如何进一步调试？
# a
在SSH命令中添加 `-v`（verbose）参数查看详细连接过程：
```bash
ssh -v -N -L 18789:127.0.0.1:18789 root@<服务器IP>
```

# q
如何将Clawdbot网关配置为systemd服务，实现后台持续运行与开机自启？
# a
1. 创建服务文件 `/etc/systemd/system/clawdbot.service`，内容如下：
```ini
[Unit]
Description=Clawdbot Gateway Service
After=network.target
Wants=network.target

[Service]
Type=simple
User=root
WorkingDirectory=/root
ExecStart=/usr/bin/clawdbot gateway --verbose
Restart=always
RestartSec=5
StandardOutput=journal
StandardError=journal
Environment=CLAWDBOT_GATEWAY_PORT=18789
Environment=DASHSCOPE_API_KEY=sk-e160a2e940374a98bf60a8105da3599c

[Install]
WantedBy=multi-user.target
```
2. 依次执行：
```bash
sudo systemctl daemon-reload
sudo systemctl enable clawdbot
sudo systemctl start clawdbot
```

# q
使用什么命令可以检查Clawdbot网关进程是否正在监听18789端口？
# a
执行：
```bash
netstat -ntlp | grep 18789
```
若端口被监听，会显示对应的进程信息。

