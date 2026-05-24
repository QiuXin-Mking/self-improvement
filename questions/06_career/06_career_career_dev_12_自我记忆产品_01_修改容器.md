# q
如何停止并删除一个名为 vue3-dev 的 Docker 容器？
# a
```bash
docker stop vue3-dev
docker rm vue3-dev
```

# q
在使用 `docker run` 创建新容器时，如何将本地目录 `D:\qiuxin_aliyun_back\04_职场修炼\12_自我记忆产品\Self-improvement` 挂载到容器内的 `/Self-improvement`？
# a
在 `docker run` 命令中添加 `-v` 参数：
```bash
-v "D:\qiuxin_aliyun_back\04_职场修炼\12_自我记忆产品\Self-improvement:/Self-improvement"
```

# q
如何将一个已存在的 Docker 容器保存为新镜像，并在新容器中使用该镜像？
# a
1. 将容器提交为新镜像：
   ```bash
   docker commit vue3-dev vue3-dev-with-deps
   ```
2. 后续使用该镜像启动新容器：
   ```bash
   docker run -d --name vue3-dev-new ... vue3-dev-with-deps tail -f /dev/null
   ```

# q
在 Docker 中，如何同时映射多个端口（例如 5173、5000、3000）并保持容器运行？
# a
在 `docker run` 命令中使用多个 `-p` 参数分别映射，并使用 `tail -f /dev/null` 作为保持前台运行的阻塞命令：
```bash
docker run -d \
  --name vue3-dev-new \
  -p 5173:5173 \
  -p 5000:5000 \
  -p 3000:3000 \
  ... \
  vue3-dev-frp \
  tail -f /dev/null
```

# q
如何在 Docker Run 命令中设置容器的工作目录？
# a
使用 `-w` 参数指定工作目录，例如：
```bash
-w /app
```

