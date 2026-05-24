# q
如何查看当前块设备的文件系统类型？
# a
使用 ```lsblk -f``` 命令，它会列出所有块设备及其文件系统类型。

# q
ext4 文件系统在线扩容应使用什么命令？
# a
使用 ```resize2fs``` 命令，例如：```resize2fs /dev/vda2```

# q
xfs 文件系统在线扩容应使用什么命令？
# a
使用 ```xfs_growfs``` 命令，例如：```xfs_growfs /```

