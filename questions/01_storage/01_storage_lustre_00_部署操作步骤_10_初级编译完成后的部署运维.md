# q
如何格式化并挂载一个独立的 Lustre MGS 服务？
# a
使用 `mkfs.lustre` 格式化块设备，指定文件系统名称和 `--mgs` 选项，然后直接 mount：
```bash
mkfs.lustre --fsname=lustre1 --mgs --reformat /dev/vdb
mount -t lustre /dev/vdb /mnt/lustre-mgs
```

# q
如何部署一个独立的 Lustre MDS 并指向已存在的 MGS？
# a
格式化时需指定文件系统名称、`--mdt` 类型、`--mgsnode` 以及唯一的 `--index`（从 0 开始）：
```bash
mkfs.lustre --fsname=lustre1 --mdt --mgsnode=172.31.0.21@tcp --index=0 --reformat /dev/vdb
mount -t lustre /dev/vdb /mnt/lustre-mds
```

# q
Lustre OSS 的 `read_cache_enable` 和 `writethrough_cache_enable` 参数有何作用？如何查看和修改？
# a
`read_cache_enable` 控制 OSS 是否启用读缓存，`writethrough_cache_enable` 控制是否启用写穿透缓存。查看和修改方法：
```bash
# 查看读缓存状态
lctl get_param osd-*.*.read_cache_enable
# 禁用读缓存
lctl set_param osd-*.*.read_cache_enable=0
# 启用读缓存
lctl set_param osd-*.*.read_cache_enable=1

# 查看写穿透缓存状态
lctl get_param osd-*.*.writethrough_cache_enable
# 设定值
lctl set_param osd-*.*.writethrough_cache_enable=0
```

# q
如何通过调整 MDS 参数优化 Lustre 大目录性能？
# a
可以设置 `mdt.*.dir_split_count` 和 `mdt.*.dir_split_delta` 参数来控制目录分裂行为，例如：
```bash
ansible lustre -m shell -a "lctl set_param mdt.*.dir_split_count=5000"
ansible lustre -m shell -a "lctl set_param mdt.*.dir_split_delta=2"
```
这可以提升大量文件在同一目录下的操作性能。

# q
对 Lustre 进行网络性能调优时，通常需要调整哪些系统参数？
# a
主要调整 TCP 缓冲区大小和网卡 Ring Buffer 大小：
```bash
# TCP 收包缓冲区最大值
sysctl -w net.ipv4.tcp_rmem="4096 8738000 629145600"
# TCP 发包缓冲区最大值
sysctl -w net.ipv4.tcp_wmem="4096 1638400 419430400"
# 将网卡 tx/rx Ring Buffer 设置为 8192
ethtool -G p259p1 tx 8192 rx 8192
ethtool -G p259p2 tx 8192 rx 8192
ethtool -G p769p1 tx 8192 rx 8192
ethtool -G p769p2 tx 8192 rx 8192
```

