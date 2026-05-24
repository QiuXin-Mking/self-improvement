# q
如何获取Lustre客户端的cache使用量？
# a
使用命令 `cat /sys/kernel/debug/lustre/llite/*/max_cached_mb | grep '^used_mb'`，单位为 MB。

# q
如何获取Lustre客户端的read ahead使用量？
# a
使用命令 `cat /sys/kernel/debug/lustre/llite/*/max_cached_mb | grep used_read_ahead_mb`，单位为 MB。

# q
如何使用`lctl`命令查看Lustre客户端的`checksum_pages`参数？
# a
执行 `lctl get_param llite.*.checksum_pages`。

# q
如何使用`lctl`命令查看Lustre客户端的`enable_filename_encryption`参数？
# a
执行 `lctl get_param llite.*.enable_filename_encryption`。

