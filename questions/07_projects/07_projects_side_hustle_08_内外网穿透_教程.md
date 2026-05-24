# q
如何使用SSH创建一个反向隧道，将家庭电脑的22端口映射到公有云服务器的指定端口？
# a
使用命令：
```bash
ssh -R <remote_port>:localhost:22 user@cloud_server_ip
```
例如：
```bash
ssh -R 3306:localhost:22 root@121.41.87.58
```
其中 `<remote_port>` 是公有云服务器上开放的端口（如 3306），`user` 是公有云用户名，`cloud_server_ip` 是公有云服务器 IP。

# q
在 `ssh -R <remote_port>:localhost:22 user@cloud_server_ip` 命令中，`<remote_port>` 的作用是什么？
# a
`<remote_port>` 指定在公有云服务器上监听的端口。当外部通过该端口访问公有云时，流量会被转发回家庭电脑的 22 端口（SSH 服务），从而实现从公网访问内网设备。

# q
建立 SSH 反向隧道后，如何通过公有云服务器实际访问家庭电脑的 SSH 服务？
# a
1. 在家庭电脑上执行 `ssh -R` 命令建立隧道。
2. SSH 登录到公有云服务器。
3. 在公有云服务器上通过 `ssh -p <remote_port> localhost` 即可连接到家庭电脑（例如 `ssh -p 3306 localhost`）。

