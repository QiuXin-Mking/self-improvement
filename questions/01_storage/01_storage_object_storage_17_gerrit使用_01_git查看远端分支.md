# q
如何查看远程仓库的所有分支？
# a
使用命令 `git branch -r` 查看所有远程分支，或使用 `git branch -a` 查看本地和远程全部分支。输出中会出现类似 `origin/v4.1` 的远程分支名。

# q
如何创建一个新的本地分支并立即追踪指定的远程分支？
# a
使用命令 `git checkout -b <本地分支名> <远程分支名>`，例如：
```
git checkout -b v4.1 origin/v4.1
```
该命令会创建本地分支 `v4.1`，切换到该分支，并自动设置上游为 `origin/v4.1`。

# q
如果本地分支已存在，如何将其设置为追踪远程分支？
# a
使用 `--set-upstream-to` 选项（较新 Git 版本）：
```
git branch --set-upstream-to=origin/v4.1 v4.1
```
或旧版本使用 `git branch -u origin/v4.1 v4.1`。

# q
在本地分支追踪远程分支后，如何获取最新的远程更改？
# a
可使用 `git pull` 拉取并合并远程分支的最新更改，或者只使用 `git fetch` 获取更改而不自动合并。

