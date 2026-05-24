# q
扩容后出现slowops的典型根因是什么？
# a
OSD weight配置不均（例如stg-139-18/21的OSD权重约为14.59，而stg-139-19/20的OSD权重约为7.32），触发数据重平衡，导致大量PG处于active+remapped+backfill_wait状态，从而引发slow requests。

# q
如何从ceph命令输出定位扩容导致的slowops问题？
# a
执行 `ceph pg dump | grep backfill_wait` 查看是否存在 `active+remapped+backfill_wait` 状态的PG，例如：
```
2.7f0 ... active+remapped+backfill_wait  2025-01-13T22:40:23.198576+0800 ...
```
再结合 `ceph osd tree` 确认各OSD weight是否明显不均衡。当PG的up set与acting set不一致且长期处于backfill_wait时，说明正在发生数据迁移。

