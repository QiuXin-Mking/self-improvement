# q
如何创建软RAID1磁盘阵列？
# a
使用 `mdadm`，示例：
```bash
sudo mdadm --create --verbose /dev/md0 --level=1 --raid-devices=2 /dev/vdb /dev/vdc
```
`--level=1` 表示镜像模式（RAID 1），`--raid-devices=2` 指定由两块磁盘组成。同理可创建多个阵列（如 `/dev/md1`、`/dev/md2`）。

# q
如何在软RAID阵列上创建LVM逻辑卷并格式化为XFS文件系统？
# a
1. 创建物理卷：`sudo pvcreate /dev/md0 /dev/md1`
2. 创建卷组：`sudo vgcreate raid_vg /dev/md0 /dev/md1`
3. 创建逻辑卷（占用全部空间）：`sudo lvcreate -l 100%VG -n raid_lv raid_vg`  
   若要多盘条带化：`sudo lvcreate -i 3 -I 64 -l 100%VG -n raid_lv raid_vg`
4. 格式化：`sudo mkfs.xfs /dev/raid_vg/raid_lv`
5. 挂载：
   ```bash
   sudo mkdir -p /mnt/raid
   sudo mount /dev/raid_vg/raid_lv /mnt/raid
   ```
6. 自动挂载：在 `/etc/fstab` 中添加
   ```
   /dev/raid_vg/raid_lv /mnt/raid xfs defaults 0 0
   ```

# q
配置NFS服务导出RAID存储并设置客户端自动挂载的标准流程是什么？
# a
1. 安装工具：`yum install nfs-utils`
2. 编辑 `/etc/exports`，添加导出条目：
   ```
   /mnt/raid *(rw,sync,no_subtree_check)
   ```
3. 重新加载配置：`sudo exportfs -ra`
4. 启动并启用 NFS 服务：
   ```bash
   systemctl start nfs-server
   systemctl status nfs-server
   ```
5. 客户端挂载：
   ```bash
   sudo mkdir -p /mnt/nfs_raid
   sudo mount 127.0.0.1:/mnt/raid /mnt/nfs_raid
   ```
6. 客户端自动挂载：在 `/etc/fstab` 中添加
   ```
   127.0.0.1:/mnt/raid /mnt/nfs_raid xfs defaults 0 0
   ```

