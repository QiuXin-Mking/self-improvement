# q
如何临时将指定OSD的debug_osd调试级别设置为20？
# a
```bash
ceph daemon /var/run/ceph/ceph-osd.22.asok config set debug_osd 20
```
该命令通过OSD的管理套接字（admin socket）文件动态修改调试日志级别，适用于OSD.22。如需对其他OSD操作，替换对应的 `.asok` 文件路径。

# q
如何为OSD.65临时开启debug_osd调试（级别20）？
# a
```bash
ceph daemon /var/run/ceph/ceph-osd.65.asok config set debug_osd 20
```
使用 `ceph daemon` 命令与指定OSD的管理套接字通信，直接设置运行时参数。此操作即时生效，无需重启OSD进程。

