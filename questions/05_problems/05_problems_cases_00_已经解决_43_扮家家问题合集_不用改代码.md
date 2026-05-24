# q
如何通过请求ID从压缩日志中检索Ceph RGW的审计日志？
# a
使用ansible在存储节点上执行shell，通过`grep`过滤请求ID和`Audit`关键字。示例命令：
```bash
ansible ceph_osd -m shell -a 'zcat /var/log/ceph/radosgw/ceph-client.rgw.$(hostname).log-20250705.gz | grep 2~3wxQ-FP1xmpzS7qJMF-rza-BHLVsNmP | grep Audit'
```
该命令会解压指定日期的gzip日志，并输出所有包含该请求ID的审计记录。

# q
Ceph RGW日志中`complete_multipart`操作记录了什么关键信息？
# a
记录分片上传完成时计算的ETag。示例日志：
```
2025-07-04T18:38:44.837+0800 ... s3:complete_multipart calculated etag: 57f103b19be2b33a5c3e2c5a9b78bbee-103
```
其中`57f103b19be2b33a5c3e2c5a9b78bbee-103`是最终对象的ETag，用于验证上传完整性。

