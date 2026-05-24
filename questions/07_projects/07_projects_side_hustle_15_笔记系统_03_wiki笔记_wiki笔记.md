# q
如何使用 Docker 快速启动 Wiki.js 实例？
# a
使用以下命令：
```bash
docker run -d -p 3000:3000 \
  --name wikijs \
  -e "DB_TYPE=sqlite" \
  -v /your/data/path:/wiki/data \
  requarks/wiki:latest
```
- `-d` 后台运行
- `-p 3000:3000` 映射容器3000端口到宿主机3000端口
- `--name wikijs` 设置容器名称为 wikijs
- `-e "DB_TYPE=sqlite"` 指定使用 SQLite 作为数据库
- `-v /your/data/path:/wiki/data` 将本地 `/your/data/path` 挂载到容器的 `/wiki/data`，用于持久化数据

# q
Wiki.js 容器中 `DB_TYPE=sqlite` 环境变量的作用是什么？
# a
指定 Wiki.js 使用 SQLite 作为数据库后端，无需额外配置外部数据库，适合快速搭建和测试。

