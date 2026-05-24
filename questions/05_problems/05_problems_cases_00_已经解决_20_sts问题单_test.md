# q
STS AssumeRole 请求认证时 STSEngine 返回 denied (-13) 的典型根因是什么？
# a
典型根因是 credential scope 中 region 字段为空，导致签名验证过程中无法构建正确的 scope。从日志中可以看到 `credential scope = 20241231//sts/aws4_request`，中间的 region 部分缺失（应为 `20241231/<region>/sts/aws4_request`），引起 `STSEngine` 签名计算错误而拒绝请求。即使最终 `LocalEngine` 验证通过，STS 自身的认证引擎仍然要求正确的 region 信息。

# q
如何从日志定位 STS AssumeRole 认证失败的具体原因？
# a
1. 查找 `STSEngine denied with reason=-13` 日志行，确认是 STS 引擎拒绝了请求。
2. 查看上下文中的 `credential scope` 字段，检查是否包含完整的 `<date>/<region>/<service>/aws4_request` 格式，尤其注意 region 是否为空。
3. 例如以下日志片段明确指出 region 缺失：
```log
v4 credential format = LUJR29GB50LRW12HCPTO/20241231//sts/aws4_request
credential scope = 20241231//sts/aws4_request
```
4. 对比成功通过 `LocalEngine` 验证的日志，会发现 `generated signature` 与客户端签名一致，说明请求本身签名正确，问题仅在于 STS 认证链对 region 的校验。

# q
解决 STS AssumeRole 认证失败（region 缺失）的标准流程是什么？
# a
1. 确认请求发送方所使用的 STS endpoint 是否正确设置了 region（例如通过 `--region` 参数或在配置文件中指定）。
2. 检查 RGW 的 STS 服务是否在正确的 region 下配置并监听，确保 `rgw_sts_key` 等配置与 endpoint 匹配。
3. 修正客户端请求中的 credential scope，使 region 字段为非空值，例如改为 `20241231/<正确region>/sts/aws4_request`。
4. 若需保留无 region 的兼容性，可以考虑调整 RGW 的认证策略链顺序，使 `LocalEngine` 优先处理，但这通常不是标准做法。推荐修复客户端或服务端 region 配置。

