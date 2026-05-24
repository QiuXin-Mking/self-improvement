# q
如何使用 boto3 客户端连接到自定义的 S3 兼容对象存储？
# a
初始化时需指定 `aws_access_key_id`、`aws_secret_access_key` 和 `endpoint_url`，例如：
```python
s3 = boto3.client(
    's3',
    aws_access_key_id='your_access_key',
    aws_secret_access_key='your_secret_key',
    endpoint_url='http://127.0.0.1:5085'
)
```

# q
检查 S3 存储桶是否存在并自动创建的核心逻辑是什么？
# a
先调用 `s3.head_bucket(Bucket=bucket_name)` 判断桶是否存在；若抛出 `ClientError`，则调用 `s3.create_bucket(Bucket=bucket_name)` 进行创建。

# q
boto3 上传文件到 S3 的方法和关键参数是什么？
# a
使用 `s3.upload_file(file_name, bucket_name, object_name)`，参数依次为本地文件路径、目标桶名和对象名。

# q
如何在 boto3 中获取 S3 对象的元数据而不下载文件内容？
# a
调用 `s3.head_object(Bucket=bucket_name, Key=object_name)`，返回的响应中包含对象的元数据。

