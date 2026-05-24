# q
如何快速查看 Ceph 集群的整体健康状态？
# a
执行命令：
```bash
ceph -s
```
该命令会显示集群 ID、健康状态（HEALTH_OK/WARN/ERR）、MON 节点、OSD 数量、PG 状态、IO 速率等简要信息。

# q
如何临时下线一个 OSD 并检查 OSD 的磁盘使用情况？
# a
通过下列命令进行操作：

1. 将 OSD 临时下线：
```bash
ceph osd out osd.<id>
```
2. 查看所有 OSD 的容量使用详情：
```bash
ceph osd df
```

# q
如何修复出现不一致的 PG？
# a
对不一致的 PG 执行修复命令：
```bash
ceph pg repair <pg_id>
```
执行前建议先用 `ceph pg <pg_id> query` 查看 PG 详细信息，确认问题类型。

# q
如何为一个 RGW 桶设置并启用配额？
# a
依次执行以下命令：

1. 设置桶的最大大小（单位：字节）：
```bash
radosgw-admin quota set --bucket=<bucket-name> --max-size=<bytes>
```
2. 启用配额：
```bash
radosgw-admin quota enable --bucket=<bucket-name>
```
3. 验证配额信息：
```bash
radosgw-admin quota get --bucket=<bucket-name>
```

# q
如何查看 Lustre 文件系统中 MDT 的使用情况？
# a
使用 `lfs df` 命令并指定 MDT 相关选项：
```bash
lfs df -m /mnt/lustre
```
该命令会列出所有 MDT 的总容量、已用空间、使用百分比等信息。

