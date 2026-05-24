# q
在Lustre冗余MGS场景中，客户端挂载命令如何指定多个MGS地址以实现故障切换？
# a
使用冒号分隔多个MGS地址，格式为：
```bash
mount -t lustre <mgs1>@tcp:<mgs2>@tcp:<mgs3>@tcp:/<fsname> /mnt/lustre
```
例如：
```bash
mount -t lustre 192.168.6.172@tcp:192.168.6.174@tcp:192.168.6.175@tcp:/nas_test /mnt/lustre
```

# q
如何从系统日志确认Lustre客户端已成功挂载并跳过恢复？
# a
检查内核日志（/var/log/messages 或 dmesg），出现类似以下输出表示客户端挂载成功：
```
Jan  3 14:58:00 stg_ssd_3-102-172 kernel: Lustre: nas_test-MDT0000: local client d9387807-9e27-46a8-8fbe-685d81154a1c w/o recovery
Jan  3 14:58:00 stg_ssd_3-102-172 kernel: Lustre: Mounted nas_test-client
```
关键信息：`w/o recovery` 表示无需恢复，`Mounted <fsname>-client` 确认挂载完成。

# q
在Lustre冗余MGS切换测试中，如何批量检查所有节点的文件系统挂载状态？
# a
使用Ansible批量执行：
```bash
ansible lustre -m shell -a 'df -h | grep lustre'
```
该命令输出各节点上所有Lustre文件系统的挂载点与容量信息，包括MDT、OST等后端设备挂载及客户端挂载（如 `/mnt/lustre`）。

# q
怎样查看Lustre服务器节点上Ceph RBD设备的映射关系？
# a
使用命令：
```bash
ansible lustre -m shell -a 'rbd showmapped'
```
输出示例：
```
id  pool         namespace  image         snap  device
0   lustre_pool             lustre_mdt3   -     /dev/rbd0
1   lustre_pool             lustre_ost12  -     /dev/rbd1
```
该命令显示每个RBD镜像所属的pool、镜像名、快照及映射到的本地块设备（如 /dev/rbd0）。

