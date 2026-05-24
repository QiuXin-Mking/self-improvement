# q
MDT故障后在本节点恢复时，如何挂载Lustre目标设备？
# a
使用标准 mount 命令挂载，例如：
```bash
mount -t lustre /dev/rbd5 /data/lustre_data
```

# q
MDT跨节点恢复时，`tunefs.lustre --erase-params --servicenode=... --mgsnode=...` 的作用是什么？
# a
该命令用于清除原有参数并重新指定服务节点（servicenode）和 MGS 节点（mgsnode），使 MDT 设备迁移到新节点后能正确连接到 MGS。例如：
```bash
tunefs.lustre --erase-params --servicenode=192.168.6.172@tcp --mgsnode=192.168.6.174@tcp /dev/rbd0
```
执行后再用 `mount -t lustre` 挂载设备。

