# q
如何理解 `git add .` 在代码提交流程中的作用？
# a
`git add .` 将当前目录下所有修改的文件添加到暂存区（staged），准备提交。如果跳过此步骤直接执行 `git commit`，会提示 `Changes not staged for commit` 错误，因为修改尚未被追踪到暂存区。

# q
`git review v4.1.1` 命令的作用是什么？
# a
`git review v4.1.1` 用于将本地提交推送到 Gerrit 上进行代码评审，并指定目标分支为 `v4.1.1`。它简化了与 Gerrit 的交互，自动创建或更新评审请求。

# q
当执行 `git commit` 后 `git status` 显示什么状态？
# a
`git status` 用于查看工作区和暂存区的状态。在成功执行 `git commit` 后，`git status` 会显示“nothing to commit, working tree clean”，表示所有修改已提交，工作区干净。

# q
在提交前如果忘记执行 `git add`，会出现什么现象？
# a
如果未执行 `git add` 而直接 `git commit`，Git 会提示 `Changes not staged for commit`，表明修改未暂存，无法继续提交，需要先通过 `git add` 将文件添加到暂存区。

