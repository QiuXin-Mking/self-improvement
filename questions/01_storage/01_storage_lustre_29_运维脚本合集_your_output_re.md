# q
如何计算 Lustre OST 的容量大小？
# a
计算步骤：首先确定集群总可用容量（总容量减去预留和已用），然后除以 OST 总数，向下取整。示例：总容量 95 T，20 个 OST，则 95 T × 1024 / 20 = 4864 G，向下取整为 4860 G。命令中使用的尺寸为 `--size 4860G`。

# q
如何根据目标文件数量估算 Lustre MDT 的容量？
# a
假设每个 inode 占用 2.4 K（1024+1536 字节），1 亿文件的总容量需求 ≈ 1亿 × 2.4 K / 1024² ≈ 228.88 G，向上取整为 250 G；再除以 MDT 数量（5），每个 MDT 约需 50 G。可通过现有集群验证：100 G MDT 容纳约 4.2 KW inode，250 G 可线性推算足够 1 亿 inode。

# q
Lustre 部署前为什么要将内核时钟源切换为 tsc？如何操作？
# a
切换的目的：确保分布式节点间时间戳的准确性和一致性，避免 TSC 不可靠导致性能或一致性问题。操作步骤：
1. 编辑 `/etc/default/grub`，在 `GRUB_CMDLINE_LINUX` 中添加 `tsc=reliable tsc=nowatchdog`；
2. 执行 `grub2-mkconfig -o /boot/grub2/grub.cfg` 更新引导配置；
3. 重启后通过 `cat /sys/devices/system/clocksource/clocksource0/current_clocksource` 检查输出是否为 `tsc`。

# q
如何使用 Ceph RBD 为 Lustre 创建存储并映射到目标节点？
# a
创建 RBD 镜像命令（示例）：
```bash
rbd create --size 4860G rbd/lustre_ost00 --image-feature layering
rbd create --size 50G rbd/lustre_mdt00 --image-feature layering
rbd create --size 10G rbd/lustre_mgt00 --image-feature layering
```
在目标节点映射：
```bash
sudo rbd map rbd/lustre_ost00 --id admin
sudo rbd map rbd/lustre_mdt00 --id admin
sudo rbd map rbd/lustre_mgt00 --id admin
```
使用 `rbd showmapped` 确认映射结果。

