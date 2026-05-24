# q
Ceph对象存储兼容什么标准协议？
# a
Ceph对象存储通过RADOS网关（radosgw）兼容Amazon S3标准，具体文档参见：https://docs.ceph.com/en/pacific/radosgw/s3/

# q
Ceph对象存储的S3接口支持哪些常见的Bucket操作？
# a
支持的Bucket操作包括：
- list buckets（列出桶）
- delete buckets（删除桶）
- create buckets（创建桶）
- bucket lifecycle（桶生命周期管理）
- policy（桶策略）
- bucket website（静态网站托管）
- bucket ACLs（访问控制列表）
- bucket location（区域位置）
- bucket notification（桶事件通知）
- bucket object version（对象版本控制）
- get bucket info（获取桶信息）

