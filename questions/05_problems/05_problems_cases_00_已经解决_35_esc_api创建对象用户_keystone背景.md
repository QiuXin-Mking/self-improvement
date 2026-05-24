# q
如何通过环境变量配置 OpenStack 客户端的认证信息？
# a
直接 `source keystonerc_admin` 或手动导出关键环境变量：
```bash
export OS_USERNAME=admin
export OS_PASSWORD=your_password
export OS_PROJECT_NAME=admin
export OS_AUTH_URL=http://keystone.example.com:5000/v3
export OS_USER_DOMAIN_NAME=Default
export OS_PROJECT_DOMAIN_NAME=Default
```
之后即可直接运行 `openstack project list` 等命令。

# q
如何使用 `curl` 向 Keystone v3 API 获取认证令牌？
# a
发送 POST 至 `/v3/auth/tokens`：
```bash
curl -X POST http://controller3:5000/v3/auth/tokens \
  -H "Content-Type: application/json" \
  -d '{
    "auth": {
      "identity": {
        "methods": ["password"],
        "password": {
          "user": {
            "name": "your_user",
            "password": "your_pass",
            "domain": {"name": "Default"}
          }
        }
      },
      "scope": {
        "project": {
          "name": "admin",
          "domain": {"name": "Default"}
        }
      }
    }
  }'
```

# q
Keystone 的 EC2 认证接口 `/v2.0/ec2tokens` 使用的签名算法是什么？
# a
采用 AWS v2 风格签名（HMAC-SHA1）。构造签名串格式为：
```
verb + "\n" + host + "\n" + path + "\n" + "AWSAccessKeyId=<ak>" + "&SignatureVersion=2" + "&Timestamp=<utc_timestamp>"
```
以 secret_key 为密钥，对上述字符串做 HMAC-SHA1 运算，再 base64 编码得到签名。示例：
```python
import hmac, hashlib, base64

def create_signature(secret_key, string_to_sign):
    dig = hmac.new(secret_key.encode('utf-8'), string_to_sign.encode('utf-8'), hashlib.sha1).digest()
    return base64.b64encode(dig).decode()
```

# q
在 Python 中如何通过 AK/SK 调用 Keystone EC2 认证获取 token？
# a
封装函数逻辑（参考 `acs_token.py`），构造 `ec2Credentials` 请求体，签名后 POST 到 `http://<keystone>:5000/v2.0/ec2tokens`，从响应中提取 `token = result['access']['token']['id']`。然后可以用该 token 调用其他 OpenStack API：
```python
from acs_token import call

ak = "your_access_key"
sk = "your_secret_key"
token = call(ak, sk, "http://controller3:5000")

headers = {"X-Auth-Token": token, "Content-type": "application/json"}
resp = requests.post("http://swift.endpoint/v1/users", json=data, headers=headers)
```

