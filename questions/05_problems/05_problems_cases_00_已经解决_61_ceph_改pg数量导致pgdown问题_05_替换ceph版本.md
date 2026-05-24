# q
如何搭建本地YUM仓库以安装自定义Ceph软件包？
# a
1. 将软件包tar.gz文件复制到目标目录（如 `/home/qiuxin/`）并解压。
2. 安装 `createrepo` 工具：`yum install -y createrepo`
3. 创建仓库元数据：`createrepo -pdo /home/qiuxin /home/qiuxin`
4. 配置repo文件（例如 `/etc/yum.repos.d/ceph_debug.repo`），内容如下：
```
[lustre]
name=lustre need Custom Repository
description=Local RPM packages
baseurl=file:///home/qiuxin
enabled=1
gpgcheck=0
```

# q
替换Ceph版本时，使用createrepo生成仓库元数据的完整命令是什么？
# a
```bash
createrepo -pdo /home/qiuxin /home/qiuxin
```
此命令在 `/home/qiuxin` 目录生成可被YUM使用的仓库元数据，选项 `-p` 保留解析的软件包，`-d` 生成SQLite数据库，`-o` 指定输出目录。

# q
如何通过repo文件将本地文件目录配置为YUM源？
# a
创建 `/etc/yum.repos.d/ceph_debug.repo` 文件，写入：
```
[lustre]
name=lustre need Custom Repository
description=Local RPM packages
baseurl=file:///home/qiuxin
enabled=1
gpgcheck=0
```
其中 `baseurl` 指定本地路径，`gpgcheck=0` 禁用GPG签名检查。

