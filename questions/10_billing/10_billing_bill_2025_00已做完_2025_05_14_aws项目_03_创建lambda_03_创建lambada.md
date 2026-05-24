# q
如何使用 AWS CLI 创建一个 Python 3.9 的 Lambda 函数？
# a
```bash
aws lambda create-function \
  --function-name my-test-lambda \
  --runtime python3.9 \
  --role arn:aws:iam::account-id:role/lambda-s3-role \
  --handler lambda_function.lambda_handler \
  --zip-file fileb://function.zip
```
前提：已将 Lambda 代码打包为 `function.zip`（入口文件在根目录），并且指定的 IAM 角色已存在且具备 Lambda 及相关服务的必要权限。

# q
为 Lambda 创建执行角色时，信任策略应如何编写？
# a
信任策略需允许 Lambda 服务代入该角色，典型 JSON 如下：
```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Action": "sts:AssumeRole"
    }
  ]
}
```
使用 CLI 创建角色：
```bash
aws iam create-role \
  --role-name lambda-s3-role \
  --assume-role-policy-document file://trust-policy.json
```

# q
如何通过 CLI 为 Lambda 添加 S3 触发器并让 S3 通知 Lambda？
# a
先允许 S3 调用 Lambda：
```bash
aws lambda add-permission \
  --function-name my-test-lambda \
  --statement-id allow-s3-invoke \
  --action lambda:InvokeFunction \
  --principal s3.amazonaws.com \
  --source-arn arn:aws:s3:::myapp-images-bucket \
  --source-account <你的账户ID>
```
然后配置 S3 存储桶的事件通知：
```bash
aws s3api put-bucket-notification-configuration \
    --bucket myapp-images-bucket \
    --notification-configuration '{
        "LambdaFunctionConfigurations": [
            {
                "LambdaFunctionArn": "arn:aws:lambda:<region>:<account-id>:function:my-test-lambda",
                "Events": ["s3:ObjectCreated:*"]
            }
        ]
    }'
```

