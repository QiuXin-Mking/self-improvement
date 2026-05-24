# q
配置远程YUM源时，nf.repo文件应包含哪些关键参数？
# a
nf.repo 文件内容为：
```c
[nf]
name=nf
baseurl=http://172.22.122.14/centos/
enabled=1
gpgcheck=0
```
将该文件放入 `/etc/yum.repos.d/` 目录，并备份或移除原有的 CentOS 源文件后，执行 `yum clean all && yum makecache` 即可生效。

# q
在 CentOS 7 上如何安装 iSCSI initiator 并检查是否安装成功？
# a
安装命令：
```bash
yum install -y iscsi-initiator-utils
```
检查是否安装：
```bash
rpm -qa | grep iscsi
```
安装后应显示 `iscsi-initiator-utils` 及其依赖包 `iscsi-initiator-utils-iscsiuio`。

# q
安装 iSCSI initiator 后需要执行哪些基本配置和服务操作？
# a
1. 设置服务开机自动启动：
```bash
chkconfig iscsi on
chkconfig iscsid on
```
2. 修改 initiator 配置文件（如需要）：
```bash
vi /etc/iscsi/iscsid.conf
```
3. 重启服务使配置生效：
```bash
service iscsi restart
```

