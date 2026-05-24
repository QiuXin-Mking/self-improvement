# q
如何在本地运行 DeepSeek R1 70B 模型？
# a
使用 Ollama 命令：`ollama run deepseek-r1:70b`

# q
如何删除本地已下载的 DeepSeek R1 70B 模型？
# a
使用命令：`ollama rm deepseek-r1:70b`

# q
如何从 Docker Hub 拉取并运行 MaxKB 容器，使其可以通过宿主机的 8080 端口访问？
# a
先拉取镜像：`docker pull cr2.fit2cloud.com/1panel/maxkb`，然后运行容器：
```
docker run -d --name=maxkb -p 8080:8080 -v ~/.maxkb:/var/lib/postgresql/data cr2.fit2cloud.com/1panel/maxkb
```

# q
MaxKB 的默认管理员用户名和密码是什么？
# a
默认用户名：`admin`，默认密码：`MaxKB@123..123` 或 `MaxKB@123..`（文档记载了两个可能值）。

# q
当 MaxKB 部署在容器内，需要访问宿主机的 Ollama 服务（端口 11434）时，如何解决容器内无法直接访问 localhost 的问题？
# a
使用 `host.docker.internal` 替代 `localhost`，即访问地址为 `http://host.docker.internal:11434`。

