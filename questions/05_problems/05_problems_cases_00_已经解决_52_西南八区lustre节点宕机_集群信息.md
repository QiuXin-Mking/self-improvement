# q
如何通过df和rbd showmapped命令定位Lustre节点宕机问题？
# a
在正常节点（如 nas_net-65-130）上执行 `df -h` 或 `df`，会看到类似 `/dev/rbd0` 等设备挂载到 `/data/lustre_mdt00`、`/data/lustre_ost00` 等目标。同时 `rbd showmapped` 会列出所有已映射的 RBD 镜像及其设备。宕机节点（如 nas_net-65-131）上，`df` 输出中完全没有 Lustre 相关的挂载，`rbd showmapped` 输出为空，说明节点上的 Ceph RBD 卷未映射或节点未完全恢复。

# q
西南八区Lustre节点宕机的典型根因是什么？
# a
典型根因包括：节点操作系统层面故障（如 node 133 重启后无法进入系统，可能是系统盘损坏或内核崩溃）；节点重启后 Ceph RBD 卷未能自动映射和挂载（如 node 131 重启后 `rbd showmapped` 为空，所有 Lustre 服务端挂载缺失），导致该节点上的 MDT/OST 存储池不可用。

# q
解决Lustre节点宕机后RBD卷未自动挂载的标准流程是什么？
# a
1. 确认节点系统已正常启动，修复任何系统级故障（如 node 133 无法进入系统需先排查硬件/引导）。  
2. 检查 RBD 映射状态：执行 `rbd showmapped`，若输出为空则需手动映射。  
3. 手动映射所需 RBD 镜像，命令顺序参考：
   ```
   rbd map rbd/lustre_mdt00
   rbd map rbd/lustre_ost00
   rbd map rbd/lustre_ost01
   rbd map rbd/lustre_ost02
   rbd map rbd/lustre_ost03
   rbd map rbd/lustre_mgt00
   ```
4. 挂载到对应目录，命令参考：
   ```
   mount -t lustre /dev/rbd0 /data/lustre_mdt00
   mount -t lustre /dev/rbd1 /data/lustre_ost00
   mount -t lustre /dev/rbd2 /data/lustre_ost01
   mount -t lustre /dev/rbd3 /data/lustre_ost02
   mount -t lustre /dev/rbd4 /data/lustre_ost03
   mount -t lustre /dev/rbd5 /data/lustre_mgt00
   ```
5. 使用 `df -h` 验证挂载点是否显示，确保存储卷恢复在线。  
6. 如需在客户端重新挂载 Lustre 文件系统，执行：
   ```
   mount -t lustre -o noatime 10.5.73.130@tcp:10.5.73.131@tcp:/nas /mnt/lustre
   ```

