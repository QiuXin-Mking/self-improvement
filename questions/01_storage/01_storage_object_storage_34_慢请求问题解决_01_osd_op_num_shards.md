# q
如何正确查询已配置的 osd_op_num_shards 参数值？
# a
使用 `ceph config get` 命令，不能通过 `ceph config show` 获取结果。  
示例：  
```sh
ceph config get osd.1 osd_op_num_shards
```

# q
osd_op_num_shards 参数如何正确修改？
# a
在线修改均会失败（无法生效），必须使用离线配置文件方式。  
- 全局配置：在 `ceph.conf` 的 `[osd]` 段添加：
  ```
  osd_op_num_shards = 10
  ```
- 单 OSD 配置：在 `ceph.conf` 的 `[osd.1]` 段添加：
  ```
  osd_op_num_shards = 16
  ```
在线命令 `ceph config set osd.1 osd_op_num_shards 16` 和 `ceph tell osd.11 injectargs '--osd_op_num_shards=16'` 均无法修改成功。

