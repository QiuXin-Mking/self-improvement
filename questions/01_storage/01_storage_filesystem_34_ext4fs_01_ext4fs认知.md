# q
如何创建一个小的ext4文件系统镜像并挂载？
# a
使用以下命令：
```bash
dd if=/dev/zero of=ext4fs.img bs=1M count=256
mkfs.ext4 ext4fs.img
mkdir mnt
sudo mount -o loop ext4fs.img mnt
```

# q
ext4文件系统第一个组（Group 0）由哪些区域组成？
# a
- 0: 超级块 (Super block)
- 1: 组描述符表 (Group descriptors)
- 2-32: 保留的组描述符表 (Reserved GDT Blocks)
- 33: 数据块位图 (Block bitmap)
- 35: inode位图 (Inode bitmap)
- 37-2084: inode表 (Inode table)
- 4139-32767: 数据块 (Data blocks)

# q
ext4中的超级块（super block）是什么，包含哪些关键信息？
# a
超级块是文件系统的“身份证”和控制面板，记录文件系统全局关键信息，包括文件系统UUID、块大小（block size）、每组块数（blocks per group）、inode总数、块总数等。

# q
ext4中的组描述符表（Group descriptors）作用是什么，记录哪些信息？
# a
组描述符表记录每个块组内重要区域的精确位置，相当于“楼层指南”。包含：块位图位置、inode位图位置、inode表位置、空闲块数、空闲inode数等。可通过 `dumpe2fs /dev/loop0 | grep -A 10 "Group 0:"` 查看。

# q
ext4中数据块位图（block bitmap）的存在价值和工作原理是什么？
# a
数据块位图用于跟踪数据块的使用状态，避免全盘扫描寻找空闲块。工作原理：查找空闲块时扫描位图找到0位并标记为1（占用）；释放块时将对应位清零（标记为空闲）。inode位图同理。

