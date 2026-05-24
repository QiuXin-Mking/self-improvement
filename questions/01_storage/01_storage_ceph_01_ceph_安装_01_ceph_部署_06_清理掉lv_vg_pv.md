# q
如何删除Ceph OSD的逻辑卷（LV）？
# a
先使用 `umount` 卸载所有相关的挂载点（例如 `/dev/mapper/ceph-*/osd-*`），再执行 `lvremove` 删除对应的逻辑卷，命令示例：  
```bash
umount /dev/mapper/ceph-*/osd-*
lvremove /dev/ceph-xxx/osd-block-yyy
```

