# q
Ceph MON节点空间不足时会看到怎样的警告信息？
# a
运行 `ceph -s` 时会显示类似 `mon ees-0-2 is low on available space` 的警告。

# q
如何排查Ceph MON节点的磁盘空间使用情况？
# a
先用 `du -sh /var/lib/ceph/mon/*` 查看 monitor 数据目录的占用情况，再用 `sudo du -ah / | sort -rh | head -n 20` 找出整个系统中最大的 20 个文件或目录，定位占用空间的大户（通常是日志类文件）。

# q
`du -ah / | sort -rh | head -n 20` 中各参数的含义是什么？
# a
- `du -ah /`：`du` 计算磁盘使用，`-a` 显示所有文件和目录，`-h` 以人类可读的单位（K/M/G）显示，`/` 从根目录遍历。  
- `sort -rh`：`sort` 排序，`-r` 降序，`-h` 正确处理带单位的数值。  
- `head -n 20`：只输出前 20 行结果。  
该命令能快速找出占用空间最大的文件或目录。

