# q
Linux下如何通过Docker运行MaxKB？
# a
执行以下命令：
```bash
docker run -d --name=maxkb --restart=always -p 8080:8080 -v ~/.maxkb:/var/lib/postgresql/data -v ~/.python-packages:/opt/maxkb/app/sandbox/python-packages registry.fit2cloud.com/maxkb/maxkb
```

# q
MaxKB安装后的默认登录凭据是什么？
# a
默认用户名：`admin`，默认密码：`MaxKB@123..`

# q
安装MaxKB时遇到`Error: image maxkb/maxkb:latest not found.`该如何处理？
# a
确认使用的镜像地址完整且正确，应使用`registry.fit2cloud.com/maxkb/maxkb`而不是简写`maxkb/maxkb`。如果问题仍存在，可先手动拉取镜像：`docker pull registry.fit2cloud.com/maxkb/maxkb`，再执行运行命令。

