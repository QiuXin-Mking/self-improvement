# q
如何使用 debugfs 查看 Lustre MGS 文件系统的元数据？
# a
使用 `debugfs /dev/vde` 进入交互模式，常用命令包括：`ls` 列出目录内容，`stat /` 查看根目录 inode 详情，`show_super_stats` 显示超级块统计信息及块组概况。

# q
debugfs 中 `stat /` 输出 Inode: 2 且 Type: directory 表示什么？
# a
Inode 2 通常是 ext 文件系统的根目录 inode 号，Type 为 directory 表示该 inode 对应的是一个目录。其他字段如 Mode 0755 表示权限，Links 8 表示硬链接数，Size 4096 表示目录占用字节数，Blocks 列出数据块位置。

# q
`show_super_stats` 输出中的 `has_journal`、`ext_attr`、`dir_index`、`flex_bg` 特性分别有什么作用？
# a
- `has_journal`：启用日志，加快一致性恢复；
- `ext_attr`：支持扩展属性，可存储额外元数据；
- `dir_index`：使用哈希索引加速目录查找；
- `flex_bg`：灵活的块组布局，优化大文件系统性能；
- `large_file`/`huge_file`：支持 2TB / 16TB 以上的大文件；
- `quota`/`project`：支持磁盘配额和项目配额。

