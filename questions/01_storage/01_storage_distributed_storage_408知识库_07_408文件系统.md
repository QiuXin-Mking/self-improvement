# q
选择文件系统时需要考虑哪些文件特征？
# a
选择文件系统前需要明确要存储的文件类型，常见特征包括：小文件但数量特别多，以及文件特别大但数量少。

# q
Ext3、Ext4、XFS 文件系统在文件大小限制和版本对应上有何核心差异？
# a
- Ext3：RHEL5 的文件系统，异常断电后可通过 `/lost+found` 目录恢复数据，Ext2 无此目录。
- Ext4：RHEL6 的文件系统，性能较 Ext3 提升，单个文件大小不能超过 1TB。
- XFS：RHEL7 默认文件系统，支持 18EB 存储容量，单个文件大小不能超过 16TB，更适合大数据场景。

# q
如何对 XFS 文件系统进行扩容？
# a
- 如果新磁盘大于 2TB，需要使用 `parted` 分区。
- 检查数据块大小和数量：`xfs_growfs info /dev/centos/root`
- 扩展到指定大小：`xfs_growfs /dev/centos/root -D 1986208`
- 自动扩展到最大可用空间：`xfs_growfs /dev/centos/root`

