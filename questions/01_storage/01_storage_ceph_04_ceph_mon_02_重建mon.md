# q
如何摧毁一个 Ceph Monitor 节点？
# a
使用 `ceph-deploy mon destroy <hostname>` 命令，例如：
```sh
ceph-deploy mon destroy ees-0-4
```

# q
如何重建一个 Ceph Monitor 节点？
# a
使用 `ceph-deploy --overwrite-conf mon add <hostname>` 命令，例如：
```sh
ceph-deploy --overwrite-conf mon add ees-0-4
```
`--overwrite-conf` 选项会覆盖目标节点上已存在的 Ceph 配置文件。

# q
重建 Monitor 时，目标节点的 `/etc/ceph/` 目录下需要有哪些文件？
# a
至少需要存在 `ceph.conf` 和 `ceph.mon.keyring` 两个文件。

# q
`ceph.mon.keyring` 文件的内容包含什么关键字段？
# a
包含 Monitor 密钥及其能力，典型格式如下：
```ini
[mon.]
    key = AQCaMfNjAAAAABAA1odkrx2o7NJP7hKzJIxeow==
    caps mon = "allow *"
```

