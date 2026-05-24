# q
如何使用 ssh-keygen 命令生成 RSA 密钥对（含空密码）？
# a
使用以下命令生成 RSA 密钥对，并将私钥存入 `~/.ssh/id_rsa`、公钥存入 `~/.ssh/id_rsa.pub`，同时设置空密码：
```bash
/usr/bin/ssh-keygen -t rsa -f ~/.ssh/id_rsa -P ''
```
其中：
- `-t rsa` 指定密钥类型为 RSA
- `-f ~/.ssh/id_rsa` 指定密钥文件路径
- `-P ''` 表示私钥密码为空

# q
`ssh-keygen -R` 命令的作用和典型用法是什么？
# a
`ssh-keygen -R` 用于从 `~/.ssh/known_hosts` 文件中删除指定主机的旧 SSH 主机密钥条目，避免因主机密钥变更导致的警告或连接问题。典型用法：
```bash
ssh-keygen -R 172.22.101.189
ssh-keygen -R 172.22.101.190
ssh-keygen -R 172.22.101.191
```
在配置互信前，需在所有节点上执行该命令清理旧的主机密钥。

# q
在多节点（如 189/190/191）间配置 SSH 互信时，如何整合各节点的公钥并分发？
# a
1. 确保各节点已生成密钥对且 `~/.ssh` 目录存在（权限 700）。
2. 在任意一个节点（如 Node1 189）上执行以下命令，将三个节点的公钥追加到 `authorized_keys` 文件中：
```bash
ssh 172.22.101.189 cat ~/.ssh/id_rsa.pub >> ~/.ssh/authorized_keys
ssh 172.22.101.190 cat ~/.ssh/id_rsa.pub >> ~/.ssh/authorized_keys
ssh 172.22.101.191 cat ~/.ssh/id_rsa.pub >> ~/.ssh/authorized_keys
```
3. 设置 `authorized_keys` 文件权限为 600：
```bash
chmod 600 ~/.ssh/authorized_keys
```
4. 将整合后的 `authorized_keys` 文件通过 `scp` 分发给其他节点：
```bash
scp /root/.ssh/authorized_keys 172.22.101.190:~/.ssh
scp /root/.ssh/authorized_keys 172.22.101.191:~/.ssh
```

# q
配置完 SSH 互信后，如何验证免密登录是否生效？
# a
在所有参与互信的节点上分别执行以下命令，测试到每个目标节点的 SSH 连接：
```bash
ssh 172.22.101.189 date
ssh 172.22.101.190 date
ssh 172.22.101.191 date
```
如果执行后无需输入密码就能直接显示目标节点的系统日期，说明 SSH 互信配置成功。

