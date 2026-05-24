# q
如何手工下线 SAS/SATA 磁盘？
# a
```shell
echo 1 > /sys/block/sdb/device/delete
```

# q
如何手工下线 NVMe 磁盘？
# a
```shell
echo 1 > /sys/block/nvme0n1/device/device/remove
```

# q
如何触发 SCSI 控制器重新扫描所有磁盘设备？
# a
```shell
echo "- - -" > /sys/class/scsi_host/host0/scan
```

# q
如何让 NVMe 磁盘重新上线？
# a
```shell
echo 1 > /sys/bus/pci/rescan
```

