# q
MBR和GPT分区表的主要区别是什么？
# a
MBR分区表不支持超过2TB的磁盘容量，而GPT分区表支持。MBR的第一个扇区包含主要开机区（Master Boot Record）和分区表。

# q
Ext2文件系统中，一个区块群组（block group）由哪些部分组成，各自的作用是什么？
# a
一个区块群组包括：
- superblock：记录文件系统的整体信息，如inode/block的总量、使用量、剩余量及文件系统格式。
- inode：记录每个文件的属性（权限等）以及该文件所占用的block号码。
- data block：实际存放文件内容，文件太大会占用多个block。

# q
如何查询磁盘的超级区块（superblock）信息？
# a
使用命令：
```sh
dumpe2fs /dev/sdb
```

# q
什么是硬链接（hard link），它与inode的关系是什么？
# a
硬链接是多个文件名称对应同一个inode号码的关联记录。它只是在某个目录下新增一条文件名链接到某inode号码的记录，不会改变inode数目和磁盘空间。可通过 `ll -i` 查看文件对应的inode信息。

# q
使用 `gdisk` 完成分区后，如何让内核立即识别新的分区表？
# a
执行命令：
```sh
partprobe -s
```
此命令会更新核心分区表，之后可用 `lsblk` 查看结果。

