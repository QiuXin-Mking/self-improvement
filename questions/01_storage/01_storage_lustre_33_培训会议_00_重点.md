# q
Lustre挂载时需要特别注意哪两个问题？
# a
1. 使用 `mount -t lustre` 时若不指定文件系统类型默认会使用 ext4，可能导致异常。
2. mount 过程中出现错误时，尽量不要按 Ctrl+C 强制中断，否则会卡死，需等待约 10 分钟自动恢复。

# q
如何分别创建Lustre的MGS、MDT和OST设备？
# a
使用 `mkfs.lustre` 命令，示例：
```
# 创建MGS
mkfs.lustre --fsname=nas --mgs --mgsnode=172.31.0.26@tcp --reformat /dev/vdg

# 创建MDT（索引0）
mkfs.lustre --fsname=nas --mdt --mgsnode=172.31.0.26@tcp --reformat --index=0 /dev/vdb

# 创建OST（索引0）
mkfs.lustre --fsname=nas --ost --mgsnode=172.31.0.26@tcp --reformat --index=0 /dev/vdc
```

# q
Lustre客户端如何挂载文件系统？
# a
执行命令：
```
mount -t lustre 172.31.0.26@tcp:/nas /mnt/lustre
```
其中 `172.31.0.26@tcp:/nas` 是MGS节点地址和文件系统名称。

# q
Lustre节点扩容通常包含哪些步骤？
# a
1. 内核替换
2. 安装Lustre RPM包
3. 配置Lustre网络
4. 添加OST设备（如需要）
5. 添加MDT设备（如需要）

添加OST示例：
```
mkfs.lustre --fsname=nas --ost --mgsnode=172.31.0.26@tcp --reformat --index=1 /dev/vdc
```
添加MDT示例：
```
mkfs.lustre --fsname=nas --mdt --mgsnode=172.31.0.26@tcp --reformat --index=1 /dev/vdb
```

