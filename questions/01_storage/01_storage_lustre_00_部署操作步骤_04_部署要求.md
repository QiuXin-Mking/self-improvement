# q
MGT（Management Target）对磁盘性能和冗余有什么要求？
# a
MGT 最大仅 100MB，不需要考虑磁盘性能，但数据非常重要，建议使用 RAID1 磁盘。

# q
MDS（Metadata Server）的 IO 特性类似什么？适合使用哪种存储介质？
# a
MDS 的 IO 特性类似数据库，存在大量寻址操作，需要读后改写（read-modify-write），因此推荐使用 NVMe SSD 或高转速的 SATA 盘。

# q
SATA 和 PCIe 接口的 SSD 在性能上有什么主要区别？
# a
SATA SSD 使用串行总线，常见版本包括 SATA I（1.5Gb/s）、SATA II（3Gb/s）和 SATA III（6Gb/s），支持热插拔；PCIe SSD 通过 PCIe 接口提供更高带宽，通常比 SATA SSD 具有更高的读写速度。

