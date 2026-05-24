# q
如何在Lustre挂载中设置文件访问时间更新策略？
# a
通过 `mount -t lustre -o remount` 命令指定 `relatime` 或 `noatime` 选项。  
`relatime`：只在文件的访问时间（atime）早于或等于其修改时间（mtime）时更新 atime，以减少不必要的磁盘写入。  
`noatime`：完全禁用 atime 更新，可以进一步提升性能。  

示例：
```
mount -t lustre -o remount,relatime 192.168.5.171@tcp:192.168.5.172@tcp:/st_nas /mnt/lustre
mount -t lustre -o remount,noatime 192.168.5.171@tcp:192.168.5.172@tcp:/st_nas /mnt/lustre
```

# q
Lustre 挂载时如何动态修改挂载选项而不卸载文件系统？
# a
使用 `mount -o remount` 命令，保持原有挂载点和文件系统类型不变，仅修改挂载选项。  
命令格式：
```
mount -t lustre -o remount,<新选项> <MGS节点>@<网络类型>:<FS名称> <挂载点>
```
例如，为挂载在 `/mnt/lustre` 的 Lustre 文件系统启用 `relatime`：
```
mount -t lustre -o remount,relatime 192.168.5.171@tcp:192.168.5.172@tcp:/st_nas /mnt/lustre
```
该操作不会中断客户端对文件系统的访问。

# q
如何解读 Lustre 挂载点 `df` 输出的容量信息？
# a
`df` 输出示例：
```
192.168.5.171@tcp:192.168.5.172@tcp:/st_nas  159T  102T   49T  68% /mnt/lustre
```
各字段含义：
- `192.168.5.171@tcp:192.168.5.172@tcp:/st_nas`：MGS/MDT 节点与文件系统名称，表示挂载的 Lustre 文件系统。
- `159T`：文件系统总容量。
- `102T`：已使用容量。
- `49T`：可用容量。
- `68%`：使用率。
- `/mnt/lustre`：客户端本地挂载点。

