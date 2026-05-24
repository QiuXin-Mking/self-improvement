# q
如何从 Ceph 集群中安全移除一个 OSD 的标准命令流程是什么
# a
首先需要设置集群标志以防止数据移动：
```bash
ceph osd set nobackfill
ceph osd set norecover
```
然后按顺序执行移除操作，假设要移除的 OSD id 为 10：
```bash
id=10;ceph osd down osd.$id; sleep 1;ceph osd out osd.$id; sleep 1; ceph osd rm osd.$id; sleep 1; ceph auth del osd.$id; umount /var/lib/ceph/osd/ceph-$id
```
最后记得清理 LVM 卷组和物理设备（如后续步骤所示）。

# q
如何彻底清理一块曾被 Ceph OSD 使用的 NVMe 磁盘（例如 nvme0n1）以重新部署
# a
完整清理与重新分区流程如下：
1. 移除 LVM 卷组：
   ```bash
   vgremove -f ceph-e92bc733-1320-41de-ad81-218db0b4b11d
   ```
2. 通知内核重新读取分区表：
   ```bash
   partprobe /dev/nvme0n1
   ```
3. 安全格式化 NVMe 命名空间（可选，取决于需求）：
   ```bash
   nvme format -s 1 /dev/nvme0n1 -l 0
   ```
4. 快速擦除开头数据：
   ```bash
   dd if=/dev/zero of=/dev/nvme0n1 bs=10M count=50
   ```
5. 创建 GPT 分区表并分配整个磁盘：
   ```bash
   parted -a optimal --script /dev/nvme0n1 mklabel gpt
   parted -a optimal --script /dev/nvme0n1 mkpart primary 0TB 100%
   partprobe /dev/nvme0n1
   ```
6. 使用 Ceph 工具清除分区上的 LVM 元数据：
   ```bash
   ceph-volume lvm zap /dev/nvme0n1p1
   ```
7. 可选：手动清除挂载点目录：
   ```bash
   umount /var/lib/ceph/osd/ceph-10/
   rm -rf  /var/lib/ceph/osd/ceph-10/*
   ```
   并再次 zap 确保干净：
   ```bash
   ceph-volume lvm zap /dev/nvme0n1p1
   ```

# q
在 OSD 移除过程中，为什么要先设置 `nobackfill` 和 `norecover` 标志
# a
设置 `nobackfill` 和 `norecover` 可以禁止集群在 OSD 上下线时启动数据回填和恢复操作，避免因移除单个 OSD 引发大规模数据迁移影响集群性能，并在可控的情况下完成维护后再恢复。

