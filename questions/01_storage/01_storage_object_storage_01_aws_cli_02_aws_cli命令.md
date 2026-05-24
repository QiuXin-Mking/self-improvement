# q
如何通过AWS CLI列出所有桶（使用自定义endpoint）？
# a
使用 `aws s3api list-buckets` 命令并指定 `--endpoint-url`。完整命令：
```bash
aws s3api --endpoint-url=http://www.qiuxin.com:5085 list-buckets
```

# q
如何通过AWS CLI创建一个桶（例如 `qiuxinbucket`，使用自定义endpoint）？
# a
使用 `aws s3api create-bucket` 命令，指定桶名和 endpoint：
```bash
aws s3api --endpoint-url=http://www.qiuxin.com:5085 create-bucket --bucket qiuxinbucket
```

# q
如何通过AWS CLI上传本地文件到桶中（使用自定义endpoint）？
# a
可以使用 `aws s3api put-object` 或 `aws s3 cp`。例如上传 `/root/qiuxin.log` 为 `q1`：
```bash
aws s3api put-object --endpoint-url=http://www.qiuxin.com:5085 --bucket qiuxinbucket --key q1 --body /root/qiuxin.log
```
或使用高级命令：
```bash
aws s3 cp --endpoint-url=http://www.qiuxin.com:5085 ./qiuxin_mount/qiuxin1 s3://qiuxinbucket/q1
```

# q
如何使用AWS CLI从桶中下载对象到本地（使用自定义endpoint）？
# a
使用 `aws s3api get-object` 或 `aws s3 cp`。例如下载 `q1` 到本地文件 `outfiles`：
```bash
aws s3api get-object --bucket qiuxinbucket --endpoint-url=http://www.qiuxin.com:5085 --key q1 outfiles
```
或者：
```bash
aws s3 cp --endpoint-url=http://www.qiuxin.com:5085 s3://qiuxinbucket/q1 ./qiuxin_mount/qiuxin1
```

# q
如何为S3对象生成一个有时效的预签名下载URL（使用自定义endpoint）？
# a
使用 `aws s3 presign` 命令，指定对象路径和过期时间。例如为 `s3://qiuxinbucket/qx2` 生成有效期3600秒的URL：
```bash
aws s3 presign --endpoint-url=http://www.qiuxin.com:5085 s3://qiuxinbucket/qx2 --expires-in 3600
```

