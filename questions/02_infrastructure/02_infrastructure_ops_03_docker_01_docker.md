# q
docker ps 命令的作用是什么？
# a
列出所有正在运行的 Docker 容器，显示容器 ID、镜像、创建时间、状态、端口映射、名称等详细信息。

# q
如何仅获取所有正在运行的 Docker 容器的 ID？
# a
使用命令 `docker ps -q` 仅返回容器 ID，添加 `-a` 则会包括已停止的容器：`docker ps -aq`。

# q
如何查看包括已停止的 Docker 容器的完整信息？
# a
使用 `docker ps -a` 列出所有容器（运行中和已停止），并显示完整详情。

