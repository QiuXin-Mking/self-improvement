# q
crontab 的五个时间字段分别代表什么含义？顺序如何？
# a
crontab 时间字段从左到右依次为：分钟（0-59）、小时（0-23）、天（1-31）、月（1-12）、星期（0-6，0 表示星期天）。每个字段用空格分隔。

# q
如何在 crontab 中每 5 分钟执行一次命令？
# a
使用 `*/5 * * * * command`，例如 `*/5 * * * * ls`。

# q
执行 crontab 任务时，命令路径有什么特殊要求？如果需要在特定目录运行脚本，如何正确配置？
# a
crontab 中只能使用绝对路径，不存在相对路径。如果要在特定目录执行脚本，可以先使用 `cd` 进入绝对目录，再执行脚本，例如：
```sh
0 * * * * cd /data/sync-data/mdd; /data/GoProjects/bin/go-elastic-index -config=mdd.toml >> mdd.log 2>&1
```
或者所有路径都写完整的绝对路径。

