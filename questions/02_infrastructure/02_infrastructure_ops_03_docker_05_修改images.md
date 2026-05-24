# q
如何将运行中的容器修改保存为新的Docker镜像？
# a
使用 `docker commit` 命令。需要提供容器ID或名称，以及新镜像的名称和标签。
```
docker commit <container_id_or_name> <new_image_name>:<new_tag>
```
例如：`docker commit mycontainer myimage:modified`

# q
如何将本地Docker镜像导出为压缩文件，以及如何从该文件导入镜像？
# a
**导出**：使用 `docker save -o <path_to_tar_file> <image_name>:<tag>`
例如：`docker save -o myimage.tar myimage:latest`

**导入**：使用 `docker load -i <path_to_tar_file>`
例如：`docker load -i myimage.tar`

