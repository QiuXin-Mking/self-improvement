# q
本次测试中，128个job 1M顺序读时，直接部署与修改nscheds参数（小/大）的Lustre性能有何差异？
# a
直接部署：BW=2992MiB/s（约2.92GiB/s），IOPS=2991。
修改为小nscheds：BW=10.9GiB/s，IOPS=11.1k，性能最高。
修改为大nscheds：BW=9.96GiB/s，IOPS=10.2k。
调整nscheds显著提升读带宽，小nscheds配置收益最大。

# q
在不修改nscheds仅重启Lustre节点后，FIO顺序读性能表现如何？
# a
只重启测试结果：BW=9668MiB/s（约9.44GiB/s），IOPS=9667。相比直接部署的2992MiB/s有明显提升，但仍低于修改小nscheds的10.9GiB/s，说明重启有一定作用但nscheds调优更为关键。

# q
修改小nscheds后，FIO读延迟的百分位分布（P50、P99）与直接部署对比有何变化？
# a
直接部署：P50≈1972ms，P99≈4178ms，延迟极高。
修改小nscheds后：P50≈174ms（大幅降低），P95≈4329ms，P99≈6745ms。中低位延迟改善显著，但高百分位延迟依然较高，呈现长尾分布。

