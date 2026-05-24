# q
如何使用curl发送PUT请求来更新S3兼容对象存储桶的配置（包括ACL、配额等）？
# a
使用以下curl命令，将JSON配置通过`--data`传递给指定存储桶的API端点：
```bash
curl -i -X PUT "http://183.220.37.211:20080/rest/s3/bucket/goofys-test/" \
     -H "User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:135.0) Gecko/20100101 Firefox/135.0" \
     -H "Accept: application/json, text/plain, */*" \
     -H "Content-Type: application/json" \
     -H "X-Requested-With: XMLHttpRequest" \
     --data '{
       "acl": "public-read-write",
       "enable_bucket_max_size_quota": true,
       "enable_bucket_max_objects_quota": true,
       "enable_bucket_object_lock": false,
       "enable_version": null,
       "max_size": 1,
       "max_objects": 44,
       "bucket_object_lock_unit": "days",
       "bucket_object_lock_time": 1
     }'
```

# q
该请求的`--data` JSON中包含了哪些配置项，如何启用存储桶的大小和对象数配额？
# a
JSON配置项及作用：
- `acl`: 设为 `"public-read-write"`，设置存储桶访问权限。
- `enable_bucket_max_size_quota`: `true`，启用存储桶最大容量配额。
- `enable_bucket_max_objects_quota`: `true`，启用存储桶最大对象数配额。
- `max_size`: `1`，最大容量限制（单位通常为MB，具体由系统定义）。
- `max_objects`: `44`，最大对象数限制。
- `enable_bucket_object_lock`: `false`，未启用对象锁。
- `bucket_object_lock_unit` / `bucket_object_lock_time`: 对象锁保留单位（`"days"`）和时长（`1`），仅在启用锁时有效。
- `enable_version`: `null`，版本控制未启用。

