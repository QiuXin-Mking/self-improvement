# q
修改PG数量后导致PG down的典型osdmap特征是什么？
# a
osdmap中目标pool的`pg_num`与`pgp_num`严重不一致，且存在大量`pg_temp`映射。例如输出中pool 4的`pg_num 750`、`pgp_num 113`、`pgp_num_target 750`，并伴随多个`pg_temp`条目（如`pg_temp 4.65 [3,6,5]`），表明PG正在重新映射但尚未完成，极易导致PG长时间处于down状态。

# q
如何通过osdmap查看OSD的运行状态？
# a
使用`ceph osd getmap -o osdmap`导出osdmap，再用`osdmaptool osdmap --print`解析输出。每个osd条目会显示`up`/`down`、`in`/`out`状态以及`up_from`、`down_at`、`last_clean_interval`。例如`osd.11 down in`表示OSD 11被标记为down但仍被集群认为in，`down_at 765`和`up_from 684`可帮助判断故障持续时长。

# q
当osdmap显示OSD为down但进程仍在时，如何进一步排查？
# a
在对应节点上执行`netstat -ntlp | grep <端口>`和`ps aux | grep <进程ID>`。如osd.11对应的PID 117677仍存在，且端口6825/6826/6827处于LISTEN状态，说明进程未退出，可能是心跳超时或网络问题导致被标记down。应检查网络连通性及`/var/log/ceph/ceph-osd.<id>.log`中的心跳日志。

