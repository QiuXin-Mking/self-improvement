# q
如何使用 Ollama 运行、删除和列出本地大模型？
# a
- 运行模型：`ollama run deepseek-r1:70b`
- 删除模型：`ollama rm deepseek-r1:70b`
- 列出已安装模型：`ollama list`

# q
如何通过 Docker 安装并启动 MaxKB，其默认登录凭据是什么？
# a
安装步骤：
1. 拉取镜像：`docker pull cr2.fit2cloud.com/1panel/maxkb`
2. 保存镜像：`docker save cr2.fit2cloud.com/1panel/maxkb -o maxkb.tar`
3. 加载镜像：`docker load < maxkb.tar`
4. 运行容器：`docker run -d --name=maxkb -p 8080:8080 -v ~/.maxkb:/var/lib/postgresql/data cr2.fit2cloud.com/1panel/maxkb`

访问地址：`http://localhost:8080`
默认用户名/密码：`admin` / `MaxKB@123..`

# q
如何解决 MaxKB 容器内部无法直接访问宿主机 Ollama 服务（localhost:11434）的问题？
# a
在容器内使用 `host.docker.internal` 代替 localhost，访问地址为 `http://host.docker.internal:11434`。

# q
如何给 Docker 镜像打标签（tag）？
# a
使用 `docker tag <IMAGE_ID> <USERNAME>/<REPOSITORY>:<TAG>`，例如：
```
docker tag b8c0977164f5 cr2.fit2cloud.com/1panel/maxkb:latest
```

