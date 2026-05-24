# q
提交到Gerrit时缺少Change-Id会报什么错误？
# a
报错信息：`remote: ERROR: missing Change-Id in commit message footer`

# q
为什么Gerrit要求每个提交都必须包含Change-Id？
# a
Change-Id（如`Change-Id: Iabc...`）是Gerrit用来唯一追踪一次代码评审过程的标识，需要放在commit message的最后（footer部分）。

# q
如何解决本地提交缺少Change-Id的问题？
# a
1. 下载并安装commit-msg hook：
```bash
gitdir=$(git rev-parse --git-dir); scp -p -P 29418 qiux1@10.3.196.2:hooks/commit-msg ${gitdir}/hooks/
```
2. 使用该hook修正最后一次提交：
```bash
git commit --amend
```

