# q
如何通过Ceph RGW日志定位PUT请求返回404错误？
# a
使用 `zcat` 和 `grep` 组合，按请求关键词、操作类型（PUT）和状态码（404）过滤审计日志：
```bash
zcat ceph-client.rgw.ees-stxx-176-3.log-20250705.gz | grep "20250704185613583" | grep PUT | grep 404 | grep Audit
```
或直接 `fgrep "\"404\""` 从已提取的日志文件中筛选，再与关键词结合缩小范围。

# q
Ceph RGW PUT上传返回404的常见根因有哪些？
# a
- 目标 bucket 或对象不存在（即使 bucket 可通过 `aws s3 ls` 列出，对象路径可能错误）。
- 客户端重试导致部分请求在对象已被删除或尚未创建时到达；本例中 `aws-sdk-dotnet` 使用 `legacy` 重试模式，产生两次 PUT，均返回 404，请求时长分别达到 100 s 和 103 s，暗示存在重试冲突或超时。

# q
排查 Ceph RGW 404 问题时，Audit 日志中应重点关注哪些字段？
# a
- 时间戳：定位问题时间点。
- 操作和状态码：如 `PUT` + `"404"`。
- bucket 与对象路径：本例为 `upload` / `17837692415/20250704185613583.7z`。
- 请求 ID：`2~QtIXkjHZHKScDo6ns5xJ7xBnf-DsyOo-20`，用于追踪单次请求。
- 客户端 UA：`aws-sdk-dotnet-35/3.7.410.10 … retry-mode#legacy`，揭示客户端重试策略。
- 请求时长：`100.296104828s`，判断是否存在长耗时或累积重试。
- 来源 IP：`223.104.112.249` 与后端 IP：`183.220.37.212`，用于定位链路。

