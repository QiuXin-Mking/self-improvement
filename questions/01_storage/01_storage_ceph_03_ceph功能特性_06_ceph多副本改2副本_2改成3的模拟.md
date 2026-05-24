# q
在Ceph中模拟存储池从2副本改为3副本时，如何调整恢复参数以减少对业务的影响？
# a
通过设置以下OSD配置参数：
- `osd_max_backfills`：限制并发回填操作数，例如设为1
- `osd_recovery_op_priority`：降低恢复操作优先级，例如设为3
- `osd_recovery_sleep`：在恢复操作之间引入休眠（秒），例如设为1
- `osd_recovery_max_active`：限制活跃恢复请求数，例如设为1

设置命令示例：
```
ceph config set osd osd_max_backfills 1
ceph config set osd osd_recovery_op_priority 3
ceph config set osd osd_recovery_sleep 1
ceph config set osd osd_recovery_max_active 1
```

# q
如何临时暂停Ceph集群的恢复和回填？
# a
使用以下命令暂停：
- 只暂停恢复：`ceph osd set norecover`
- 只暂停回填：`ceph osd set nobackfill`

# q
在模拟2副本转3副本时，创建测试Pool并执行副本数变更的命令是什么？
# a
创建Pool并初始化配置：
```
ceph osd pool create 2to3 64
ceph osd pool set 2to3 size 2
ceph osd pool set 2to3 pg_num 512
ceph osd pool set 2to3 pgp_num 512
```
执行副本数从2改为3：
```
ceph osd pool set 2to3 size 3
```

# q
如何使用`rados bench`对特定Pool进行读写压测？
# a
写测试（60秒，不清理数据）：
```
rados bench -p 2to3 60 write --no-cleanup
```
顺序读测试（60秒）：
```
rados bench -p 2to3 60 seq
```

