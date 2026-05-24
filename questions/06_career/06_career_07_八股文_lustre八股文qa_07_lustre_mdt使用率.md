# q
如何用 top、grep、awk 命令查询 Lustre 所有 MDT 线程的 CPU 使用率总和？
# a
使用以下命令：
```bash
top -bn1 | grep mdt | awk 'BEGIN{s=0}{s+=$9}END{print s}'
```
其中 `top -bn1` 输出一次系统进程信息，`grep mdt` 筛选出 MDT 线程，`awk` 累加第9列（CPU使用率）并打印总和。

