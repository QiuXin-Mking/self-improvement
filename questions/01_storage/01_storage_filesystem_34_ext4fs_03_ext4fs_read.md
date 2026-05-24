# q
inode 是什么
# a
inode​ 是文件的“元数据载体”，描述文件的属性和数据存储位置；

# q
entry是什么
# a
directory entry​ 是目录的“索引项”，负责将文件名映射到对应的 inode。录 entry 是“文件名→inode 编号”的桥梁

# q
ext4 支持两种目录存储格式，直接影响 entry 的组织方式：是哪两种？
# a
1. 线性目录（Linear Directory）
2. Hash 树目录（Hash Tree Directory，ext4 特有优化）

# q
ext3 文件系统 和 ext4 文件系统的差异？
# a
最大文件系统约 16TB（需 64 位系统），单文件最大 2TB（因 inode 中块指针是 32 位，最多寻址 2³² 个块，若块大小 4KB 则单文件 16TB？实际因实现限制通常为 2TB）

