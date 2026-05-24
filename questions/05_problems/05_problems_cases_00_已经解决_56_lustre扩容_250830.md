# q
如何查看Lustre文件系统的当前容量和各OST使用率？
# a
使用命令 `lfs df -h`，它会列出每个MDT和OST的UUID、总容量、已用空间、可用空间和使用率百分比，以及整个文件系统的汇总信息（filesystem_summary）。例如：
```
filesystem_summary:        19.8T       16.5T        2.3T  88% /mnt/lustre
```

# q
Lustre在线扩容新增OST的标准流程是什么？
# a
1. 在任意节点（如管理节点）创建RBD image：  
   ```bash
   rbd create --size 4096G rbd/mj_ost05 --image-feature layering
   ```
2. 在目标OSS节点映射RBD设备：  
   ```bash
   sudo rbd map rbd/mj_ost05 --id admin
   ```
3. 格式化为Lustre OST：  
   ```bash
   mkfs.lustre --fsname=mj_nas --ost --mgsnode=192.168.6.188@tcp --mgsnode=192.168.6.189@tcp --index=5 /dev/rbd3
   ```
   （关键参数：`--fsname` 文件系统名称，`--ost` 标记为对象存储目标，`--mgsnode` 管理节点NID，`--index` OST索引）
4. 创建挂载点并挂载：  
   ```bash
   mkdir -p /data/mj_ost05
   mount -t lustre /dev/rbd3 /data/mj_ost05
   ```
5. 使用 `lfs df -h` 确认新OST（如OST:5、OST:6）已加入且容量变化正确。

