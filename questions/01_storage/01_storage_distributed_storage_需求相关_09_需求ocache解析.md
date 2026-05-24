# q
ocache中“回环”指的是什么？
# a
回环是指一次process read产生一个ocache io的处理流程，形成读请求与缓存命中的闭环。

# q
ocache_readlift_cmd_hit_blk函数的核心作用是什么？
# a
用于处理ocache读提升时命令命中block的情况，提升缓存的读命中率。

# q
刷盘过程中ocache如何处理预读命令？
# a
刷盘中，预读命令就不要下去了，从而避免预读命令干扰刷盘操作。

# q
顺序读写超过95在ocache中代表什么含义？
# a
当顺序读写请求占比超过95%时，可能触发ocache的特定优化逻辑（如调整读提升或预读策略）。

