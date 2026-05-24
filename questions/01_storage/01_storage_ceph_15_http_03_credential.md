# q
HTTP中的“Credential”指的是什么，主要用途是什么？
# a
在HTTP中，“Credential”（凭证）是用于身份验证和授权的信息，允许客户端（如浏览器或API客户端）向服务器证明自己的身份，以获取对特定资源的访问权限。

# q
HTTP基本认证的凭证在请求中如何传输？
# a
HTTP基本认证的凭证包含在请求的 `Authorization` 头部字段中，内容为“Basic”后跟一个空格，然后是**Base64编码**的`用户名:密码`对（例如 `Authorization: Basic QWxhZGRpbjpvcGVuIHNlc2FtZQ==`）。

# q
HTTP摘要认证与基本认证在凭证处理上的核心区别是什么？
# a
基本认证直接发送Base64编码的用户名密码；摘要认证则使用更复杂的**哈希和加密过程**，通过挑战-应答机制，避免明文传输等效信息，从而提供更高的安全性。

# q
现代Web应用中，除了传统认证方式，凭证常以什么形式存在？
# a
现代Web应用常使用OAuth、OpenID Connect等协议，凭证通常以**令牌（token）** 的形式出现，例如访问令牌（access token）或刷新令牌（refresh token），通过 `Authorization` 头部携带。

# q
传输HTTP凭证时，如何保障安全性？
# a
凭证信息应避免被截获或泄露，实际应用中通常使用**HTTPS（HTTP Secure）** 对整个HTTP通信进行加密，确保凭证和其他敏感数据的安全传输。

