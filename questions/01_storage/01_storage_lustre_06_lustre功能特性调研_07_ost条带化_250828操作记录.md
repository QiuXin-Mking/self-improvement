# q
Lustre 中用于设置条带化（stripe）的命令是什么？关键参数有哪些？
# a
使用 `lfs setstripe` 命令。关键参数：
- `-c` 或 `--count`：指定条带计数（stripe count），即文件分布的 OST 数量。
- `-S` 或 `--size`：指定条带大小（stripe size），如 `-S 1M` 表示 1 MiB。
示例：`lfs setstripe -c 12 -S 1M /mnt/lustre`

# q
如何在不删除数据的情况下更改 Lustre 文件或目录的条带配置？
# a
使用 `lfs migrate` 命令，该命令会将数据从旧条带布局迁移到新布局。例如：
`lfs migrate -c 12 -S 1M /mnt/lustre/$(hostname)/Fio*`
根据记录，将 16 GiB 数据从 12 条带迁移为 1 条带耗时约 13 分钟，吞吐约 1.2 GB/分钟。

# q
查看 Lustre 文件或目录的条带布局信息用什么命令？
# a
使用 `lfs getstripe` 命令。示例：
`lfs getstripe /mnt/lustre/$(hostname)_c12s1/Fio.0.0`
输出会包含 obdidx、objid、group 等布局细节。

# q
条带化数量对 Lustre 读写性能有何影响？
# a
条带数量显著影响 I/O 吞吐。在 fio 测试中，将条带从 1 增加到 4：
- 顺序读从约 116 MiB/s 提升至约 460 MiB/s（约 4 倍）
- 顺序写从约 115 MiB/s 提升至约 428 MiB/s（约 3.7 倍）
多 OST 并行可以有效提升聚合带宽。

# q
在 Lustre 性能测试中，fio 的 numjobs 参数应如何设置以匹配条带化？
# a
建议将 `-numjobs` 设为 OST 数量（或条带计数）的整数倍。例如测试教训指出：12 个 OST 时应使用 12 的倍数（如 60），以均匀打满各 OST，避免单线程成为瓶颈。

