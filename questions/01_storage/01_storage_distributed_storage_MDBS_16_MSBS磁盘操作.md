# q
如何使用 disk_scan_offline.sh 脚本下线指定 IP 上的磁盘？
# a
```shell
./disk_scan_offline.sh 192.168.234.104 off sde
```
该命令将 192.168.234.104 上的 sde 磁盘下线。

# q
如何触发 SCSI 总线重新扫描以发现新添加的磁盘？
# a
使用以下命令扫描所有 SCSI 主机通道：
```shell
echo "- - -" > /sys/class/scsi_host/host0/scan
```

# q
如何通过 sysfs 立即删除 /dev/sdb 设备（使其离线）？
# a
```shell
echo 1 > /sys/block/sdb/device/delete
```
该命令会从系统中移除 sdb 磁盘，通常用于安全下线。

