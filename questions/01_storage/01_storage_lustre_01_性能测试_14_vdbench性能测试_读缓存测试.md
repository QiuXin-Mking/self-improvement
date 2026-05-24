# q
在Lustre读操作过程中，为何节点内存的buff/cache会大幅增长？
# a
Lustre客户端执行读操作时，会将读取的数据缓存到Linux页面缓存（page cache）中，即free命令中的buff/cache部分。测试中128K读期间，buff/cache从不足3GB飙升到约250~260GB，反映了将近250GB数据被缓存到内存，以加速后续访问。这是标准Linux内核的缓存行为，并非Lustre特有。

# q
在Linux系统中，如何手动丢弃页面缓存来释放内存？具体命令是什么？
# a
使用命令：
```
echo 3 > /proc/sys/vm/drop_caches
```
其中参数3表示同时清理page cache、dentry和inode缓存。示例中执行`ansible lustre -m shell -a "echo 3 > /proc/sys/vm/drop_caches"`后，所有节点的buff/cache从约270GB回落至数百MB。

# q
在示例测试中，128K写操作前后内存buff/cache的变化情况如何？
# a
128K写操作前，各节点buff/cache仅为约500MB~900MB；写操作后，buff/cache上升至约2.8~3.2GB，说明写操作也会产生少量页面缓存，但远低于读操作产生的缓存量。

