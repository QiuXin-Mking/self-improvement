# q
如何根据文件数量估算Lustre的MDT容量？
# a
假设单个inode大小为2.4K，文件数量为1亿，则总inode占用空间为 100000000 × 2.4K / 1024 / 1024 ≈ 228.88G，向上取整为250G。再除以MDT数量（如5个）可得到每个MDT约需50G。验证方式：若已知现有集群MDT 100G对应4.2千万inode，则扩容至250G时，线性估算inode数约为 4.2KW × 2.5 = 1.05亿，满足1亿文件需求。

# q
在Ceph中为Lustre创建RBD镜像的常用命令是什么？
# a
使用 `rbd create` 命令，指定镜像大小和特性：
```bash
rbd create --size 50G rbd/lustre_mdt00 --image-feature layering
rbd create --size 4860G rbd/lustre_ost00 --image-feature layering
rbd create --size 10G rbd/lustre_mgt00 --image-feature layering
```
其中 `--image-feature layering` 是必须启用的特性。

# q
Lustre部署中如何将内核时钟源切换为tsc？
# a
编辑 `/etc/default/grub`，在 `GRUB_CMDLINE_LINUX` 中添加 `tsc=reliable tsc=nowatchdog`，然后执行 `grub2-mkconfig -o /boot/grub2/grub.cfg`（UEFI系统可能需更新 `/boot/efi/EFI/centos/grub.cfg`），重启后通过 `cat /sys/devices/system/clocksource/clocksource0/current_clocksource` 验证是否为tsc。

# q
将Ceph RBD镜像映射到节点供Lustre使用的基本步骤是什么？
# a
在节点上执行 `sudo rbd map rbd/<镜像名> --id admin`，例如：
```bash
sudo rbd map rbd/lustre_mdt00 --id admin
sudo rbd map rbd/lustre_ost00 --id admin
```
映射后通过 `rbd showmapped` 查看映射列表。每个节点需按规划映射对应的MDT、OST和MGT镜像。

