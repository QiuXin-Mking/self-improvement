# q
如何使用最少的 cap-add 权限启动 goofys 容器挂载 S3 存储桶？
# a
使用以下命令，仅添加 `SYS_ADMIN` 能力：
```bash
docker run \
-e AWS_ACCESS_KEY_ID=10QOENRYZLOIJF1ME6OQ  \
-e AWS_SECRET_ACCESS_KEY=MTFf02dPaJ63kA0yCW01EdOMFtPuokiqAyDzclUV  \
-e AWS_S3_ENDPOINT=http://ees-southwest-8.edgeray.cn:5085 -e BUCKET_NAME=goofys-test -e S3_MOUNT=/mnt/goofys1 \
--device /dev/fuse \
--cap-add SYS_ADMIN  \
--network host \
--user root \
--security-opt apparmor=unconfined \
--name=goofys_test1 \
-d goofys:test11
```

# q
如何在特权模式下启动 goofys 容器以进行调试？
# a
使用 `--privileged` 标志运行容器：
```bash
docker run \
-e AWS_ACCESS_KEY_ID=10QOENRYZLOIJF1ME6OQ  \
-e AWS_SECRET_ACCESS_KEY=MTFf02dPaJ63kA0yCW01EdOMFtPuokiqAyDzclUV  \
-e AWS_S3_ENDPOINT=http://ees-southwest-8.edgeray.cn:5085 -e BUCKET_NAME=goofys-test -e S3_MOUNT=/mnt/goofys1 \
--privileged \
--name=goofys_test5_test \
-d goofys_ollama:test5
```

