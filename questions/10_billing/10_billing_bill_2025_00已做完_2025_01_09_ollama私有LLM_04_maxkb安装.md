# q
在 Linux 上如何通过在线方式一键部署 MaxKB？
# a
```bash
yum install -y docker
systemctl start docker
docker run -d --name=maxkb --restart=always -p 8080:8080 \
  -v ~/.maxkb:/var/lib/postgresql/data \
  -v ~/.python-packages:/opt/maxkb/app/sandbox/python-packages \
  registry.fit2cloud.com/maxkb/maxkb
```
该命令会自动拉取镜像并启动容器，数据卷挂载在主目录下。

# q
`docker run` 拉取 MaxKB 镜像时报错 `Error: image maxkb/maxkb:latest not found`，原因和解决办法是什么？
# a
原因：使用了错误的镜像名 `maxkb/maxkb`。  
正确镜像地址是 `registry.fit2cloud.com/maxkb/maxkb`，Docker 默认拉取 Docker Hub，而该镜像位于 fit2cloud 的私有 registry 中，必须使用完整地址。

# q
MaxKB 安装完成后，默认的访问地址和管理员登录凭据是什么？
# a
- 访问地址：`http://目标服务器IP:8080`  
- 用户名：`admin`  
- 默认密码：`MaxKB@123..`

# q
在 Windows 上安装 MaxKB 需要预先完成什么前提条件？
# a
必须预先解决 Docker 环境问题（如安装 Docker Desktop 并确保可用），然后参考官方离线安装文档：https://maxkb.cn/docs/installation/offline_installtion/ 进行部署。

