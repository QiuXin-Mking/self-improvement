# q
df命令在Lustre客户端挂住时的典型根因是什么？如何排查？
# a
典型根因是MDT（元数据目标）不存活。可先执行 `strace df` 观察卡在哪个挂载点，通常为 `/mnt/lustre`。随后使用 `lfs check all` 确认MDT状态，若异常则执行 `umount -l /mnt/lustre` 解除df阻塞。

# q
mount mdt进程处于I状态（不可中断睡眠）如何解决？
# a
原因通常是mdt与mgt（管理目标）无法通信，导致mdt一直等待mgt信号。解决方案：先将mgt挂载上，或等待超过600秒超时。若急需恢复，可先挂载mgt。

