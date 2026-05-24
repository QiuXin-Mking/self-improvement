# q
OSD Slow Op 问题的典型根因是什么？
# a
典型根因是 `osd_op_num_shards` 参数配置过低（常见于错误值 5），导致并发处理分片不足，无法应对负载。正确值应设为 16–32（根据磁盘类型调整）。通过 `ceph osd perf` 可观察到 OSD 的慢操作延迟，`ceph osd pool stats` 可查看请求队列堆积情况。

# q
如何从日志定位 RGW 桶索引损坏问题？
# a
当 `s3cmd ls` 或其他列表操作返回空时，说明索引可能损坏。直接使用 `radosgw-admin bucket check --bucket=<bucket_name>` 命令，日志或命令输出会提示索引 inconsistency。若确认损坏，需执行 `radosgw-admin bucket reshard` 或调用数据同步 `radosgw-admin data sync` 重建索引。

# q
解决 Ceph PG inconsistent 问题的标准流程是什么？
# a
1. 用 `ceph pg stat` 或 `ceph pg dump | grep inconsistent` 确认异常 PG ID。
2. 执行 `ceph pg repair <pg_id>` 自动修复合并不一致数据。
3. 若自动修复失败（如位图错误），需使用 `ceph-objectstore-tool --op repair --data-path <osd_path>` 对问题 OSD 进行离线修复，然后重新启动 OSD。修复后进行 `ceph pg scrub <pg_id>` 校验。

# q
Lustre MDT 使用率达到 1000% 的根因是什么？
# a
根因是元数据热点（Metadata Hotspot），即大量并发请求集中在单个 MDT 上，导致其负载远超出能力指标，显示为 1000% 的 CPU 或 I/O 压力。通过 `lfs df -m <mount_point>` 查看 MDT 使用分布，若某一 MDT 远超其他，则需通过目录重分布、使用分布式 MDT 配置（如 DNE）或 `lfs migrate` 迁移文件平衡负载。

# q
BCache 缓存失效导致 OSD Down 的标准排查流程是什么？
# a
1. 使用 `lsblk` 确认 BCache 关联的设备是否掉线。
2. 执行 `bcache-super-show <backing_device>` 检查超级块状态，若显示错误则缓存层损坏。
3. 若缓存设备故障，先执行 `echo 0 > /sys/fs/bcache/<cache_set>/stop` 停用缓存，然后更换故障设备，使用 `make-bcache -C <new_cache_device>` 重建缓存，再重新注册并挂载 backing 设备。最后重启 OSD 服务 `systemctl restart ceph-osd@<id>`。

