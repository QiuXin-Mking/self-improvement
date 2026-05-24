# q
如何查看当前安装的 WSL 版本？
# a
使用以下命令可以查看 WSL 的版本信息：
```bash
wsl --version
```

# q
如何列出所有已安装的 WSL 发行版及其状态？
# a
使用以下命令可以列出所有 WSL 发行版及其运行状态（-l 列出，-v 详细信息）：
```bash
wsl -l -v
```

# q
如何进入一个特定的 WSL 子系统（例如 Ubuntu）？
# a
使用 `-d` 参数指定发行版名称即可进入：
```bash
wsl -d Ubuntu
```

# q
如何将某个 WSL 发行版设置为默认子系统？
# a
使用 `--set-default` 参数指定发行版名称：
```bash
wsl --set-default Ubuntu
```

