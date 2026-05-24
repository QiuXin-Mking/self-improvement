# q
如何查看 Ceph radosgw 实例的缓存相关配置（如是否启用、LRU 大小）？
# a
使用以下命令，替换实际的 ASOK 套接字路径：
```bash
ceph daemon /var/run/ceph/client.rgw.ees-stxx-176-4.asok config show | grep rgw_cache_enabled
ceph daemon /var/run/ceph/client.rgw.ees-stxx-176-4.asok config show | grep rgw_cache_lru_size
```

# q
如何动态开启或关闭 Ceph radosgw 的调试日志，以便在问题分析期间获取详细日志？
# a
使用 `debug_rgw` 参数动态调整日志级别：
- 开启调试（设置级别为 20）：
  ```bash
  ceph daemon /var/run/ceph/client.rgw.ees-stxx-176-2.asok config set debug_rgw 20
  ```
- 关闭调试（设置级别为 0）：
  ```bash
  ceph daemon /var/run/ceph/client.rgw.ees-stxx-176-2.asok config set debug_rgw 0
  ```

# q
如何导出 OSD 的历史慢操作信息以辅助定位后端存储性能瓶颈？
# a
使用以下命令，将指定 OSD 的慢操作历史追加到文件：
```bash
ceph daemon osd.114 dump_historic_slow_ops >> daemonperf_osd114
```
该文件可用于进一步分析 OSD 的延迟和慢请求情况。

