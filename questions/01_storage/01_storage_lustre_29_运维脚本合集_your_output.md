# q
在 Lustre 集群规划中，如何根据已知的总可用容量和 OST 数量计算每个 OST 的容量？
# a
根据预留后的集群总容量（例如 95T）除以总 OST 数量（例如 20），并将结果向下取整。计算过程：  
```bash
95T * 1024 / 20 = 4864G  # 向下取整为 4860G
```
每个 OST 的容量设为 4860G。

# q
如何估算 Lustre MDT 所需的存储容量来支持 1 亿个文件？
# a
计算单个 inode 大小（如 1024 + 1536 = 2.4K）。  
```bash
1 亿 * 2.4K / 1024 / 1024 ≈ 228.88G  # 向上取整为 250G
```
再将总 MDT 容量除以 MDT 数量（如 5），得每个 MDT 50G。并通过已有集群验证：MDT 100G 承载约 4.2 千万 inode，线性推测 250G 可承载约 1.05 亿 inode。

# q
如何使用 Ceph RBD 为 Lustre 创建 MGT、MDT 和 OST 块设备？
# a
使用 `rbd create` 命令，指定大小和 RBD 镜像名称，并启用 layering 特性。示例：
```bash
rbd create --size 10G rbd/lustre_mgt00 --image-feature layering
rbd create --size 50G rbd/lustre_mdt00 --image-feature layering
rbd create --size 4860G rbd/lustre_ost00 --image-feature layering
```
创建后再在各节点映射：
```bash
sudo rbd map rbd/lustre_mdt00 --id admin
rbd showmapped
```

# q
在 Lustre 部署过程中，如何将内核时钟源切换为 TSC？
# a
修改 `/etc/default/grub`，在 `GRUB_CMDLINE_LINUX` 行末尾添加 `tsc=reliable tsc=nowatchdog`，然后执行：
```bash
grub2-mkconfig -o /boot/grub2/grub.cfg
reboot
```
重启后通过以下命令确认：
```bash
cat /sys/devices/system/clocksource/clocksource0/current_clocksource
# 输出应为 tsc
```

