# q
如何使用 vim 编辑 AWS 凭证文件？
# a
运行命令：
```
vim ~/.aws/credentials
```

# q
AWS 凭证文件中的默认配置示例包含哪些关键字段？
# a
```
[default]
aws_access_key_id = AKID1234567890
aws_secret_access_key = MY-SECRET-KEY
```
其中 `[default]` 是配置段名称，`aws_access_key_id` 和 `aws_secret_access_key` 分别存储访问密钥 ID 和私有访问密钥。

