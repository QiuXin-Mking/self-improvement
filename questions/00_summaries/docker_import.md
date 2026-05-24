# q
如何将正在运行的 Docker 容器导出为一个 tar 归档文件？
# a
使用 `docker export` 命令，将容器的文件系统导出到标准输出，可以重定向到文件。例如：
```bash
docker export 4fade8eafc46 > image.tar
```
其中 `4fade8eafc46` 是容器 ID（可通过 `docker ps` 获取），`image.tar` 是输出的 tar 文件名。

# q
`docker ps` 命令的输出中哪个字段可以用于 `docker export` 命令？
# a
`CONTAINER ID` 字段（如 `4fade8eafc46`），可直接作为 `docker export` 的参数指定要导出的容器。

