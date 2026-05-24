# q
HTTP Accept-Encoding 请求头的作用是什么
# a
用于告知服务器客户端支持哪些内容编码（如 gzip、deflate），以便服务器选择适当的压缩方式发送响应，从而减少传输数据量。

# q
Accept-Encoding 头中 identity 值的含义是什么
# a
`identity` 表示不接受任何内容编码，即要求服务器以原始、未压缩的格式发送响应，不进行任何转换。

# q
什么场景下适合使用 Accept-Encoding: identity
# a
- 需要确保响应完整性（避免解压错误）
- 响应内容本身已为最优编码或二进制数据，不适合再次压缩
- 客户端直接处理原始数据，无需解压缩步骤

