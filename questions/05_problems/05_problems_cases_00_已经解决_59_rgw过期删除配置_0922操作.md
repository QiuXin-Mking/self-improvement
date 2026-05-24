# q
排查RadosGW生命周期过期删除策略未生效的标准流程是什么？
# a
1. 检查桶基本状态与分片：
   ```bash
   radosgw-admin bucket stats --bucket <bucket_name>
   radosgw-admin bucket radoslist --bucket <bucket_name>
   ```
2. 检查并修复桶索引：
   ```bash
   radosgw-admin bucket check --bucket <bucket_name> --fix
   ```
3. 查看多部分上传的生命周期索引条目：
   ```bash
   radosgw-admin bi list --bucket <bucket_name> --max-entries 100 --category multipart
   ```
4. 刷新生命周期策略：执行脚本 `00_document\03_对象存储\05_boto3_py\04_生命周期_fresh.py`
5. 确认生命周期线程是否存活：用 `top -H -p <radosgw_pid> -bn1 | grep life` 检查线程状态。

# q
如何手动触发某个存储桶的生命周期策略刷新？
# a
执行生命周期刷新脚本：`00_document\03_对象存储\05_boto3_py\04_生命周期_fresh.py`，并可通过 `top -H -p <radosgw_pid> -bn1 | grep life` 确认生命周期线程已启动。

