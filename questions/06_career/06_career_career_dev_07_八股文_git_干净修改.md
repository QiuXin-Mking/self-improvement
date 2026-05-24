# q
发现有一个不要的commit记录，已经合并在谦哥的修改里面了，怎么删除掉
# a
1. 确保本地工作区干净（没有未提交的改动），然后切换到其他分支：
```bash
git status
```
2. 删除本地分支：
```bash
git branch -D feature/login
```
3. 从远程重新拉取并建立同名分支：
```bash
git checkout -b feature/login origin/feature/login
```

