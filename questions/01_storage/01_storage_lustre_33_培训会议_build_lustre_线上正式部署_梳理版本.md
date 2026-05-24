# q
在Lustre部署规划中，如何根据总容量与文件数量计算每个OST和MDT的容量大小？
# a
1. 计算Lustre可用总容量：从Ceph rbd池的USED(93 TiB) + MAX AVAIL(30 TiB) = 123 TiB，预留8 TiB并扣除赞齐项目20 TiB后，Lustre集群可用容量为95 TiB。  
2. 计算每个OST容量：总容量95 TiB ÷ 20个OST = 4864 GiB，向下取整为4860 GiB。  
3. 计算每个MDT容量：按1亿文件估算，单个inode约2.4 KiB(1024+1536)，总inode占用为1亿×2.4 KiB = 约228.88 GiB，向上取整250 GiB；MDT总容量250 GiB ÷ 5个MDT = 50 GiB。  
4. 验证：总容量 = 4860×20 + 50×5 + 10×1 = 97,460 GiB ≈ 95.18 TiB，基本符合预期。

# q
部署Lustre前需要完成哪些关键系统配置（内核版本与时钟源）？
# a
- 内核：必须安装并切换至 `kernel-4.18.0-3.2.1.er1.x86_64` 及其对应的开发与头文件包，使用 `rpm -ivh --nodeps --force` 安装。  
- 时钟源：修改 `/etc/default/grub`，在 `GRUB_CMDLINE_LINUX` 行尾添加 `tsc=reliable tsc=nowatchdog`，然后执行 `grub2-mkconfig -o /boot/grub2/grub.cfg` 生效。  
- 重启后通过 `uname -r` 确认内核版本，通过 `cat /sys/devices/system/clocksource/clocksource0/current_clocksource` 确认时钟源为 `tsc`。  
- 若无法切换内核（UEFI启动）或无法切换tsc，需额外处理（切换BIOS或咨询研发）。

# q
使用Ceph RBD为Lustre后端创建块设备时，标准的`rbd create`命令格式及关键选项是什么？
# a
```bash
rbd create --size <SIZE> rbd/<image_name> --image-feature layering
```
- `--size`：指定容量，如 `10G`、`50G`、`4860G`。  
- `rbd/<image_name>`：指定池rbd下的映像名称，如 `lustre_mgt00`、`lustre_mdt00`、`lustre_ost00`。  
- `--image-feature layering`：启用layering特性，用于后续克隆等操作。

# q
在离线或本地环境下，如何创建本地Yum源并安装Lustre依赖包？
# a
1. 安装`createrepo`工具：`yum install -y createrepo`。  
2. 在存放rpm的目录（如`/home/lustre_need`）执行：`createrepo -pdo /home/lustre_need /home/lustre_need`。  
3. 创建repo文件 `/etc/yum.repos.d/lustre.repo`，内容为：
```ini
[lustre]
name=lustre need Custom Repository
description=Local RPM packages
baseurl=file:///home/lustre_need
enabled=1
gpgcheck=0
```
4. 清理缓存并安装依赖：`yum clean all && yum makecache fast`，然后`yum install libmount-devel libyaml-devel e2fsprogs-devel e2fsprogs`。

