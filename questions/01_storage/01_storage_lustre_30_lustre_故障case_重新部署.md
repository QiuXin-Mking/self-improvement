# q
如何从头开始重新部署一个Lustre文件系统（MGS、MDT、OST）？
# a
标准流程：
1. 在 MGS 节点上格式化并挂载管理服务：
   ```bash
   mkfs.lustre --fsname=nas_test --mgs --reformat /dev/vdc
   ```
2. 在各 MDT 节点上格式化并挂载元数据服务（示例 index=0）：
   ```bash
   mkfs.lustre --fsname=nas_test --mdt --index=0 --mgsnode=172.31.0.26@tcp --reformat /dev/vdh
   mount -t lustre /dev/vdh /data/lustre_mdt
   ```
3. 在 OST 节点上格式化并挂载对象存储服务（示例 index=1）：
   ```bash
   mkfs.lustre --fsname=nas_test --ost --mgsnode=172.31.0.26@tcp --index=1 --reformat /dev/vdc
   mount -t lustre /dev/vdc /data/lustre_ost_1/
   ```
4. 在客户端挂载文件系统：
   ```bash
   mount -t lustre 172.31.0.26@tcp:/nas_test /mnt/lustre
   ```
所有操作之前需先 `umount` 目标目录，确保旧挂载已卸载。

# q
挂载 Lustre OST 时出现 “mount.lustre: increased 'max_sectors_kb' from 1280 to 16384” 是什么原因？需要处理吗？
# a
这不是错误，而是 Lustre 挂载过程中自动调整内核参数 `max_sectors_kb` 的行为，从默认的 1280 增加到 16384，目的是优化大块 I/O 的吞吐量。该调整由挂载命令触发，正常完成挂载后无需任何人工干预。

# q
在 Lustre 部署中，多次执行 `mkfs.lustre --reformat` 会有什么风险？如何规避？
# a
`--reformat` 会强制重新格式化设备，清除原有所有数据。在运维脚本或批量部署时，若误对正在服务的设备执行此操作，会导致数据丢失。规避方法：
- 执行前先 `umount` 对应设备并确认挂载点已卸载。
- 使用前明确设备角色，检查是否有残留文件系统。
- 在自动化脚本中添加确认条件或备份步骤，避免重复运行格式化命令。

