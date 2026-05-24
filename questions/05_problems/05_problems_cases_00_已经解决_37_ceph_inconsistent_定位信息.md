# q
如何识别 Ceph PG 发生 inconsistent 并需要修复的状态？
# a
通过 `ceph pg ls inconsistent` 查看，PG 状态会显示为 `active+clean+scrubbing+inconsistent+failed_repair` 等组合，日志中会出现 `scrub` 相关错误和 `repop` 反复重放。

# q
Ceph PG inconsistent 的典型根因是什么？
# a
通常是因为 PG 内对象副本不一致，可能由存储介质错误、OSD 故障或网络闪断导致。日志中会显示特定对象（如 `rbd_data.1222247a9395d3.00000000000018b7`）的 scrub 错误，以及 `might_have_unfound` 的 OSD（如 `osd.1`、`osd.13`）。

# q
解决 Ceph PG inconsistent 的标准流程是什么？
# a
1. 先尝试 `ceph pg repair <pgid>`（如 `ceph pg repair 4.d`）；
2. 若修复失败，可对不一致对象涉及的 OSD 做 `ceph osd pg-upmap-items` 重新映射（如 `ceph osd pg-upmap-items 4.d 1 2` 和 `ceph osd pg-upmap-items 4.d 13 12`）；
3. 再次执行 `ceph pg repair 4.d` 和 `ceph pg 4.d deep_scrub`；
4. 观察 PG 状态是否恢复正常。

# q
如何从日志中定位 PG inconsistent 的具体对象？
# a
通过 `ceph pg <pgid> query` 可以查看 `scrubber.start`、`scrubber.end` 字段，其中包含对象标识（如 `4:b671c3ed:::rbd_data.7b70996d74511d.00000000000018c2:16`）。也可使用 `grep -i "inconsistent" /var/log/ceph/ceph-osd.<id>.log` 和 `grep -i "scrub" ... | grep -i error` 搜索日志。

