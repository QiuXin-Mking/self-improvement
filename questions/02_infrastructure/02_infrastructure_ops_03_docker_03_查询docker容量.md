# q
如何快速查看 Docker 的整体磁盘使用和剩余空间？
# a
使用 `docker system df` 命令。输出会以表格形式展示 Images、Containers、Local Volumes 等对象的磁盘情况，`SIZE` 列为总大小，`RECLAIMABLE` 列为可回收空间。

# q
`docker system df` 与 `docker system df -v` 的区别是什么？
# a
`docker system df` 提供汇总的磁盘使用摘要；`docker system df -v` 提供详细版本，会列出每个镜像和容器的具体磁盘占用情况。

# q
如何查看 Docker 卷的占用及特定容器内部的磁盘使用情况？
# a
- 卷占用：先用 `docker volume ls` 列出所有卷，再结合宿主机命令（如 `du`）查看其挂载点目录的大小。`docker volume ls` 本身不直接显示卷的剩余空间。
- 容器内部：先用 `docker ps` 找到容器名或 ID，然后执行 `docker exec -it <容器名或ID> /bin/bash` 进入容器，在容器内执行 `df -h` 查看磁盘使用。

