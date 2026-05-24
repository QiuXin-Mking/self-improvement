# q
如何在容器启动脚本中为Jupyter Lab添加自定义base_url和token支持？
# a
在启动脚本中设置环境变量和参数：
```bash
DEFAULT_NOTEBOOK_TOKEN="3bv8uoqs855ukot7kr3c0s9nawj5b908srkb7tur89r248sn"
ENV_NOTEBOOK_TOKEN="${NOTEBOOK_TOKEN:-$DEFAULT_NOTEBOOK_TOKEN}"

DEFAULT_NOTEBOOK_BASE_PATH="/"
ENV_NOTEBOOK_BASE_PATH="${NOTEBOOK_BASE_PATH:-$DEFAULT_NOTEBOOK_BASE_PATH}"

jupyter lab --ip=0.0.0.0 --port=8888 --allow-root --no-browser \
    --NotebookApp.token="$ENV_NOTEBOOK_TOKEN" \
    --NotebookApp.base_url="$ENV_NOTEBOOK_BASE_PATH" > /var/log/jupyter.log 2>&1 &
```
通过环境变量 `ENV_NOTEBOOK_BASE_PATH` 控制 base_url，token 可通过 `NOTEBOOK_TOKEN` 外部传入，缺省使用默认值。

# q
容器镜像需要新增TensorBoard服务时，标准的修改流程是什么？
# a
在容器启动脚本（如 `/start_services.sh`）中添加TensorBoard启动命令，确保日志目录存在：
```bash
tensorboard --logdir=/var/log/tensorboard --host=0.0.0.0 --port=6006 > /var/log/tensorboard.log 2>&1 &
```
之后重新构建镜像或替换现有容器的启动脚本并重启容器。若需持久化日志，可将 `/var/log/tensorboard` 挂载到宿主机或存储卷。

# q
如何通过Goofys让容器启动时自动挂载S3兼容存储？
# a
在启动脚本中加入以下命令：
```bash
mkdir -p $S3_MOUNT
goofys -f -o allow_other --endpoint=$AWS_S3_ENDPOINT $BUCKET_NAME $S3_MOUNT > /var/log/goofys.log 2>&1 &
```
其中 `$S3_MOUNT` 是挂载点路径，`$AWS_S3_ENDPOINT` 为S3服务地址，`$BUCKET_NAME` 为存储桶名。这些变量需通过环境变量传入或定义为常量。Goofys 以后台方式运行，日志重定向到 `/var/log/goofys.log`。

