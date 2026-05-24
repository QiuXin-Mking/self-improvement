# q
如何使用 Docker 部署思源笔记服务？
# a
运行以下命令启动容器，将宿主机端口 6806 映射到容器 6806，并挂载本地工作目录：
```sh
docker run -d -p 6806:6806 \
  --name siyuan \
  -v /your/local/path:/siyuan/workspace \
  b3log/siyuan
```
如果部署在 Docker 中，必须通过环境变量 `SIYUAN_ACCESS_AUTH_CODE` 设置访问授权码，否则容器会启动失败。

# q
思源笔记 Docker 部署后无法访问，有哪些调试手段？
# a
可使用以下命令进行排查：
```sh
# 从远程测试端口通断
curl -v http://<服务器IP>:6806/

# 在宿主机本地测试
curl -v http://127.0.0.1:6806

# 查看容器日志
docker logs siyuan

# 检查端口监听
netstat -ntlp | grep 6806
```

# q
部署思源笔记 Docker 容器时，日志提示 “the access authorization code command line parameter (--accessAuthCode) must be set when deploying via Docker” 如何解决？
# a
必须在启动容器时通过环境变量设置访问授权码。示例命令：
```sh
docker run -d -p 6806:6806 \
  --name siyuan \
  -v /your/local/path:/siyuan/workspace \
  -e SIYUAN_ACCESS_AUTH_CODE=your_code_here \
  b3log/siyuan
```
如果之前已创建了未设置授权码的容器，需要先删除再重新启动：
```sh
docker rm siyuan
```

