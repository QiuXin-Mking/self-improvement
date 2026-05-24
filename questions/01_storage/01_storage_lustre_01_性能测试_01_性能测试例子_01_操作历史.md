# q
在 Lustre 性能测试环境准备中，如何检查各节点上 Docker 容器是否都已关闭？
# a
使用 `docker ps` 命令，输出仅显示表头（`CONTAINER ID   IMAGE     COMMAND   CREATED   STATUS    PORTS     NAMES`）而无任何容器记录，即表示所有 Docker 容器已关闭。

# q
Lustre 性能测试前，需通过 `/etc/modprobe.d/modprobe.conf` 设置哪些内核模块参数？
# a
需添加以下配置：
```bash
echo "options ptlrpc ptlrpcd per cpt max=32" >> /etc/modprobe.d/modprobe.conf
echo "options ksockInd credits=2560" >> /etc/modprobe.d/modprobe.conf
```
其中 `ptlrpcd per cpt max=32` 控制每个 CPT 的 ptlrpcd 线程数上限，`ksockInd credits=2560` 设置 ksocklnd 传输信用数。

# q
如何批量停止 Lustre 测试节点上的 Ceph 相关服务？
# a
依次执行以下 `systemctl stop` 命令：
```bash
systemctl stop ceph-mds.target
systemctl stop ceph-mgr.target
systemctl stop ceph-mon.target
systemctl stop ceph-osd.target
systemctl stop ceph.target
```
或通过 Ansible 批量执行：
```bash
ansible lustre -m shell -a 'systemctl stop ceph-mds.target'
ansible lustre -m shell -a 'systemctl stop ceph-mgr.target'
ansible lustre -m shell -a 'systemctl stop ceph-mon.target'
ansible lustre -m shell -a 'systemctl stop ceph-osd.target'
ansible lustre -m shell -a 'systemctl stop ceph.target'
```

