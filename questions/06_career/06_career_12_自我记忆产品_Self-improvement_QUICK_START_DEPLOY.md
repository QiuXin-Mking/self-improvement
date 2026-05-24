# q
快速部署本系统有哪三种方式？
# a
方式一：PyInstaller 打包（最简单，上传 deploy/ 目录执行）；方式二：Docker 部署（使用 Dockerfile 和 docker-compose.yml）；方式三：直接部署（使用 deploy.sh 或手动创建虚拟环境安装依赖）。

# q
使用 PyInstaller 打包后，在服务器上如何初始化并运行训练程序？
# a
```bash
cd /opt/spaced-repetition-train
chmod +x train
./train --init
./train
```

# q
使用 Docker 部署时，如何执行首次初始化并开始训练？
# a
```bash
docker-compose run --rm train --init
docker-compose run --rm train
```

# q
在直接部署方式中，推荐的自动部署命令是什么？
# a
```bash
chmod +x deploy.sh
sudo ./deploy.sh
```

# q
常用的运行命令有哪些？
# a
- 初始化知识库：`./train --init`
- 开始训练：`./train`
- 查看统计：`./train --stats`

