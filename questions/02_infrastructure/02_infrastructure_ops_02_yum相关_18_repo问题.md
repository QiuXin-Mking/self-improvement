# q
yum 操作时报错 "error was [Errno 2] Local file does not exist" 的典型根因是什么？
# a
yum 源仓库的元数据（repodata）损坏或与实际 RPM 包状态不一致，导致本地文件索引无效。

# q
如何解决 yum 源 "Local file does not exist" 错误？
# a
在 yum 源服务器上删除 `/mnt/yum-iso/repodata` 目录中的旧元数据文件，然后重新生成：
```
createrepo --update /mnt/yum-iso
```

