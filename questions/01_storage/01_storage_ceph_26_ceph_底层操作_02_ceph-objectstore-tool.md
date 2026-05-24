# q
如何使用 ceph-objectstore-tool 列出指定 PG 中的所有对象？
# a
使用 `--op list` 并指定 `--pgid`，例如：
```bash
ceph-objectstore-tool --data-path /var/lib/ceph/osd/ceph-23 --pgid 4.120 --op list
```

# q
在 ceph-objectstore-tool 的 list 输出中，哪个 snapid 值代表对象的 head 版本？
# a
`snapid` 值为 `-2` 的条目代表对象的 head 版本。

# q
如何使用 ceph-objectstore-tool 导出某个对象的具体数据到文件？
# a
使用 `get-bytes` 操作，需要从 list 输出中复制完整的 JSON 对象描述作为参数，并指定输出文件路径。例如：
```bash
ceph-objectstore-tool --data-path /var/lib/ceph/osd/ceph-23 --pgid 4.120 \
  '["4.120",{"oid":"rbd_data.18570ff66275d5.000000000000a4da","key":"","snapid":3365,"hash":3096615712,"max":0,"pool":4,"namespace":"","max":0}]' \
  get-bytes rbd_data.18570ff66275d5.000000000000a4da_snapid_3365_bin
```

# q
ceph-objectstore-tool 如何删除一个底层对象？
# a
使用 `--op remove` 并通过 `--oid` 指定对象 ID（格式如 `rbd_data.7b70996d74511d.0000000000000bb5:16_head`）：
```bash
ceph-objectstore-tool --data-path /var/lib/ceph/osd/ceph-3 --op remove --oid 'rbd_data.7b70996d74511d.0000000000000bb5:16_head'
```

# q
ceph-objectstore-tool 支持哪些对象属性与 omap 操作？
# a
支持以下操作：
- 属性操作：`set-attr`、`get-attr`、`rm-attr`、`list-attrs`
- omap 操作：`set-omap`、`get-omap`、`rm-omap`、`list-omap`
- omap 头操作：`get-omaphdr`、`set-omaphdr`

