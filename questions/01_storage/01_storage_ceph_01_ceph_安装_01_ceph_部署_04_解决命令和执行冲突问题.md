# q
如何解决使用 ceph-deploy 创建 monitor 时出现的 “config file /etc/ceph/ceph.conf exists with different content” 错误？
# a
在 `ceph-deploy mon create` 命令中添加 `--overwrite-conf` 选项强制覆盖现有配置文件：
```bash
ceph-deploy mon create --overwrite-conf ees24, ees23
```

