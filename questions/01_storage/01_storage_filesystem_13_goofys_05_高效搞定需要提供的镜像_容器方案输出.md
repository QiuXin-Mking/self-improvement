# q
如何理解运行集成goofys的容器时必须添加`--device /dev/fuse`和`--cap-add SYS_ADMIN`参数？
# a
`--device /dev/fuse` 允许容器访问 FUSE 设备，goofys 依赖 FUSE 将 S3 存储桶挂载为本地文件系统。`--cap-add SYS_ADMIN` 授予容器管理挂载操作所需的系统权限。同时需要结合 `--user root` 和 `--security-opt apparmor=unconfined` 以禁用安全限制，确保挂载能够成功执行。

# q
运行 goofys 容器时需要配置哪些环境变量，各代表什么含义？
# a
- `AWS_ACCESS_KEY_ID` 和 `AWS_SECRET_ACCESS_KEY`：访问 S3 兼容存储的凭证。
- `AWS_S3_ENDPOINT`：S3 兼容服务的端点 URL，例如 `http://ees-southwest-8.edgeray.cn:5085`。
- `BUCKET_NAME`：要挂载的存储桶名称，例如 `goofys-test`。
- `S3_MOUNT`：容器内的挂载点路径，例如 `/mnt/storage`。

# q
构建和保存 goofys 集成容器镜像的基本命令是什么？
# a
构建镜像：
```bash
docker build -t <镜像名:标签> .
```
保存镜像为压缩包：
```bash
docker save -o <输出文件名>.tar.gz <镜像名:标签>
```
示例：
```bash
docker build -t goofys_pytorch:test2 .
docker save -o goofys_pytorch.tar.gz goofys_pytorch:test2
```

