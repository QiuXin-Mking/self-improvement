# q
如何重新计算ACS API请求的`Signature`以解决签名验证失败问题？
# a
需按照ACS签名流程：  
1. 准备请求参数（如`Action`、`Timestamp`、`SignatureNonce`等）。  
2. 构造规范化查询字符串：将参数按照参数名排序，对值进行URL编码，用`&`拼接。  
3. 构建`string_to_sign`：`HTTPMethod&%2F&` 后接对规范化查询字符串的URL编码。  
   - 示例：  
     ```
     GET&%2F&AccessKeyId%3Ddf9bd5db2eb14fd0963758bb358f1f6f%26Action%3DDescribeContainers%26Format%3DJSON%26PageNumber%3D1%26PageSize%3D100%26RegionId%3Dyour-region%26SignatureMethod%3DHMAC-SHA1%26SignatureNonce%3D3c58853f-cdc1-461c-b511-b62c3da05a17%26SignatureVersion%3D1.0%26Timestamp%3D2025-05-30T06%253A47%253A11Z%26Version%3D1.0
     ```
4. 使用AccessKeySecret作为密钥，对`string_to_sign`进行HMAC-SHA1加密，然后进行Base64编码得到`Signature`。  
5. 将`Signature`进行URL编码后附加到请求中。

# q
构造签名请求时，`Timestamp`参数应如何处理？
# a
- `Timestamp`必须设置为当前UTC时间，格式为ISO 8601：`YYYY-MM-DDTHH:MM:SSZ`，例如`2025-05-30T06:47:11Z`。  
- 在构建`string_to_sign`时，需对`Timestamp`值进行双重URL编码：  
  - 冒号`:`第一次编码为`%3A`，即`2025-05-30T06%3A47%3A11Z`。  
  - 第二次编码（拼接进`string_to_sign`时）将`%`编码为`%25`，最终显示为`2025-05-30T06%253A47%253A11Z`。

# q
在keystone日志中看到签名错误，如何通过复制现有请求快速构造有效的测试请求？
# a
1. 从日志或已有请求中获取一个完整的curl命令（如向`/v2.0/ec2tokens`的POST请求）。  
2. 手动修改`Timestamp`为当前UTC时间和新的`SignatureNonce`（例如UUID）。  
3. 按ACS签名规则重新计算`Signature`：  
   - 抽取出所有业务参数（如`Action`、`Version`、`AccessKeyId`等）。  
   - 构造`string_to_sign`并进行HMAC-SHA1签名。  
4. 将新生成的`Signature`进行URL编码后替换到请求中，重新发送。

# q
计算签名的Python脚本需要实现哪些关键步骤？
# a
以下为关键步骤（基于源内容）：  
```python
import hmac
import hashlib
import base64
import urllib.parse
import datetime

def sign(access_key_secret, params):
    # 1. 参数排序并构建规范化查询字符串
    sorted_params = sorted(params.items())
    canonical_query = "&".join(
        f"{k}={urllib.parse.quote(str(v), safe='')}" for k, v in sorted_params
    )
    # 2. 构建string_to_sign
    string_to_sign = f"GET&{urllib.parse.quote('/', safe='')}&{urllib.parse.quote(canonical_query, safe='')}"
    # 3. HMAC-SHA1签名
    sig = hmac.new(
        access_key_secret.encode(), string_to_sign.encode(), hashlib.sha1
    ).digest()
    # 4. Base64编码
    return base64.b64encode(sig).decode()
```
注意：需要对特殊字符正确处理，特别是Timestamp中的冒号需编码为`%3A`。

