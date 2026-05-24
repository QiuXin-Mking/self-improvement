# q
如何使用ceph-deploy创建Ceph对象存储网关(RGW)？
# a
使用以下命令：
```sh
ceph-deploy --overwrite-conf rgw create test-1
```
该命令在节点 `test-1` 上创建 RGW 实例，`--overwrite-conf` 表示覆盖现有配置文件。

# q
如何查看Ceph集群的认证信息？
# a
执行：
```sh
ceph auth ls
```
可以列出所有 Ceph 认证实体及其权限。集群配置可通过 `cat /etc/ceph/ceph.conf` 查看。

# q
如何为 RGW 的 bootstrap 客户端生成 keyring 文件？
# a
使用 `ceph-authtool` 命令：
```sh
ceph-authtool /etc/ceph/ceph.client.bootstrap-rgw.keyring --create-keyring --name=client.bootstrap-rgw --gen-key --cap mon 'allow *' --cap osd 'allow *'
```
它会创建 keyring、生成密钥，并授予对 monitor 和 OSD 的完全访问权限。

# q
如何在 RHEL/CentOS 上安装 AWS CLI 工具？
# a
使用 yum 安装：
```sh
yum -y install awscli
```

