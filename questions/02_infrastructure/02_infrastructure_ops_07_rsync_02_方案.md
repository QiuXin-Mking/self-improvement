# q
rsync 如何实现仅将远端变化的文件同步到本地，且本地变化不影响远端？
# a
使用 `rsync -avz -e ssh root@ees-0-3:/mnt/myrbd/ /mnt/myrbd`，
其中 `-a` 归档模式，`-v` 显示详情，`-z` 压缩传输，`-e ssh` 指定远程 shell 为 SSH。
该命令将远端目录 `/mnt/myrbd/` 的内容**单向同步**到本地 `/mnt/myrbd`，远端新增或修改的文件会被拉取，本地变动不会推送到远端。

# q
unison 如何实现本地与远端目录的双向同步？
# a
使用命令：
```
unison /mnt/myrbd ssh://root@ees-0-3///mnt/myrbd
```
unison 会检测两端的文件变化并自动合并冲突，实现双向同步。注意远程路径中使用 `///` 表示绝对路径。

# q
在 Ceph 中，如何创建 RBD 块设备、映射到本地并格式化为 XFS 文件系统？
# a
关键步骤与命令：
1. 创建 10G 的 RBD 镜像：
   ```
   rbd create testpool/rbd1 --size 10240
   ```
2. 映射到本地设备：
   ```
   rbd map testpool/rbd1
   ```
3. 查看映射结果：
   ```
   rbd showmapped
   ```
4. 创建 XFS 文件系统：
   ```
   mkfs.xfs /dev/rbd0
   ```
5. 挂载使用：
   ```
   mkdir /mnt/myrbd
   mount /dev/rbd0 /mnt/myrbd
   ```

# q
如何安全地取消 RBD 映射并卸载对应的挂载点？
# a
1. 查看当前映射：
   ```
   rbd showmapped
   ```
2. 取消映射指定设备（如 `/dev/rbd0`）：
   ```
   rbd unmap /dev/rbd0
   ```
3. 检查挂载信息：
   ```
   df -h
   ```
4. 卸载目录：
   ```
   umount /mnt/myrbd
   ```

# q
实现 inotify + rsync 实时同步方案需要安装哪些基础工具？
# a
在 RHEL/CentOS 系统上执行：
```
yum install inotify-tools rsync
```
其中 `inotify-tools` 用于监控文件系统变化，`rsync` 负责增量同步。

