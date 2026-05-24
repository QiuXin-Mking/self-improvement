# q
如何通过 AWS CLI 获取指定存储桶的策略？
# a
使用 `aws s3api get-bucket-policy` 命令，并指定 `--bucket` 和 `--endpoint-url` 参数。例如：
```bash
aws s3api get-bucket-policy --bucket fe-assets --endpoint-url http://ees-southwest-10.edgeray.cn:5085
```

# q
如何使用 Python boto3 获取存储桶策略并处理异常？
# a
创建 S3 客户端时指定 `endpoint_url` 和签名版本 `s3v4`，调用 `get_bucket_policy` 方法，并用 `ClientError` 捕获异常。示例代码：
```python
import boto3
from botocore.config import Config
from botocore.exceptions import ClientError

def get_bucket_policy(bucket_name, endpoint_url):
    s3_client = boto3.client('s3', endpoint_url=endpoint_url, config=Config(signature_version='s3v4'))
    try:
        response = s3_client.get_bucket_policy(Bucket=bucket_name)
        policy = response['Policy']
        print('Bucket Policy:', policy)
    except ClientError as e:
        print('Error:', e)

if __name__ == '__main__':
    bucket_name = 'fe-assets'
    endpoint_url = 'http://ees-southwest-10.edgeray.cn:5085'
    get_bucket_policy(bucket_name, endpoint_url)
```

