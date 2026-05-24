# q
Ceph 池从二副本变更为三副本的标准操作流程是什么
# a
使用 `ceph osd pool set` 命令调整副本数，并同步调整 PG 数量。
1. 设置副本数：
   ```bash
   ceph osd pool set testpool size 3
   ```
2. 查看当前 PG 数与 PGP 数：
   ```bash
   ceph osd pool get testpool pg_num
   ceph osd pool get testpool pgp_num
   ```
3. 增大 PG 数（注意：只能增大，不能减小，除非重建 pool）：
   ```bash
   ceph osd pool set testpool pg_num 512
   ceph osd pool set testpool pgp_num 512
   ```
4. 验证数据写入性能：
   ```bash
   rados bench -p testpool 36000 write --no-cleanup -b 1M -t 64
   rados bench -p testpool 18000 write --no-cleanup -b 128M -t 64
   ```

# q
调整 Ceph 池 PG 数量时的核心限制是什么
# a
PG 数只能增大，不能减小，除非重建 pool。此外，应将 `pgp_num` 设置为与 `pg_num` 相同的值，确保数据分布一致。操作示例：
```bash
ceph osd pool set <pool> pg_num 512
ceph osd pool set <pool> pgp_num 512
```

