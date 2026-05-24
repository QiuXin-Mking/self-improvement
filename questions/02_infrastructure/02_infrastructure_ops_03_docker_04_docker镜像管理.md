# q
如何配置Docker以使用私有的Harbor镜像仓库？
# a
修改Docker配置文件（如`/etc/docker/daemon.json`），在`insecure-registries`中添加Harbor服务器地址（例如`172.16.155.12:8091`），并可同时配置`registry-mirrors`。保存后执行以下命令重启Docker服务：
```bash
sudo systemctl restart docker
```

# q
登录Harbor仓库的命令是什么？
# a
使用`docker login`命令并指定仓库地址，例如：
```bash
docker login 172.16.155.12:8091
```
然后根据提示输入用户名和密码。

# q
如何将本地镜像标记并推送到Harbor仓库？
# a
首先使用`docker tag`将本地镜像重新标记为目标仓库的路径和标签，例如：
```bash
docker tag myapp:latest 172.16.155.12:8091/library/myapp:v1.0
```
然后推送该镜像：
```bash
docker push 172.16.155.12:8091/library/myapp:v1.0
```

# q
从Harbor拉取镜像的命令格式是什么？
# a
命令格式为：
```bash
docker pull harbor.mycompany.com/library/your-image:tag
```
其中`harbor.mycompany.com`、项目路径`library`、镜像名和标签需根据实际情况替换。

