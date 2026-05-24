# q
Lustre客户端挂载命令的基本格式是什么？
# a
```bash
mount -t lustre -o noatime <MGS NID1>@tcp:<MGS NID2>@tcp:/<文件系统名称> <挂载点>
```
示例：`mount -t lustre -o noatime 10.5.73.130@tcp:10.5.73.131@tcp:/nas /mnt/lustre`

# q
如何从系统已有挂载信息自动提取MGS NID和文件系统名称来生成客户端恢复挂载命令？
# a
通过 `df` 命令中匹配挂载点 `/mnt/lustre` 的行，提取该行的第一个字段（格式为 `<NID1>@tcp:<NID2>@tcp:/<fsname>`），再拼接完整的 `mount -t lustre -o noatime ...` 命令写入脚本。

# q
如何根据 `rbd showmapped` 的输出生成RBD映射恢复脚本（lustre_map.sh）？
# a
解析 `rbd showmapped` 的输出（跳过标题行），对每一行提取 `pool` 和 `image` 列的值，生成 `rbd map <pool>/<image>` 命令。例如 `rbd showmapped` 中有 `rbd   lustre_mdt00`，则生成 `rbd map rbd/lustre_mdt00`。

# q
如何从 `df -h | grep rbd` 的输出生成Lustre后端设备挂载恢复脚本（lustre_mount.sh）？
# a
过滤出 `df -h` 输出中包含 `rbd` 的行，提取设备路径（如 `/dev/rbd1`）和挂载点（如 `/data/lustre_ost00`），为每一对生成 `mount -t lustre <设备> <挂载点>` 命令。

# q
Lustre节点恢复需要的三个关键脚本及其内容来源是什么？
# a
1. `lustre_mount_client.sh`：用于重新挂载客户端，内容通过解析 `df` 输出中 `/mnt/lustre` 的源地址生成。  
2. `lustre_map.sh`：用于重新映射RBD设备，内容基于 `rbd showmapped` 输出的 pool/image 生成。  
3. `lustre_mount.sh`：用于挂载本地Lustre目标（MDT/OST），内容基于 `df -h | grep rbd` 输出的设备与挂载点关系生成。

