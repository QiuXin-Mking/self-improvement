# q
排查goofys挂载失败问题时应查看哪个日志文件？
# a
应查看 `/var/log/goofys.log`，因为容器启动脚本中将goofys的标准输出和错误都重定向到了该文件：
```
goofys -o allow_other --endpoint=$AWS_S3_ENDPOINT $BUCKET_NAME /mnt/goofys > /var/log/goofys.log 2>&1 &
```

# q
制作goofys容器镜像时，如何确保启动脚本具有可执行权限？
# a
在Dockerfile中生成 `/start_services.sh` 后，通过以下命令添加执行权限：
```
chmod +x /start_services.sh
```

# q
容器启动后goofys挂载S3桶的标准命令是什么？涉及哪些关键环境变量？
# a
标准命令为：
```
goofys -o allow_other --endpoint=$AWS_S3_ENDPOINT $BUCKET_NAME /mnt/goofys
```
依赖两个环境变量：`$AWS_S3_ENDPOINT`（S3兼容端点地址）和 `$BUCKET_NAME`（目标桶名称），挂载点固定为 `/mnt/goofys`。

