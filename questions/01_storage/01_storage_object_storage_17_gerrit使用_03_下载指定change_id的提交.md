# q
如何使用git review下载指定change_id的提交？
# a
执行命令 `git review -d $change_id`，该命令会基于指定的change创建一个新的本地分支，并自动跟踪该修改。

