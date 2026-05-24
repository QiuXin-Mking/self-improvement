# q
如何列出本地主机上的所有 Docker 镜像？
# a
使用命令 `docker images`

# q
docker run 命令的作用是什么？
# a
创建一个新的容器并运行一个命令

# q
如何停止一个正在运行的容器？
# a
使用 `docker stop` 命令，可以停止一个或多个容器

# q
docker ps 和 docker images 的区别是什么？
# a
`docker ps` 列出当前运行中的容器，`docker images` 列出本地所有镜像

# q
如何进入一个正在运行的容器并执行命令？
# a
使用 `docker exec` 命令，例如 `docker exec -it <容器名或ID> bash`

