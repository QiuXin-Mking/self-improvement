# q
这个脚本的核心功能是什么？
# a
实现多台服务器间的 SSH 互信（免密登录），并统一设置各服务器的主机名。

# q
脚本中生成 SSH 密钥的命令是什么？
# a
```sh
ssh-keygen -t rsa -N "" -f ~/.ssh/id_rsa
```
该命令会在不设置密码（-N ""）的情况下生成RSA密钥对，私钥保存在 `~/.ssh/id_rsa`。

# q
脚本中如何将本机公钥分发到远程服务器？
# a
使用 `ssh-copy-id -i ~/.ssh/id_rsa.pub "$SERVER"` 命令，将指定公钥追加到远程服务器的 `~/.ssh/authorized_keys` 文件中。

# q
脚本中最终设置主机名用的是什么命令？
# a
```sh
ssh "$ser" hostnamectl set-hostname $ser
```
通过 SSH 远程执行 `hostnamectl set-hostname` 将主机名修改为对应的服务器名。

