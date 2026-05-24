# q
如何通过 ceph osd tree 查看 OSD 与主机的对应关系？
# a
使用 `ceph osd tree` 可以列出集群中所有 OSD 所在的主机和状态，例如：
```
-7         14.71277      host cmp_trv-167-3
 6    ssd   3.67819          osd.6               up   1.00000  1.00000
 7    ssd   3.67819          osd.7               up   1.00000  1.00000
 8    ssd   3.67819          osd.8               up   1.00000  1.00000
12    ssd   3.67819          osd.12              up   1.00000  1.00000
```
该输出显示 `cmp_trv-167-3` 节点上有 osd.6、7、8、12，均为 SSD，状态为 up。

# q
如何通过 LVM 命令查看 Ceph OSD 使用的物理卷？
# a
使用 `pvs` 命令查看物理卷及其所属的卷组，例如：
```
[root@cmp_trv-167-3 ~]# pvs
  PV           VG                                        Fmt  Attr PSize  PFree
  /dev/bcache0 ceph-17512a3a-a295-4997-8403-2dff5104df74 lvm2 a--  <3.64t    0
  /dev/bcache1 ceph-64155d61-e669-494f-b6a4-d1e76403a79f lvm2 a--  <3.64t    0
  /dev/bcache2 ceph-8b34456b-42bf-40c3-a301-0f1a25445c6f lvm2 a--  <3.64t    0
  /dev/bcache3 ceph-5d44221c-8ce2-46c6-8649-26e58379b58b lvm2 a--  <3.64t    0
```
每块 bcache 设备对应一个 Ceph 卷组，PV 大小约 3.64 TiB，无剩余空间。

# q
如何通过 lvs 查看 Ceph OSD 对应的逻辑卷？
# a
使用 `lvs` 命令可列出逻辑卷及其对应的卷组：
```
[root@cmp_trv-167-3 ~]# lvs
  LV                                             VG                                        Attr       LSize
  osd-block-f1d6c276-e7dc-46d8-ac36-76aa313c96e6 ceph-17512a3a-a295-4997-8403-2dff5104df74 -wi-ao---- <3.64t
  osd-block-6847ad01-1053-49c9-9eb3-5a74716899f1 ceph-5d44221c-8ce2-46c6-8649-26e58379b58b -wi-ao---- <3.64t
  osd-block-02404221-143e-40e8-8a19-36910368a3ab ceph-64155d61-e669-494f-b6a4-d1e76403a79f -wi-ao---- <3.64t
  osd-block-014f1112-5942-48ea-8048-d45c082987c5 ceph-8b34456b-42bf-40c3-a301-0f1a25445c6f -wi-ao---- <3.64t
```
每个 OSD 的块设备以 `osd-block-<uuid>` 命名，通过卷组与 PV 一一对应。

# q
如何通过 lsblk 查看 Ceph OSD 的 bcache 和 LVM 层次结构？
# a
执行 `lsblk` 可以看到从磁盘到 bcache 再到 LVM 的完整设备栈，例如：
```
sdc              8:32   0   3.7T  0 disk
└─sdc1           8:33   0   3.7T  0 part
  └─bcache0    253:0    0   3.7T  0 disk
    └─ceph--...--osd--block--f1d6c276... 252:1    0   3.7T  0 lvm
```
这表明物理磁盘 sdc 分区 sdc1 作为 bcache 的后端设备，bcache0 上直接承载 LVM 逻辑卷，该 LV 即为 OSD 的实际数据块设备。

