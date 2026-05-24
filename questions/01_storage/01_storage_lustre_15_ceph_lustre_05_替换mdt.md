# q
在Lustre中，恢复超时(recovery timed out)后系统会如何处理客户端？
# a
系统会记录 `recovery is timed out, evict stale exports`，随后断开（disconnecting）过时客户端连接并驱逐它们。例如日志中显示 `disconnecting 4 stale clients`，最终恢复完成后可能报告已恢复的客户端数和被驱逐的客户端数，如 `of 10 clients 6 recovered and 4 were evicted`。

# q
如何查看Lustre文件或目录的条带分布信息？
# a
使用 `lfs getstripe` 命令，例如：
```bash
lfs getstripe /mnt/lustre/test_174/vdb.1_1.dir
```

# q
如何手动挂载一个Lustre OST设备（例如基于RBD的块设备）？
# a
使用 `mount -t lustre` 指定设备路径和挂载点，例如：
```bash
mount -t lustre /dev/rbd1 /data/lustre_ost_1/
```

# q
在替换Lustre MDT后，恢复过程可能报告哪些客户端状态？
# a
恢复完成后，日志可能报告类似 `Recovery over after 2:30, of 10 clients 6 recovered and 4 were evicted` 的信息，表明部分客户端恢复成功，部分因超时或状态陈旧被驱逐。

