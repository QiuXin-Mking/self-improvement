# q
为什么在 Git 中需要设置用户名和邮箱地址？
# a
这两个信息会被用于提交记录中，标识是谁做了这次提交。Git 要求设置用户名和邮箱后才能正常提交。

# q
如何全局设置 Git 用户名和邮箱，使所有仓库都生效？
# a
使用 `--global` 选项执行：
```bash
git config --global user.name "你的用户名"
git config --global user.email "你的邮箱地址"
```

# q
如何仅对当前仓库设置独立的用户名和邮箱？
# a
在仓库目录下省略 `--global` 选项执行：
```bash
git config user.name "你的用户名"
git config user.email "你的邮箱地址"
```
这样设置仅对该仓库有效，不会影响其他仓库。

# q
如果已经提交但发现作者信息错误，如何修正最近一次提交的用户名和邮箱？
# a
使用 `git commit --amend --author` 命令修改：
```bash
git commit --amend --author="Author Name <email@example.com>"
```
注意：此操作会改变提交哈希值；对于已推送到远程仓库的提交，修改更复杂，需谨慎处理，避免扰乱历史记录。

