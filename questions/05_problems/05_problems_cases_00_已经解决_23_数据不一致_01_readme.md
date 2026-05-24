# q
如何从日志中发现Ceph PG数据不一致问题？
# a
在OSD日志或mon日志中搜索`scrub`关键字。典型错误日志示例：
```
log_channel(cluster) log [ERR] : 10.3ffs0 scrub : stat mismatch, got 408032/408033 objects, 0/0 clones, 408032/408033 dirty, 0/0 omap, 0/0 pinned, 0/0 hit_set_archive, 0/0 whiteouts, 46939015712/46939046528 bytes, 0/0 manifest objects, 0/0 hit_set_archive bytes.
```
同时健康检查会报`1 scrub errors (OSD_SCRUB_ERRORS)`和`Possible data damage: 1 pg inconsistent (PG_DAMAGED)`。

# q
如何查看一个处于inconsistent状态的PG（如10.3ff）的详细信息和所属OSD？
# a
- `ceph pg 10.3ff query`：查看PG详细状态，包括`last_scrub_stamp`、`last_clean_scrub_stamp`等时间戳，以及`up`/`acting` OSD列表。
- `ceph pg map 10.3ff`：输出当前OSD映射，如`osdmap e80324 pg 10.3ff (10.3ff) -> up [47,36,13,10,26] acting [47,36,13,10,26]`。
- `ceph pg dump | grep 10.3ff`：查看PG的概要状态、OSD集合、对象数、缺失对象等指标。

# q
解决Ceph PG inconsistent状态的标准修复流程是什么？
# a
1. 尝试停止相关OSD（`systemctl stop ceph-osd@<id>.service`）。
2. 刷写日志：`ceph-osd -i <id> --flush-journal`。
3. 启动OSD：`systemctl start ceph-osd@<id>.service`。
4. 若仍处于inconsistent，执行`ceph pg <pgid> repair`让集群自动修复。
5. 极端情况下，可用`ceph-objectstore-tool`导出PG数据（`--op export`），必要时移除（`--op remove`）并重新导入（`--op import`），但需谨慎。

# q
如何从pool对象列表中定位属于特定PG（如10.3ff）的对象？
# a
编写脚本遍历对象列表，使用`ceph osd map <pool> <object>`命令，检查输出中是否包含目标PGID。示例逻辑：
```bash
POOL=".product.rgw.buckets.ec"
PGID="10.3ff"
while read obj; do
  if ceph osd map "$POOL" "$obj" | grep -q "$PGID"; then
    echo "$obj" >> pg_objects.txt
  fi
done < object_list.txt
```
对象列表可通过`rados -p <pool> ls > object_list.txt`获取。

