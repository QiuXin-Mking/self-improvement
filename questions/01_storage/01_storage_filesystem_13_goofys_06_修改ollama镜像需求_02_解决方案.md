# q
运行Goofys容器以实现S3挂载，需要哪些关键的Docker参数和权限配置？
# a
关键配置包括：
- 设备映射：`--device /dev/fuse` 和权限提升 `--cap-add SYS_ADMIN`，以支持FUSE文件系统。
- 网络模式：`--network host`（使用宿主机网络）。
- 用户模式：`--user root` 和 `--security-opt apparmor=unconfined` 关闭AppArmor限制。
- 环境变量：需指定 `AWS_ACCESS_KEY_ID`、`AWS_SECRET_ACCESS_KEY`、`AWS_S3_ENDPOINT`、`BUCKET_NAME` 和 `S3_MOUNT` 挂载点。
示例命令：
```bash
docker run -d \
  -e AWS_ACCESS_KEY_ID=xxx \
  -e AWS_SECRET_ACCESS_KEY=xxx \
  -e AWS_S3_ENDPOINT=http://endpoint:port \
  -e BUCKET_NAME=goofys-test \
  -e S3_MOUNT=/mnt/storage \
  --device /dev/fuse \
  --cap-add SYS_ADMIN \
  --network host \
  --user root \
  --security-opt apparmor=unconfined \
  --name container_name \
  image:tag
```

# q
如何配置Ollama服务以监听所有网络接口并允许跨域请求？
# a
在启动Ollama服务前，设置以下环境变量：
```bash
export OLLAMA_HOST="0.0.0.0:11434"
export OLLAMA_ORIGINS=*
```
然后将服务进程后台启动并记录日志：
```bash
ollama serve > /var/log/ollama.log 2>&1 &
```

# q
如何将运行中的容器修改保存为新镜像？
# a
使用 `docker commit` 命令，语法为：
```bash
docker commit -m "提交信息" -a "作者" 容器ID 新镜像名:标签
```
例如将容器 `ed4c4edfc6a9` 保存为 `goofys_ollama:test4`：
```bash
docker commit -m "add ollama" -a "qiuxin" ed4c4edfc6a9 goofys_ollama:test4
```

