# q
如何进入一个正在运行的 Docker 容器并启动交互式 shell？
# a
使用 `docker exec -it <container_id> /bin/bash` 命令，该命令会在指定容器内启动一个交互式的 bash shell。

# q
如何将 Docker 容器内的文件复制到宿主机？
# a
使用 `docker cp` 命令，格式为：`docker cp <container_id>:/path/to/container/file /host/path/target`。该命令将容器内指定路径的文件复制到宿主机的目标路径。

