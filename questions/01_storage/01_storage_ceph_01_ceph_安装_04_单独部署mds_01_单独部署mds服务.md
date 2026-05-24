# q
单独部署ceph MDS服务的标准流程包含哪些步骤？
# a
1. 创建临时工作目录并进入：
```bash
mkdir -p /tmp/cephfs-cluster
cd /tmp/cephfs-cluster
```
2. 从第一个mon节点收集认证密钥：
```bash
ceph-deploy --overwrite-conf gatherkeys <mon_first_hostname>
```
3. 从目标mds节点拉取当前配置：
```bash
ceph-deploy --overwrite-conf config pull <mds_first_hostname>
```
4. 部署mds服务：
```bash
ceph-deploy --overwrite-conf mds <mds_first_hostname>
```

# q
`ceph-deploy gatherkeys` 命令在生产环境部署MDS时承担什么作用？
# a
从指定的mon守护进程节点收集集群认证密钥（包括`ceph.client.admin.keyring`等），使部署节点获得管理集群的权限，为后续推送配置和启动mds服务提供安全基础。

# q
为什么单独部署mds前需要执行 `ceph-deploy config pull`？
# a
从目标节点拉取现有的ceph配置文件到本地临时目录，确保部署时使用的配置与集群当前状态一致，避免因配置漂移导致mds启动失败。带上 `--overwrite-conf` 会直接覆盖本地旧配置。

