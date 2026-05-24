# q
在 Ceph 环境中使用 vdbench 进行性能测试时，如何找到要测试的块设备路径？
# a
通过 `ll /dev/ceph-*/osd-*` 命令查看，会发现 OSD 的逻辑卷被映射到 `/dev/dm-X` 设备，例如 `/dev/dm-2`，该设备用于 vdbench 脚本中的 `lun` 参数。

# q
示例 vdbench 测试脚本 `raw_test` 的核心配置项有哪些？
# a
脚本包含三部分配置：
- `sd=sd1,lun=/dev/dm-2,openflags=o_direct`：定义存储设备并设置直接 I/O 标志
- `wd=wd1,sd=sd1,xfersize=4k,rdpct=75`：定义工作负载，传输大小 4K，读请求占比 75%
- `rd=run1,wd=wd*,iorate=100,elapsed=7200,interval=1`：定义运行参数，限制 IOPS 为 100，运行 7200 秒，每秒输出一次统计

# q
如何执行 vdbench 测试并提取平均性能指标？
# a
- 测试前验证环境：`./vdbench -t`
- 执行脚本：`./vdbench -f raw_test`
- 提取结果：`cat output/summary.html | grep avg`

