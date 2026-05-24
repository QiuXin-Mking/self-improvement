# q
修改Ceph池的pg_num触发PG分裂时，导致部分PG进入down状态的典型根因是什么？
# a
根因通常是PG分裂过程中OSD负载突增，bluestore缓存来不及回收（trim）或缓存空间不足，导致IO延迟过高或超时，从而触发PG被标记为down。可通过调整`bluestore_cache_trim_interval`和`bluestore_cache_size_ssd`来缓解。

# q
在执行`ceph osd pool set <pool> pg_num <num>`扩容PG数量时，如何通过调整配置避免PG down？
# a
可临时调低bluestore缓存trim间隔并增加缓存大小，例如：
```
ceph config set osd bluestore_cache_trim_interval 0.005
ceph config set osd bluestore_cache_size_ssd 500
```
测试表明：从200增至300时，`bluestore_cache_trim_interval`设为0.005可避免PG down；从300增至600时可能需要更低的值（如0.0005），但必须结合集群实际压力调整，避免过度降低影响其他IO。

# q
如何从日志定位PG分裂时引起的PG down问题？
# a
1. 在操作前开启OSD和bluestore调试日志：
```
ceph tell osd.* config set debug_osd 20
ceph tell osd.* config set debug_bluestore 20
```
2. 执行`ceph osd pool set <pool> pg_num <new_num>`触发分裂。
3. 观察集群中PG down数量的变化（如`ceph -s`）。
4. 检查OSD日志中与慢IO、超时、trim相关记录，重点关注bluestore缓存相关报错，确认是否因缓存压力导致PG被标记为down。操作完成后恢复日志级别：
```
ceph tell osd.* config set debug_osd 0
ceph tell osd.* config set debug_bluestore 0
```

