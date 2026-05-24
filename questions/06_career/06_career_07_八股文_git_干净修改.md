# q
如何让本地分支与远程完全一致，丢弃本地多余的 commit？
# a
```bash
# 1. 确保工作区干净
git status
# 2. 切换到其他分支（例如 master）
git checkout master
# 3. 强制删除目标本地分支
git branch -D feature/login
# 4. 从远程重新创建同名本地分支
git checkout -b feature/login origin/feature/login
```

# q
为什么删除本地分支前必须切换到其他分支？
# a
因为 Git 不允许删除当前所在分支，必须先移动到其他分支才能删除目标分支。

