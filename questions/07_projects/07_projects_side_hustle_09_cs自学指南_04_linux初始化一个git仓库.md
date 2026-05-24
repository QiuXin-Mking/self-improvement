# q
如何初始化一个新的Git仓库？
# a
使用 `git init` 命令在当前目录创建一个空的Git仓库。

# q
如何将当前目录的所有文件添加到Git暂存区？
# a
使用 `git add .` 命令会将当前目录及其子目录下的所有文件添加到暂存区。

# q
如何配置Git的全局用户名和邮箱？
# a
使用以下命令分别设置全局用户名与邮箱：
```bash
git config --global user.email "132203516@qq.com"
git config --global user.name "qiuxin"
```

# q
如何提交暂存区的更改并添加提交信息？
# a
使用 `git commit -m "提交信息"` 命令，例如：
```bash
git commit -m "first_commit_2.15"
```

