# q
如何使用 `ssh-keygen` 生成 RSA 密钥对并指定私钥路径且不设置密码？
# a
使用以下命令：
```bash
/usr/bin/ssh-keygen -t rsa -f ~/.ssh/id_rsa -P ''
```
其中 `-t rsa` 指定密钥类型为 RSA，`-f ~/.ssh/id_rsa` 指定私钥文件输出路径，`-P ''` 表示生成的密钥不设置密码短语（空密码）。

# q
如何将生成的公钥复制到 `authorized_keys` 文件中以实现免密登录？
# a
直接使用 `cp` 命令复制：
```bash
cp ~/.ssh/id_rsa.pub ~/.ssh/authorized_keys
```
这将公钥内容写入授权密钥文件，使持有对应私钥的客户端可以免密登录本机。

# q
`ssh-keygen` 命令中 `-P ''` 参数的作用是什么？
# a
`-P ''` 表示为生成的私钥设置空密码，即不设置密码短语。这样在使用该私钥进行 SSH 认证时不需要额外输入密码，常用于自动化脚本或免交互登录场景。

