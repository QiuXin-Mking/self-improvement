# q
如何使用 ceph-bluestore-tool 查看块设备的标签信息？
# a
使用 `show-label` 命令，指定设备路径：
```bash
ceph-bluestore-tool show-label --dev /dev/sdX
```

# q
如何对 BlueStore OSD 的文件系统进行一致性检查与修复？
# a
使用 `fsck` 命令，指定 OSD 的数据目录：
```bash
ceph-bluestore-tool fsck --path /var/lib/ceph/osd/ceph-X
```

# q
如何导出 BlueFS 的内部数据到指定目录？
# a
使用 `bluefs-export` 命令，指定 OSD 路径和输出目录：
```bash
ceph-bluestore-tool bluefs-export --path /var/lib/ceph/osd/ceph-X --out-dir /path/to/export
```

# q
如何在不依赖 MON 配置的情况下修复 OSD 中的对象？
# a
使用 `repair` 命令并添加 `--no-mon-config` 选项：
```bash
ceph-bluestore-tool repair --path /var/lib/ceph/osd/ceph-X --no-mon-config
```

# q
如何重置一个 OSD 的块设备，清除所有 BlueStore 数据？
# a
使用 `zap-device` 命令，指定 OSD 路径：
```bash
ceph-bluestore-tool zap-device --path /var/lib/ceph/osd/ceph-X
```

