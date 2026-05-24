# q
如何查看本地分支？
# a
```bash
git branch
```

# q
如何查看远端分支？
# a
```bash
git branch -r
```

# q
如何克隆一个远程仓库？
# a
```bash
git clone http://xxxx
```

# q
如何创建并切换到一个追踪远端分支的本地分支（例如 tracking origin/octopus）？
# a
```bash
git checkout -b octopus origin/octopus
```

# q
如何克隆指定分支到本地并重命名本地目录？
# a
```bash
git clone --branch octopus https://github.com/ceph/ceph.git ceph-octopus
```

