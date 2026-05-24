# q
如何创建一个新分支并立即切换到该分支？
# a
使用 `git checkout -b <branch-name>` 命令，例如：
```
git checkout -b feature-x
```
此命令会基于当前所在的提交创建新分支，并自动切换到新分支。

# q
如何让 Git 忽略某个目录下的所有文件？
# a
在仓库根目录创建或编辑 `.gitignore` 文件，添加要忽略的模式，例如忽略 `example_folder` 目录下的所有内容：
```
example_folder/*
```
然后将 `.gitignore` 文件提交到仓库：
```
git add .gitignore
git commit -m "Add .gitignore to ignore example_folder"
```
注意：`.gitignore` 只对未跟踪的文件生效，之前已被跟踪的文件不受影响，需要先取消跟踪。

