# q
如何在 Git 中创建一个新分支？
# a
使用命令 ```git checkout -b <branch-name>```，这会在当前 HEAD 位置创建新分支并立即切换过去。

# q
如何让 Git 忽略指定文件或文件夹？
# a
1. 打开或创建 `.gitignore` 文件  
2. 添加忽略规则，例如忽略 `example_folder` 整个目录，写入：  
```
example_folder/*
```
3. 提交 `.gitignore` 到仓库：  
```
git add .gitignore
git commit -m "Add .gitignore to ignore example_folder"
```

