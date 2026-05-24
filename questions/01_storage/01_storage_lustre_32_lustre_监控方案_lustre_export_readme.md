# q
docker run 命令中的 `--privileged` 标志有什么作用？
# a
`--privileged` 赋予容器几乎所有的宿主机权限，包括访问硬件设备、修改内核参数等，通常用于需要执行系统级管理操作（如挂载文件系统、加载模块）的场景。

# q
docker run 命令中的 `-v` 参数用在什么场景？上述命令中 `/usr/sbin:/host-sbin` 的含义是什么？
# a
`-v` 用于挂载卷，将宿主机目录或文件映射到容器内部。
`/usr/sbin:/host-sbin` 表示将宿主机的 `/usr/sbin` 目录挂载到容器内的 `/host-sbin` 路径，从而在容器内可以访问宿主机的系统二进制文件。

# q
如何自定义 Docker 容器的入口点（entrypoint）？
# a
在 `docker run` 命令中使用 `--entrypoint` 参数，例如：
```bash
docker run -it --entrypoint /bin/bash <镜像名>
```
这会覆盖镜像默认的 ENTRYPOINT，启动容器后直接进入 Bash shell。

