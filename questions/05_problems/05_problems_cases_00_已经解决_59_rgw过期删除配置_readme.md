# q
RGW生命周期过期删除规则不生效，日志中无delete操作记录的典型根因是什么？
# a
典型根因是生命周期执行进程卡住未处理工作项。常见情况：
- lc_pool中有待处理任务但进程停滞，需要手动触发`radosgw-admin lc process`（该命令为全集群校验）。
- 桶的索引分片数过多（例如`bucket_index_max_shards`为11），导致扫描效率下降，处理超时。
- 碎片的过期规则（如碎片删除2天）与生命周期策略的并行处理冲突，导致upload桶状态停留。

# q
如何通过命令行定位RGW生命周期过期规则是否卡住？
# a
按以下步骤操作：
1. 查看所有桶的生命周期状态：`radosgw-admin lc list`
2. 手动触发全集群生命周期处理：`radosgw-admin lc process`
3. 检查lc_pool中是否有积压的工作项（根据zone配置找到lc_pool名称，如`.product.rgw.log:lc`），执行：
   ```bash
   rados -p .product.rgw.log --namespace=lc ls
   ```
4. 观察日志中是否产生delete操作记录，确认规则是否开始生效。

# q
解决RGW生命周期过期删除规则不生效的标准流程是什么？
# a
标准排查与修复流程：
1. **确认配置**：`radosgw-admin lc list`，确保桶的过期规则（如7天删除、碎片2天删除）已正确应用。
2. **手动触发**：执行`radosgw-admin lc process`在全集群强制启动一次生命周期扫描，观察状态变化。
3. **检查工作队列**：进入lc_pool（示例`.product.rgw.log:lc`），用`rados -p <lc_pool> --namespace=lc ls`列出待处理对象，确认是否有阻塞。
4. **修复分片异常**：若仍无效，执行`radosgw-admin lc reshard fix`修复生命周期索引分片问题。
5. **监控日志**：查看rgw日志，确认delete字段是否出现。
6. 若桶分片数过多（如11），评估是否需要调整`bucket_index_max_shards`或拆分桶以提升处理效率。

