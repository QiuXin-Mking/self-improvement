# q
ceph 判定 osd down 的流程。
# a
1. osd 定期向 mon 发送 MOSDBeacon 消息，mon会捕获所有osd最后的心跳，检查是否超过超时时间，超时就mark down.
2. 其他 OSD 报告失败（check_failures）检查报告者数量，并查询宽限时间，超过报告这数量，超过查询宽限时间就认为down.

# q
ceph 判定 osd down 的流程中，如何避免 交换机20%丢包故障，导致频繁假死，ceph有无特殊设计？
# a
如果 OSD 被标记为 down 但随后恢复（boot_epoch > 0），会更新 laggy_probability 和 laggy_interval
下次判定时会动态增加宽限时间（grace_time）

自动学习网络抖动模式
对频繁“假死”的 OSD 增加容忍度

