# q
如何使用 ssh-keygen 生成 RSA 密钥对？
# a
执行命令：
```
ssh-keygen -t rsa -b 4096 -C "your_email@example.com"
```
参数说明：`-t rsa` 指定密钥类型为 RSA，`-b 4096` 指定密钥长度 4096 位，`-C` 添加注释（通常是邮箱）。

# q
如何将本地公钥复制到远程服务器实现 SSH 免密登录？
# a
使用 `ssh-copy-id` 命令：
```
ssh-copy-id user@remote_host
```
例如：
```
ssh-copy-id root@121.41.87.58
```
该命令会将本地 `~/.ssh/id_rsa.pub` 的内容追加到远程服务器的 `~/.ssh/authorized_keys` 文件中。

# q
SSH config 文件中如何配置别名和密钥路径？
# a
在 `~/.ssh/config` 中添加如下格式的 Host 块：
```
Host myserver
    HostName remote_host
    User user
    IdentityFile ~/.ssh/id_rsa
```
之后可以使用 `ssh myserver` 代替完整的 `ssh user@remote_host -i ~/.ssh/id_rsa`。

