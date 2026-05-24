# q
docker命令终端有，但sh里面不能用，如何检查？
# a
使用 `command -v docker` 检查docker命令是否在PATH中，以及它是什么类型（别名、函数、二进制文件等）。

# q
ssh-copy-id 怎么使用？
# a
基本用法：
```bash
ssh-copy-id [远程服务器用户名]@[远程服务器IP]
```
该命令会将本地SSH公钥自动复制到远程服务器的 `~/.ssh/authorized_keys` 中，实现免密登录。

