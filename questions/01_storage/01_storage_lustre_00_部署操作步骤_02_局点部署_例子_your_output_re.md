# q
在Lustre部署 中，如何根据Ceph提供的总容量计算每个节点所需的OST数量？
# a
从 `ceph df` 获取可用总容量（例如25 TiB），按产品规格每个OST固定2T计算所需总OST数量：`总容量 / 2T`（25T / 2T = 12个）。再除以节点数得到每个节点的OST数：`12 / 3 = 4`，即每个节点4个OST。

# q
部署Lustre时，将内核时钟源切换为tsc的关键步骤是什么？
# a
1. 编辑 `/etc/default/grub`，在 `GRUB_CMDLINE_LINUX` 参数末尾添加 `tsc=reliable tsc=nowatchdog`。  
2. 执行 `grub2-mkconfig -o /boot/grub2/grub.cfg` 使配置生效。  
3. 重启后通过 `cat /sys/devices/system/clocksource/clocksource0/current_clocksource` 验证输出为 `tsc`。

# q
创建Ceph RBD供Lustre使用的典型命令是什么？如何缩减已创建的RBD卷？
# a
创建命令：  
```bash
rbd create --size <大小>G rbd/<名称> --image-feature layering
```
例如：`rbd create --size 10G rbd/lustre_mgt00 --image-feature layering`  
缩减已有卷：  
```bash
rbd resize rbd/<名称> --size <新大小>G --allow-shrink
```
例如：`rbd resize rbd/lustre_ost00 --size 2048G --allow-shrink`

# q
使用 `mkfs.lustre` 格式化 MGS、MDT、OST 时分别需要哪些关键参数？
# a
所有格式均需指定 `--fsname=<文件系统名>` 和对应的块设备。  
- **MGS**：`--mgs`，并指定所有 MGS 节点 `--mgsnode=<IP>@tcp`（可多个）。  
  例：`mkfs.lustre --fsname=nas --mgs --mgsnode=10.5.73.130@tcp --mgsnode=10.5.73.131@tcp /dev/rbd5`  
- **MDT**：`--mdt --index=<n>`（从0开始），同样需 `--mgsnode=... --reformat`。  
  例：`mkfs.lustre --fsname=nas --mdt --index=0 --mgsnode=10.5.73.130@tcp --mgsnode=10.5.73.131@tcp --reformat /dev/rbd0`  
- **OST**：`--ost --index=<n>`，同样需 `--mgsnode=... --reformat`。  
  例：`mkfs.lustre --fsname=nas --ost --index=0 --mgsnode=10.5.73.130@tcp --mgsnode=10.5.73.131@tcp --reformat /dev/rbd1`

# q
Lustre网络配置中，如何指定 Lustre 使用的网络接口并使其生效？
# a
1. 编辑 `/etc/modprobe.d/lustre.conf`，写入：  
   `options lnet networks=tcp(<网卡名>)`  
   例如：`tcp(bond0.5)`  
2. 执行 `lustre_rmmod` 移除原有 Lustre 内核模块，再执行 `modprobe lustre` 重新加载。  
3. 使用 `lctl list_nids` 查看 NID 列表，确认网络配置已生效。

