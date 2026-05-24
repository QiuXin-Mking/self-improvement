# q
如何正确构造STS AssumeRole信任策略，避免角色创建失败？
# a
信任策略需使用 JSON 格式，指定 `"Effect": "Allow"`、`"Principal": {"AWS": ["arn:aws:iam:::user/用户名"]}` 和 `"Action": ["sts:AssumeRole"]`。完整的 CLI 创建命令示例：
```bash
aws iam create-role \
  --endpoint-url http://szys01.storage.wcsapi.com:5085 \
  --role-name 1220 \
  --path / \
  --assume-role-policy-document '{"Version":"2012-10-17","Statement":[{"Effect":"Allow","Principal":{"AWS":["arn:aws:iam:::user/liyq"]},"Action":["sts:AssumeRole"]}]}'
```
注意：`Principal` 必须是对象而非数组，且 ARN 需符合 IAM 用户格式。

# q
如何通过 CLI 快速测试 STS 角色创建、附加内联策略及清理流程？
# a
标准测试流程包括三个命令：
1. 创建角色（带信任策略）
2. 附着内联策略（授权具体操作和资源）
3. 测试完成后删除角色
示例：
```bash
# 创建角色
aws iam create-role --endpoint-url http://szys01.storage.wcsapi.com:5085 --role-name 1220 --path / --assume-role-policy-document '{"Version":"2012-10-17","Statement":[{"Effect":"Allow","Principal":{"AWS":["arn:aws:iam:::user/liyq"]},"Action":["sts:AssumeRole"]}]}'

# 附加策略（允许所有S3操作）
aws iam put-role-policy --endpoint-url http://szys01.storage.wcsapi.com:5085 --role-name 1220 --policy-name 12policy --policy-document '{"Version":"2012-10-17","Statement":[{"Effect":"Allow","Action":["s3:*"],"Resource":"arn:aws:s3:::*"}]}'

# 删除角色
aws iam delete-role --endpoint-url http://szys01.storage.wcsapi.com:5085 --role-name 1220
```
若测试更精细的权限，可调整 `Resource`，如限制到特定存储桶路径：`"arn:aws:s3:::liyq/123/*"`。

# q
在STS测试中附加角色内联策略时，为了限制权限范围，如何指定特定的S3资源路径？
# a
使用 `put-role-policy` 命令，在策略文档的 `Resource` 字段中限制 ARN。例如，仅允许对 `liyq/123/` 前缀下的对象进行操作：
```bash
aws iam put-role-policy --endpoint-url http://szys01.storage.wcsapi.com:5085 --role-name 1220 --policy-name 1220policy --policy-document '{"Version":"2012-10-17","Statement":[{"Effect":"Allow","Action":["s3:*"],"Resource":"arn:aws:s3:::liyq/123/*"}]}'
```
`Resource` 的值必须为 `arn:aws:s3:::<bucket-name>/<path>/*` 格式，路径末尾需包含通配符 `*` 以匹配该目录下所有对象。

