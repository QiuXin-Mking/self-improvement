# q
在Windows的Bash脚本中启动Docker Desktop时，为什么要用引号括起可执行文件路径？
# a
因为路径中包含空格（例如 `Program Files`），如果不加引号，Shell会将其解析为多个参数，导致错误。正确写法：
```bash
"/drives/C/Program Files/Docker/Docker/Docker Desktop.exe" &
```

# q
如何编写一个等待Docker引擎完全就绪的Shell循环？
# a
使用 `while ! docker info > /dev/null 2>&1; do ... done` 循环，检测 `docker info` 命令是否成功执行，直到返回成功才继续。

# q
在Shell脚本中如何启动一个已存在的Docker容器？
# a
使用 `docker start` 命令，后跟容器名称或ID。例如：
```bash
docker start vue3-dev
```

# q
从本地通过MobaXterm使用SSH连接远程服务器的命令格式是什么？
# a
```bash
ssh root@115.190.235.149 -p 5555
```
其中 `root` 是用户名，`115.190.235.149` 是远程主机IP，`-p 5555` 指定非默认的SSH端口5555。

