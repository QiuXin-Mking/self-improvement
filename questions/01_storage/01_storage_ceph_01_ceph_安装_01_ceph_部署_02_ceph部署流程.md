# q
SELinux是什么？它的三种运行模式分别是什么？
# a
SELinux（Security-Enhanced Linux）是Linux内核的安全子系统，提供基于类型强制（Type Enforcement）的访问控制安全策略。三种模式：
- Enforcing（强制模式）：安全策略被强制实施。
- Permissive（宽容模式）：仅打印警告信息，不强制实施策略。
- Disabled（禁用模式）：SELinux完全不运行。

# q
chrony服务由哪两个主要程序组成？chronyd的作用是什么？
# a
chrony服务由 `chronyd` 和 `chronyc` 两个程序组成。
- chronyd：后台守护进程，负责与NTP服务器同步系统时钟，并平滑调整时间速率。
- chronyc：命令行用户界面，用于监控性能和进行配置。

# q
使用 `ceph-deploy new` 命令初始化Ceph集群时，通常需要指定哪些网络参数？该命令会生成哪些关键文件？
# a
需要指定 `--public-network`（公有网络，如 `192.168.93.0/24`）和 `--cluster-network`（集群网络，如 `192.168.94.0/24`），并跟随首个Monitor节点主机名（如 `test-1`）。命令会生成三个文件：
- `ceph.conf`：Ceph配置文件。
- `ceph.mon.keyring`：Monitor的密钥环。
- `ceph-deploy-ceph.log`：部署日志文件。

# q
执行 `ceph-deploy mon create-initial` 创建初始Monitor后，会生成哪些用于后续认证的bootstrap keyring文件？
# a
生成的keyring文件包括：
```
ceph.bootstrap-mds.keyring
ceph.bootstrap-mgr.keyring
ceph.bootstrap-osd.keyring
ceph.bootstrap-rgw.keyring
ceph.client.admin.keyring
ceph.mon.keyring
```

# q
使用 `ceph-deploy` 添加OSD时，如何指定数据盘？磁盘在Ceph集群中最终以什么格式存在？
# a
使用命令 `ceph-deploy osd create --data /dev/sdb <节点名>` 指定数据盘。磁盘会被创建为LVM格式，在物理卷（PV）上创建 `ceph-` 前缀的卷组（VG），并在其中生成 `osd-block-` 前缀的逻辑卷（LV），例如：
```
/dev/sdb -> ceph-... VG -> osd-block-... LV
```

