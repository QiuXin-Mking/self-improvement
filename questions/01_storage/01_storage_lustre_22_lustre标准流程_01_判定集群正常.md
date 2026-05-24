# q
判定Lustre集群正常需要执行哪些核心检查命令？
# a
执行以下命令并观察其输出：
- `lfs check all`
- `rbd showmapped`
- `df -h`
- `lctl dl`
- `df -h | grep mgt`

# q
检查MGT（管理目标）挂载情况的命令是什么？
# a
```bash
df -h | grep mgt
```

