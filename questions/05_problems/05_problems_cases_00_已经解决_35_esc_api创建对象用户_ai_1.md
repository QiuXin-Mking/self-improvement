# q
AK/SK鉴权请求中缺少Signature或AccessKeyId参数时的错误日志和返回响应是什么？
# a
日志中会输出“signature or accessKeyId key does not exist”，并返回由`_reject_request()`生成的拒绝响应。

# q
AK/SK签名校验失败时，日志中会记录什么关键错误信息？
# a
日志中记录“check signature error.”，随后返回`_reject_request()`响应。

# q
该ACS token中间件处理AK/SK鉴权的完整流程是怎样的？
# a
1. 从请求参数中提取`Signature`和`AccessKeyId`；缺失则记录日志并拒绝请求。
2. 构建`credentials`字典，先查询缓存中是否存在该Access对应的secret。
3. 若缓存命中且签名匹配，则使用缓存的token设置`X-Auth-Token`并转发请求。
4. 若缓存未命中或签名不匹配，通过POST到`CONF.keystone_acs_token.url`获取token；若返回403则直接返回Unauthorized响应。
5. 从响应中提取`token_id`和`user_id`，缓存token，再调用GET接口获取credential secret并缓存。
6. 将`X-Auth-Token`设置为新token，交由后续应用处理。

