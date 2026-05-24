# q
at命令的作用是什么？
# a
at命令用于设置一次性定时计划任务，任务执行一次后不再重复执行。

# q
atq和atrm命令的功能分别是什么？
# a
atq用于列出用户的计划任务（超级用户将列出所有用户的任务），输出格式为：作业号、日期、小时、队列和用户名；atrm用于根据作业号（job number）删除at任务。

# q
如何设置一分钟后执行指定文件（如`/tmp/date`）中的命令？
# a
使用命令 `at -f /tmp/date now +1 minutes`，其中 `-f` 选项表示从文件中读取任务，`now +1 minutes` 是相对时间表示一分钟之后执行。

# q
使用at命令的必要前提是什么？
# a
必须保证atd守护进程正在运行，可通过 `ps -ef | grep atd` 检查。

# q
at命令的用户权限控制规则是什么？
# a
由 `/etc/at.allow` 和 `/etc/at.deny` 文件控制：
- 若 `at.allow` 存在，仅其中列出的用户可使用at；
- 若 `at.allow` 不存在，`at.deny` 存在，则不在 `at.deny` 中的用户可使用at；
- 若两者都不存在，只有root用户可使用at；
- 若 `at.deny` 为空（默认配置），则所有用户均可使用at；
- `at.allow` 的优先级高于 `at.deny`。

