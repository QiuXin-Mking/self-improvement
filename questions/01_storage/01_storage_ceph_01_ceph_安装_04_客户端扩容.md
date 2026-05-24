# q
客户端执行 `ceph -s` 报错 `RADOS object not found (error connecting to the cluster)` 时，需要检查哪三个关键要素？
# a
需要检查：
1. **密钥环文件**：`/etc/ceph/ceph.client.admin.keyring`
2. **配置文件**：`/etc/ceph/ceph.conf`
3. **网络连通性**：使用 `ping <monitor-ip>` 和 `telnet <monitor-ip> 6789` 测试到 monitor 的连通性

# q
当客户端缺失 Ceph 配置或密钥文件时，如何从正常节点恢复？
# a
1. 备份原有文件（如存在）：
   ```bash
   cp /etc/ceph/ceph.conf /etc/ceph/ceph.conf.bak
   cp /etc/ceph/ceph.client.admin.keyring /etc/ceph/ceph.client.admin.keyring_back
   ```
2. 使用 `scp` 从正常节点复制：
   ```bash
   scp stg_ssd_3-102-172:/etc/ceph/ceph.conf /etc/ceph/ceph.conf
   scp stg_ssd_3-102-172:/etc/ceph/ceph.client.admin.keyring /etc/ceph/ceph.client.admin.keyring
   ```
3. 重新执行 `ceph -s` 验证集群连接

