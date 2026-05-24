# q
Docker 容器中 SSH 登录在密钥交换后立即中断的常见原因是什么？
# a
原因是容器内通常没有完整的 `systemd` 或 PAM 初始化进程，导致 SSH 服务在尝试 PAM 认证时崩溃，连接在 `SSH2_MSG_KEX_ECDH_REPLY` 后意外终止。

# q
如何通过禁用 UsePAM 来解决容器内 SSH 闪退问题？
# a
1. 以 root 身份进入容器：  
```bash
docker exec -it -u root <容器名或ID> /bin/bash
```
2. 编辑 SSH 配置文件：  
```bash
vi /etc/ssh/sshd_config
```
3. 将 `UsePAM` 设置为 `no`（若无该行则手动添加）：  
```
UsePAM no
```
4. 重启 SSH 服务：  
```bash
service ssh restart
```

