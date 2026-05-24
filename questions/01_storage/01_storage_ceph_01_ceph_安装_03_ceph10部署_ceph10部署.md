# q
部署 Ceph 集群前需要安装哪些必备软件？
# a
需要安装 `unzip` 和 `createrepo`，使用命令：
```bash
yum install -y unzip createrepo
```

# q
如何基于本地 RPM 包目录创建私有 YUM 仓库？
# a
使用 `createrepo` 命令，指定输出目录和源目录：
```bash
createrepo -pdo /home/yumrepo /home/yumrepo
```
其中：
- `-p` 表示解析依赖
- `-d` 生成 SQLite 数据库
- `-o` 指定输出仓库目录

# q
Ceph 部署时需要创建哪些关键工作目录？
# a
需要创建以下目录：
```bash
mkdir -p /var/lib/ceph/bootstrap-mds
mkdir -p /var/lib/ceph/bootstrap-osd
mkdir -p /var/lib/ceph/bootstrap-rgw
mkdir -p /var/lib/ceph/mon
mkdir -p /var/lib/ceph/ospd
mkdir -p /var/lib/ceph/tmp
mkdir -p /var/run/ceph/
```
这些目录分别用于 MDS、OSD、RGW 的引导过程、MON 服务、OSD 服务、临时文件以及运行时文件。

# q
Ceph 相关目录的属主和权限应如何设置？
# a
所有 `/var/lib/ceph/`、`/var/log/ceph/` 和 `/var/run/ceph/` 下的目录，其属主和属组都应设置为 `ceph:ceph`，例如：
```bash
chown -R ceph:ceph /var/lib/ceph/
chown -R ceph:ceph /var/log/ceph/
chown -R ceph:ceph /var/run/ceph/
```

