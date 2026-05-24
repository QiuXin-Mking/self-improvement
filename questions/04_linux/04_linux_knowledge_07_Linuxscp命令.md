# q
如何将本地目录递归传输到远程服务器的指定路径？
# a
使用 `scp -r` 命令，例如：
```bash
scp -r images/ 172.22.80.80:/root/install/
```
`-r` 表示递归复制整个目录及其内容。

# q
如何将 SSH 公钥文件传输到远程服务器的 `.ssh` 目录？
# a
使用 `scp -r` 命令，以下示例将 `authorized_keys` 文件复制到远程的 `/root/.ssh/` 下：
```bash
scp -r authorized_keys 172.22.251.104:/root/.ssh/
```

