# q
在Lustre文件系统中，为什么建议使用 noatime 挂载选项，并如何设置？
# a
每次访问inode都会更新访问时间（accesstime），这会产生大量的小写IO，影响性能。使用 `noatime` 可禁止此更新，挂载命令为：
```
mount -t lustre -o noatime <远端服务> <挂载点>
```

# q
Samba 的 op_lock 锁存在什么潜在问题？
# a
op_lock 锁可能因缓存问题导致数据损坏，或错误地报告“文件正在被使用”。

