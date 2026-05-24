# q
如何通过配置减缓 BlueStore 缓存 trim 行为过于激进导致的资源释放冲突？
# a
可以调整以下参数来减少频繁的 trim 行为，间接减少并发释放和锁竞争：
- `bluestore_cache_trim_interval`：减小该值可减缓 trim 频率（例如设置为 `0.0001` 或 `5`）。
- `bluestore_cache_trim_max_skip_pinned_adjust` 等相关参数。
此外，`bluestore_warn_on_transaction_too_big` 等参数可对“大事务”进行警告，辅助发现相关问题。

# q
`bluestore_cache_trim_max_skip_pinned` 参数目前是否生效？
# a
该参数在当前代码中已废弃，没有 `_conf` 的配置，虽然可以通过 `ceph config get osd bluestore_cache_trim_max_skip_pinned` 查到值（如 64），但代码中已不再使用。

# q
如何使用命令行查看并临时修改 BlueStore 的缓存大小配置？
# a
查看缓存大小（以 SSD 和 HDD 为例）：
```
ceph config get osd bluestore_cache_size_ssd
ceph config get osd bluestore_cache_size_hdd
```
临时设置为 0：
```
ceph config set osd bluestore_cache_size_ssd 0
ceph config set osd bluestore_cache_size_hdd 0
```

