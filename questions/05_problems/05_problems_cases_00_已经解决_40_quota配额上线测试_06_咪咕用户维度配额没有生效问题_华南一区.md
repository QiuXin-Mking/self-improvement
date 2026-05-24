# q
如何从日志定位咪咕用户维度配额未生效的问题？
# a
在 RGW 节点上通过 admin socket 提升日志级别，执行触发配额检查的 S3 操作，然后检查 RGW 日志：
```bash
# 开启调试日志
ceph --admin-daemon /var/run/ceph/client.rgw.ees-gdla-250-4.asok config set debug_rgw 20

# 执行上传操作，触发配额检查
aws s3 cp --endpoint-url http://127.0.0.1:5085 --recursive test s3://quota/

# 查看 RGW 日志，关注配额计算与检查相关条目（日志路径视具体部署而定）
# 分析完毕后关闭调试日志
ceph --admin-daemon /var/run/ceph/client.rgw.ees-gdla-250-4.asok config set debug_rgw 0
```

# q
解决咪咕用户维度配额没有生效问题的标准流程是什么？
# a
标准排查与修复流程：
1. 在 RGW 节点提升日志级别以捕获详细信息：
```bash
ceph --admin-daemon /var/run/ceph/client.rgw.ees-gdla-250-4.asok config set debug_rgw 20
```
2. 使用客户提供的 AK/SK 及 endpoint 配置 AWS CLI，执行触发配额的操作：
```bash
aws s3 cp --endpoint-url http://127.0.0.1:5085 --recursive test s3://quota/
```
3. 检查 RGW 日志中配额统计与检查信息，确认配额是否因缓存未更新。
4. 若发现配额缓存 TTL 过长，动态调小 `rgw_bucket_quota_ttl` 参数：
```bash
ceph --admin-daemon /var/run/ceph/client.rgw.ees-gdla-250-4.asok config set rgw_bucket_quota_ttl 2400
# 验证参数已生效
ceph --admin-daemon /var/run/ceph/client.rgw.ees-gdla-250-4.asok config get rgw_bucket_quota_ttl
```
5. 再次执行上传操作，验证配额限制是否生效。
6. 问题解决后恢复日志级别：
```bash
ceph --admin-daemon /var/run/ceph/client.rgw.ees-gdla-250-4.asok config set debug_rgw 0
```

# q
咪咕用户维度配额没有生效的典型根因是什么？
# a
典型根因是 `rgw_bucket_quota_ttl` 参数值设置过大，导致桶配额统计信息缓存长时间不更新。当用户写入数据后，RGW 仍使用旧的已用容量判断是否超配额，使得配额限制未能及时触发。将该参数调整为较小值（如 2400 秒）可强制更频繁地刷新配额缓存，使配额立即生效。调整命令示例：
```bash
ceph --admin-daemon /var/run/ceph/client.rgw.ees-gdla-250-4.asok config set rgw_bucket_quota_ttl 2400
```

