# q
如何删除未使用的 Docker 镜像？
# a
使用 `docker image prune` 命令；添加 `-a` 选项会删除所有未使用的镜像，而不仅仅是悬空镜像。

# q
如何列出所有悬空镜像（dangling images）？
# a
```bash
docker images -f "dangling=true"
```

# q
如何删除所有悬空镜像？
# a
```bash
docker rmi $(docker images -f "dangling=true" -q)
```
也可以使用 `docker image prune` 删除悬空镜像。

