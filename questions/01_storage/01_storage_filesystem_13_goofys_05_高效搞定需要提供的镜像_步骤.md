# q
在定制用于Goofys的Docker镜像时，需要解决哪些基础环境问题？
# a
- 设置中文编码为 `UTF-8`，避免中文异常。  
- 将 pip 默认源更换为国内源。  
- 创建缺失的 `/var/log/tensorboard` 目录。  
- 安装 `vim`（必要时使用 `-f` 强制安装）。  
- 安装必需工具：`iproute2`、`net-tools`、`telnet`、`curl`、`vim`。  
- 确认 goofys 可执行文件路径正确。  
- 基础镜像可选 `ollama.tar` 和 `pytorch-2.1.2-py3.10-cuda11.8-u20.04.tar.gz`。

# q
运行Goofys容器挂载S3桶时，需要传递哪些关键环境变量和设备映射？
# a
必须设置以下环境变量：
- `AWS_ACCESS_KEY_ID`、`AWS_SECRET_ACCESS_KEY`：S3 访问凭证。  
- `AWS_S3_ENDPOINT`：S3 兼容存储端点（如 `http://ees-southwest-8.edgeray.cn:5085`）。  
- `BUCKET_NAME`：要挂载的存储桶名称（如 `goofys-test`）。  
- `S3_MOUNT`：容器内挂载点（如 `/mnt/storage` 或 `/mnt/goofys1`）。  

设备与权限映射：
- 挂载 `/dev/fuse` 设备（`--device /dev/fuse`）。  
- 添加 `SYS_ADMIN` 能力（`--cap-add SYS_ADMIN`）。  
- 网络使用 `host` 模式（`--network host`）。  
- 以 root 用户运行（`--user root`）。  
- 关闭 AppArmor 限制（`--security-opt apparmor=unconfined`）。

