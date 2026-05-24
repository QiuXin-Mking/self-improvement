# q
在 bash 脚本中启动 Windows Docker Desktop 时，如何处理路径中的空格？
# a
使用双引号将包含空格的完整路径括起来，例如：
```bash
"/drives/C/Program Files/Docker/Docker/Docker Desktop.exe" &
```

# q
在脚本中如何等待 Docker 引擎完全就绪后再执行后续命令？
# a
使用 `while` 循环检测 `docker info` 命令是否成功，例如：
```bash
while ! docker info > /dev/null 2>&1; do
    echo "Docker 尚未就绪，等待 3 秒..."
    sleep 3
done
```

# q
如何通过 SSH 连接远程服务器并指定非默认端口？
# a
使用 `ssh 用户名@IP地址 -p 端口号`，例如：
```bash
ssh root@115.190.235.149 -p 5555
```

# q
启动已存在的 Docker 容器的命令是什么？
# a
```bash
docker start 容器名
```
例如 `docker start vue3-dev`。

# q
如何给 shell 脚本添加可执行权限？
# a
使用 `chmod +x 脚本文件名`，例如：
```bash
chmod +x start_vue_dev.sh
```

