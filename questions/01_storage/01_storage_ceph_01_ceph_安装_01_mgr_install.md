# q
安装 Ceph mgr 需要满足哪些条件？
# a
需要满足以下条件：
1. 已经安装好 ceph-deploy 工具（例如版本 2.0.1），可通过 `rpm -qa | grep ceph-deploy` 检查。
2. 需要在 `/etc/ceph` 目录下执行，因为需要使用 ceph.conf 配置文件。
3. `/etc/hosts` 文件中需要包含主机域名的解析条目。

# q
如何使用 ceph-deploy 创建 mgr 守护进程？
# a
在 `/etc/ceph` 目录下执行命令：
```bash
ceph-deploy mgr create ees23
```
其中 `ees23` 是目标主机名，该命令会在该节点上部署并启动 Ceph Manager（mgr）守护进程。

