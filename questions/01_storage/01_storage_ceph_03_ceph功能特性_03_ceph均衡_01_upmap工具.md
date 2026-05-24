# q
使用Ceph upmap工具需要满足哪些前提条件？
# a
- Ceph版本必须为12.2.x及后续版本。
- Ceph client特性至少需支持luminous，以确保client能解读pg-upmap的新映射结构。

# q
如何使用upmap工具计算需要迁移的PG并生成调整命令？
# a
1. 导出当前osdmap：
```bash
ceph osd getmap -o thisosdmap
```
2. 可选：用导出的osdmap测试PG映射情况：
```bash
osdmaptool --test-map-pgs --pool 5 ./thisosdmap
```
3. 计算需要移动的PG并生成upmap命令文件：
```bash
osdmaptool thisosdmap --upmap afterupmap --upmap-pool lab-zone1.rgw.buckets.data --upmap-max 100 --upmap-deviation 2
```
其中 `--upmap-pool` 指定目标存储池，`--upmap-max` 设定最大重映射条目数，`--upmap-deviation` 控制允许的偏差。

# q
如何执行生成的PG upmap迁移命令？
# a
通过source命令应用生成的文件：
```bash
source afterupmap
```
使用 `cat afterupmap` 可以查看具体生成的调整命令，执行后会实际触发PG的重映射迁移。

