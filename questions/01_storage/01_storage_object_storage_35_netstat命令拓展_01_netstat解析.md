# q
如何显示系统的路由表？
# a
使用 `-r` 选项
```bash
netstat -r
```

# q
netstat 命令的 `-i` 选项有什么作用？
# a
显示网络接口表（interface table），列出所有网络接口的统计信息
```bash
netstat -i
```

# q
如何让 netstat 以数字形式显示地址和端口，避免名称解析？
# a
使用 `-n`（`--numeric`）或更细粒度的选项，如 `--numeric-hosts`、`--numeric-ports`
```bash
netstat -n
```

# q
如何只显示处于监听状态的服务器套接字？
# a
使用 `-l` 或 `--listening` 选项
```bash
netstat -l
```

# q
如何查看每个套接字对应的进程 ID 和程序名称？
# a
使用 `-p`（`--programs`）选项
```bash
netstat -p
```

