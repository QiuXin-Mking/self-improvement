# q
将修改添加到Git暂存区的命令是什么？
# a
```
git add .
```
如果未执行此步骤直接提交，会看到 `Changes not staged for commit` 错误。

# q
如何使用命令行直接为Git提交添加描述信息？
# a
使用 `git commit -m "描述信息"` 命令，例如：
```
git commit -m "提交监控脚本和部署脚本"
```

# q
`git review` 命令在代码提交流程中的作用是什么？
# a
`git review` 是将本地提交推送到 Gerrit 代码审查系统的工具，通常用于发起代码审查。可以在命令后指定目标分支，如 `git review v4.1.1` 或 `git review master`。

# q
`git status` 命令的用途是什么？
# a
`git status` 用于显示当前工作目录和暂存区的状态，包括哪些文件已修改、已暂存或未被跟踪。

