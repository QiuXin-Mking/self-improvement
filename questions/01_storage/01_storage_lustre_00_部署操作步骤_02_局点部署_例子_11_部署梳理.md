# q
如何根据总容量和产品规格计算 Lustre 部署所需的 OST 数量？
# a
每个 OST 固定容量为 2T，总容量需求为 25T（来自 ceph 的可用容量），因此 OST 总数 = 25T / 2T = 12 个。部署在 3 个节点上，每个节点分配 4 个 OST。

# q
Lustre 部署中如何配置内核时钟源为 TSC？
# a
修改 /etc/default/grub，在 GRUB_CMDLINE_LINUX 参数末尾增加 `tsc=reliable tsc=nowatchdog`，然后执行 `grub2-mkconfig -o /boot/grub2/grub.cfg`（UEFI 启动需使用对应路径），重启后通过 `cat /sys/devices/system/clocksource/clocksource0/current_clocksource` 检查是否为 tsc。

# q
Lustre 部署时如何配置 LNET 网络并使其生效？
# a
编辑 `/etc/modprobe.d/lustre.conf`，写入 `options lnet networks=tcp(bond0.5)`，其中 bond0.5 为存储网卡名称。然后执行 `lustre_rmmod` 卸载 Lustre 内核模块，再执行 `modprobe lustre` 加载模块，最后用 `lctl list_nids` 查看配置是否生效。

