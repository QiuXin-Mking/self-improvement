# q
在 Lustre 中，`lfs migrate -i` 命令的作用是什么？
# a
用于将指定文件强制迁移到特定索引的 OST（对象存储目标）上。命令格式为 `lfs migrate -i <OST_index> /path/to/file`，其中 `-i` 参数指定目标 OST 的索引号。

# q
如何用脚本批量生成将 40 个测试文件按索引分配到不同 OST 的 `lfs migrate` 命令？
# a
可以使用简单的循环生成，例如 Python 代码：
```python
for i in range(40):
    print(f"lfs migrate -i {i} /mnt/lustre/chenlou/test_ost_perf/test_file_{i}.dat")
```
这会将 `test_file_0.dat` 迁移到 OST 0，`test_file_1.dat` 到 OST 1，依此类推，用于 OST 性能测试。

# q
根据文档，Lustre 可拨测脚本的改进需求中对命令结果日志的管理策略是什么？
# a
- 日志文件必须写入 `/var/log/lustre_test/lustre_command.log`。
- 需要对日志进行每日压缩。
- 保留 7 天内的信息，删除 7 天前的压缩日志。

# q
脚本对读写失败或超时的异常处理有什么要求？
# a
一旦出现读写失败或超时，脚本应直接返回远超常理的时间信息（例如极大值），以便监控系统识别异常。

