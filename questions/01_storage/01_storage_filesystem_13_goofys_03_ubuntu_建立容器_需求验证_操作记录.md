# q
在 start_services.sh 中,如何通过环境变量优雅地配置 Jupyter 的访问 token？
# a
使用 Shell 参数扩展语法设置默认值：
```bash
DEFAULT_NOTEBOOK_TOKEN="3bv8uoqs855ukot7kr3c0s9nawj5b908srkb7tur89r248sn"
ENV_NOTEBOOK_TOKEN="${NOTEBOOK_TOKEN:-$DEFAULT_NOTEBOOK_TOKEN}"
```
如果环境变量 `NOTEBOOK_TOKEN` 未设置或为空，则使用 `DEFAULT_NOTEBOOK_TOKEN`；否则优先使用外部传入的值。启动 Jupyter 时传入该变量：
```bash
jupyter lab ... --NotebookApp.token="$ENV_NOTEBOOK_TOKEN" ...
```

# q
如何通过 Dockerfile 快速构建一个自动启动多个服务的容器镜像？
# a
基于基础镜像 `goofys:formal8`，复制新的 `start_services.sh` 脚本并赋予执行权限，最后指定容器启动命令：
```dockerfile
FROM goofys:formal8
COPY ./start_services.sh /start_services.sh
RUN chmod +x /start_services.sh
CMD ["/start_services.sh"]
```
构建命令：
```
docker build -t goofys:formal9 .
```

# q
goofys 容器化部署时，挂载部分需要哪些关键环境变量和命令？
# a
需要 `AWS_S3_ENDPOINT`、`BUCKET_NAME` 和 `S3_MOUNT` 环境变量。脚本中先创建挂载目录，然后执行：
```bash
mkdir -p $S3_MOUNT
goofys -o allow_other --endpoint=$AWS_S3_ENDPOINT $BUCKET_NAME $S3_MOUNT > /var/log/goofys.log 2>&1 &
```
参数 `-o allow_other` 允许其他用户访问挂载点，`--endpoint` 指定兼容 S3 的对象存储地址。

# q
以分离模式运行 goofys 容器时，如何通过 `docker run` 传入全部所需环境变量？
# a
使用多个 `-e` 参数，配合 `--privileged` 和 `--network host` 以获取完整系统权限和主机网络：
```bash
docker run -d \
  -e AWS_ACCESS_KEY_ID=10QOENRYZLOIJF1ME6OQ \
  -e AWS_SECRET_ACCESS_KEY=MTFf02dPaJ63kA0yCW01EdOMFtPuokiqAyDzclUV \
  -e AWS_S3_ENDPOINT=http://ees-southwest-8.edgeray.cn:5085 \
  -e BUCKET_NAME=goofys-test \
  -e S3_MOUNT=/mnt/goofys1 \
  -e NOTEBOOK_TOKEN="11" \
  --name=test3 \
  --privileged \
  --network host \
  goofys:formal9
```

