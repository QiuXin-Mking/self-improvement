# q
如何使用 `ceph-deploy` 指定自定义的 `ceph.conf` 文件？
# a
通过 `--ceph_conf` 参数指定配置文件的完整路径，例如：
```bash
ceph-deploy --ceph_conf /path/to/your/ceph.conf mon create node1 node2
```
这样 `ceph-deploy` 将使用该路径下的 `ceph.conf` 进行部署，而忽略当前目录下的 `ceph.conf`。

# q
如何用 `ceph-deploy` 管理多个不同配置的 Ceph 集群？
# a
为每个集群准备独立的 `ceph.conf` 文件，执行命令时用 `--ceph_conf` 指向对应的配置文件即可切换环境，例如：
```bash
ceph-deploy --ceph_conf /data/ceph_cluster_x.conf mon add mon2
```
通过参数即时指定配置源，无需手动替换当前目录下的默认文件。

# q
如何查看 `ceph-deploy` 支持的全部命令行选项？
# a
运行命令：
```bash
ceph-deploy --help
```

