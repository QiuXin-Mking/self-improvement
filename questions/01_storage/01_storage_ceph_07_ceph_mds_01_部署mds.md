# q
如何部署 Ceph MDS（元数据服务器）的标准步骤是什么？
# a
在部署节点上执行以下操作：
```bash
mkdir -p /tmp/cephfs-cluster
cd /tmp/cephfs-cluster
ceph-deploy --overwrite-conf gatherkeys <mon_first_hostname>
ceph-deploy --overwrite-conf config pull <mds_first_hostname>
ceph-deploy --overwrite-conf mds <mds_first_hostname>
```
这些命令会创建临时目录，从 MON 节点收集密钥，拉取已有配置，然后安装并启动 MDS 服务。

# q
在 MDS 部署过程中，`ceph-deploy gatherkeys` 和 `ceph-deploy config pull` 命令的作用是什么？
# a
- `ceph-deploy gatherkeys <mon_first_hostname>`：从指定的 MON 节点收集集群的认证密钥（如 `ceph.client.admin.keyring`、`ceph.bootstrap-mds.keyring` 等），并保存到本地临时目录中。
- `ceph-deploy config pull <mds_first_hostname>`：从已有的集群节点拉取 Ceph 配置文件（`ceph.conf`），并覆盖本地配置，以确保后续 MDS 节点使用一致的集群配置。

