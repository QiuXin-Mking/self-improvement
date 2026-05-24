# q
Ceph源码包是什么性质的包？是否只能在Debian/Ubuntu上使用？
# a
Ceph源码包（如ceph_16.2.7.orig.tar.gz）是官方Ceph项目发行的中性源码包，理论上可以用于任何Linux发行版，不限于Debian/Ubuntu，也可以在CentOS/RHEL上编译安装。

# q
如何下载Ceph 16.2.7的源码包？
# a
```bash
wget https://download.ceph.com/debian-16.2.7/pool/main/c/ceph/ceph_16.2.7.orig.tar.gz
```

# q
如何解压`ceph_16.2.7.orig.tar.gz`？
# a
```bash
tar xf ceph_16.2.7.orig.tar.gz
```

# q
`tar xf` 命令中 x 和 f 参数分别代表什么？`tar xzf` 中的 z 又代表什么？
# a
- `x`：extract，解包（解压）归档文件。
- `f`：file，指定要操作的归档文件名（后面紧跟文件名）。
- `z`：用gzip解压（针对`.tar.gz`或`.tgz`文件），现代tar通常能自动识别压缩格式，所以`xf`即可解压`.tar.gz`。

