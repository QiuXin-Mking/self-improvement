# q
Goofys 在 Docker 中挂载时权限过大导致无法实现租户隔离的典型根因是什么？
# a
典型根因是容器启动时使用了 `--privileged` 参数，赋予容器几乎所有内核能力（包括访问 `/dev/fuse` 的完全权限），这会绕过所有权限检查，使得 FUSE 挂载点对其他容器或进程可见并可直接访问，从而无法实现租户隔离。此外，即便不使用 `--privileged`，若添加了过多的 `--cap-add`（如 `SYS_ADMIN`、`DAC_READ_SEARCH`、`SYS_CHROOT`、`CAP_SYS_ADMIN` 全部叠加），也会造成权限过大。

# q
如何为 Goofys 容器配置最小权限以实现 FUSE 挂载？
# a
最小权限配置需要：
- 使用 `--device /dev/fuse` 将宿主机 FUSE 设备暴露给容器
- 添加必要的内核能力：`--cap-add SYS_ADMIN`（可酌情保留 `DAC_READ_SEARCH` 和 `SYS_CHROOT`，但避免 `CAP_SYS_ADMIN` 冗余叠加）
- 避免使用 `--privileged`
- 容器内挂载时使用 `-o allow_other` 选项，并确保宿主机 `/etc/fuse.conf` 已开启 `user_allow_other`

示例启动命令：
```
docker run \
  -e AWS_ACCESS_KEY_ID=... \
  -e AWS_SECRET_ACCESS_KEY=... \
  -e AWS_S3_ENDPOINT=http://ees-southwest-8.edgeray.cn:5085 \
  -e BUCKET_NAME=goofys-test \
  -e S3_MOUNT=/mnt/goofys1 \
  --device /dev/fuse \
  --cap-add SYS_ADMIN --cap-add DAC_READ_SEARCH --cap-add SYS_CHROOT \
  --network host \
  --user root \
  --name=goofys_test \
  -d goofys:test11
```

# q
启动 Goofys 挂载时需要在容器内执行哪些关键命令？
# a
首先在容器宿主机或容器内配置 FUSE 允许其他用户访问：
```
echo "user_allow_other" >> /etc/fuse.conf
```

然后在容器内挂载 S3 存储桶，推荐使用前台调试模式进行验证：
```
goofys -f --debug_fuse -o allow_other --endpoint=$AWS_S3_ENDPOINT $BUCKET_NAME /mnt/goofys
```

若需后台运行，去掉 `-f` 即可：
```
goofys -o allow_other --endpoint=$AWS_S3_ENDPOINT $BUCKET_NAME /mnt/goofys
```

若需更详细的 S3 层面调试，可添加 `--debug_s3` 并使用 `strace` 跟踪系统调用：
```
strace goofys -f --debug_s3 --debug_fuse -o allow_other --endpoint=$AWS_S3_ENDPOINT $BUCKET_NAME /mnt/goofys
```

