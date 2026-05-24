# q
Ceph Monitor 的 debug 机制涉及哪些核心配置参数？
# a
```plaintext
debug_ms = 1
debug_mon = 20
debug_paxos = 15
```

# q
如何设置 Ceph Monitor 的调试级别以获得详细日志？
# a
在 ceph.conf 中配置以下参数并重启 monitor 服务：
```ini
[mon]
debug_ms = 1
debug_mon = 20
debug_paxos = 15
```

