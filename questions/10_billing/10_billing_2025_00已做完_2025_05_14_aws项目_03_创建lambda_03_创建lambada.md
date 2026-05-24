# q
使用 AWS CLI 创建 Lambda 函数的基本命令是什么？
# a
```bash
aws lambda create-function \
  --function-name my-test-lambda \
  --runtime python3.9 \
  --role arn:aws:iam::xxxxxxxx:role/lambda-s3-role \
  --handler lambda_function.lambda_handler \
  --zip-file fileb://function.zip
```
其中 `lambda_function.lambda_handler` 表示入口模块文件名 `lambda_function.py` 中的 `lambda_handler` 函数。必须先使用 `zip function.zip lambda_function.py` 将代码打包为 ZIP 文件。

# q
Lambda 执行角色信任策略的核心内容是什么？
# a
信任策略必须允许 Lambda 服务代入该角色，典型 `trust-policy.json` 内容为：
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
使用 CLI 创建角色：`aws iam create-role --role-name lambda-s3-role --assume-role-policy-document file://trust-policy.json`。如果遇到 `AccessDenied` 错误，说明当前用户没有 `iam:CreateRole` 权限。

# q
通过 CLI 为 Lambda 添加 S3 触发器需要哪些步骤和命令？
# a
先授权 S3 调用 Lambda：
```bash
aws lambda add-permission \
  --function-name my-test-lambda \
  --statement-id allow-s3-invoke \
  --action lambda:InvokeFunction \
  --principal s3.amazonaws.com \
  --source-arn arn:aws:s3:::myapp-images-bucket \
  --source-account <你的账户ID>
```
再配置桶通知：
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

# q
打包 Lambda 代码时有哪些关键要求？
# a
- 必须使用标准 ZIP 格式，入口文件须在 ZIP 根目录。
- 入口函数格式为 `文件名.函数名`，Python 默认是 `lambda_function.lambda_handler`。
- 若使用第三方库，需本地 `pip install -t .` 安装到项目目录后一起打包，或通过 Lambda Layer 提供。

