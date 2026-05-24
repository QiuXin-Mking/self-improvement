# q
在高带宽NAS-stg_ssd3环境中，Ceph Monitor 节点有哪些 IP？
# a
Ceph Monitor 节点 IP 列表如下：
- 10.176.102.184
- 10.176.102.185
- 10.176.102.186
另外，存储节点 10.176.102.187 也属于该环境。

# q
一台 UEFI 启动的机器（如 197）关机后可能无法正常引导，如何修复 GRUB 配置？
# a
对于 UEFI 启动的机器，需要重新生成 GRUB 配置文件。执行以下命令之一（根据系统提示选择对应的输出路径）：
```bash
grub2-mkconfig -o /boot/efi/EFI/edgerayos/grub.cfg
grub2-mkconfig -o /boot/grub2/grub.cfg
grub2-mkconfig -o /boot/efi/EFI/centos/grub.cfg
```
若不确定，可先检查 `/boot/efi/EFI` 下的目录，以确定实际使用的引导条目。

