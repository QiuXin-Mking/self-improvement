# q
在 Linux x86_64 上安装 AWS CLI v2 的完整步骤是什么？
# a
```bash
curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip"
unzip awscliv2.zip
sudo ./aws/install
```
安装后通过 `aws --version` 验证。

# q
使用 aws s3api 查看 S3 存储桶 ACL 的命令是什么？
# a
```bash
aws s3api get-bucket-acl --bucket <bucket-name>
```
该命令返回包含 Owner 和 Grants 的 JSON 结构。

# q
如何通过 aws s3api 将桶的 ACL 设置为“私有”（private）？
# a
```bash
aws s3api put-bucket-acl --bucket <bucket-name> --acl private
```

# q
使用 JSON 文件设置自定义桶 ACL 的正确命令是什么？常见的错误是什么？
# a
正确命令：
```bash
aws s3api put-bucket-acl --bucket <bucket-name> --access-control-policy file://acl.json
```
常见错误：误用 `--acl file://acl.json`，该语法无效；应使用 `--access-control-policy` 指定自定义策略文件。

