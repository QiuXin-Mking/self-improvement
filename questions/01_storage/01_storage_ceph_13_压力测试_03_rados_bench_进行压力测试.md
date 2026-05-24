# q
如何使用 rados bench 对 Ceph 存储池进行写性能压力测试？
# a
先创建一个测试存储池，然后使用 `rados bench` 命令执行写测试。常用命令为：
```
ceph osd pool create testpool 128
rados -p testpool bench 300 write --no-cleanup
```
`300` 表示测试持续300秒，`--no-cleanup` 表示测试结束后保留数据，便于后续读测试。

# q
如何使用 rados bench 对已写入数据的存储池进行读性能测试？
# a
在写测试（使用 `--no-cleanup` 保留数据）后，运行顺序读测试：
```
rados -p testpool bench 300 seq
```
可添加参数 `--object-size` 指定对象大小，`--total` 指定读写总数据量，例如：
```
rados -p testpool bench 300 seq --object-size 4M --total 124M
```

# q
使用 rados bench 测试后，如何彻底清理测试数据与存储池？
# a
先清理测试数据，再删除存储池：
```
rados -p testpool cleanup
ceph osd pool delete testpool testpool --yes-i-really-really-mean-it
```

# q
在压力测试过程中，常用哪些命令监控 OSD 性能与状态？
# a
常用监控命令：
```
ceph -s          # 查看集群整体状态
ceph osd perf    # 查看所有 OSD 的延迟和吞吐
ceph osd df      # 查看每个 OSD 的磁盘使用情况
```

