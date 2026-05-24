# q
df命令执行挂住（hang）的典型根因是什么？
# a
通常是由于Lustre客户端所挂载的MDT（元数据目标）不存活。可使用 `lfs check all` 确认。根因是客户端无法访问元数据服务器。

# q
如何用strace命令定位df挂住的具体位置？
# a
执行 `strace df` 可观察到进程卡在某个 `stat` 调用上，例如：
```
stat("/mnt/lustre", ^C
```
这说明对Lustre挂载点的stat调用阻塞，表明Lustre客户端无法连接到MDT。

# q
解决df命令因MDT不存活而挂住的标准流程是什么？
# a
先执行 `umount -l /mnt/lustre` 强制卸载挂载点，然后再次执行 `df` 命令即可恢复正常。

# q
mount.lustre进程处于I（不可中断睡眠）状态的典型根因是什么？
# a
MDT与MGS（管理服务）通讯出现故障，导致mount进程持续等待MGS的回应信号，陷入I状态。

# q
解决mount mdt进程因等待MGS而卡在I状态的标准方法是什么？
# a
恢复并挂载MGS服务，确保MDT能与MGS正常通讯；或者等待默认超时（超过600秒）后进程自动退出。

