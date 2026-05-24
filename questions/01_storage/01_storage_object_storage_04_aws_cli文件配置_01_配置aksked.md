# q
如何为AWS CLI配置自定义endpoint（如私有对象存储）？
# a
在 `/root/.aws/config` 文件的 `[default]` 配置段中设置 `endpoint_url` 参数：
```ini
[default]
region = oss-cn-beijing
output = json
endpoint_url = http://ees-southwest-8.edgeray.cn:5085
```
如果目录不存在，先执行 `mkdir -p /root/.aws` 创建。

# q
AWS CLI的AKSK如何配置？
# a
编辑 `/root/.aws/credentials` 文件，在 `[default]` 段中添加 `aws_access_key_id` 和 `aws_secret_access_key`：
```ini
[default]
aws_access_key_id = 7Y3GQ0ZW72XMWJ4B032N
aws_secret_access_key = 2R8WWUxPfLpF5p6qSKzfKq7j4QKZwQSKuPiStEDH
```

# q
AWS CLI的配置文件和凭证文件分别是什么？
# a
配置文件是 `/root/.aws/config`，用于设置 region、output、endpoint_url 等；凭证文件是 `/root/.aws/credentials`，用于存放访问密钥（aws_access_key_id 和 aws_secret_access_key）。两者都使用 INI 格式，默认配置段为 `[default]`。

