# q
如何配置SSH免密登录的核心步骤是什么？
# a
1. 在客户端生成密钥对：`ssh-keygen -t rsa -b 4096`
2. 将公钥上传至目标服务器：`ssh-copy-id 用户@目标IP`
`ssh-copy-id` 会自动完成公钥上传和 `.ssh` 目录、`authorized_keys` 文件的权限设置，是最安全便捷的方式。

# q
`ssh-copy-id` 命令的作用是什么？
# a
将本机的SSH公钥复制到远程服务器的 `~/.ssh/authorized_keys` 文件中，并自动设置正确的目录和文件权限，从而实现免密登录。

# q
如何为特定用户生成RSA密钥对？
# a
登录该用户后执行：
```
ssh-keygen -t rsa -b 4096
```
按提示操作（通常直接回车使用默认路径和无密码短语）即可在 `~/.ssh/` 下生成 `id_rsa`（私钥）和 `id_rsa.pub`（公钥）。

# q
如何配置从A服务器ubuntu用户到B服务器root用户的免密登录？
# a
1. 登录A服务器，切换到ubuntu用户
2. 生成密钥对（如尚未生成）：`ssh-keygen -t rsa -b 4096`
3. 执行：`ssh-copy-id root@B服务器IP`
完成后即可从A的ubuntu用户免密SSH到B的root用户。

