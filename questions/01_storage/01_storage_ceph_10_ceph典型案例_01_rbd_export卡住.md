# q
rbd export 卡住时，OSD 日志中可能出现哪些典型错误信息？
# a
日志中常出现以下错误：
- `kernel: VFS: file-max limit 1025500 reached`
- `AdminSocket: do_accept error: '(23) Too many open files in system`

# q
导致 rbd export 卡住的根本原因是什么？
# a
通常是由于 scrub 过程中发现数据不一致（inconsistent），部分 OSD 上的 PG 数据丢失。scrub 结束后 OSD 需要执行 recovery，该过程会大量打开元数据文件，导致系统文件句柄数超过 `file-max` 限制，从而阻塞包括 export 在内的外部请求。

# q
如何临时解决因打开文件数达到上限导致的 rbd export 卡住问题？
# a
通过调整系统最大打开文件数临时恢复，查看和设置方法如下：
```bash
# 查看当前最大文件数限制
cat /proc/sys/fs/file-max
# 查看当前已打开文件数量
cat /proc/sys/fs/file-nr
# 临时增大最大文件数（例如设为 3250）
echo 3250 > /proc/sys/fs/file-max
```

