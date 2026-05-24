# q
`part_file_clear` 函数的功能是什么？
# a
该函数用于清除 `/dev` 目录下指定磁盘分区（`part_name`）的残留文件。它通过 `os.path.exists` 检查文件是否存在，若存在则调用 `os.remove` 删除，并返回操作是否成功的布尔值。若删除失败会记录错误日志。

# q
`part_osd` 方法在初始化磁盘时执行了哪些关键步骤？
# a
1. **容量对齐**：将 `md_cap` 从 GiB（1024 进制）转换为 MB（1000 进制），公式 `int((int(md_cap) << 20) / 1000 / 1000)`。
2. **创建 GPT 标签**：`parted -s /dev/<disk> mklabel gpt`
3. **清理旧元数据分区**：调用 `part_file_clear` 删除 `/dev/<disk>1`
4. **创建元数据分区**：`parted -s /dev/<disk> "mkpart primary 0 <md_cap_MB>"`
5. **清理旧数据分区**：调用 `part_file_clear` 删除 `/dev/<disk>2`
6. **创建数据分区**：`parted -s /dev/<disk> "mkpart primary <md_cap_MB> -1"`（从元数据分区末尾到磁盘末尾）
7. **清空元数据分区首部**：`dd if=/dev/zero of=/dev/<disk>1 bs=4096 count=1`

# q
在 `part_osd` 中，`dd if=/dev/zero of=/dev/%s1 bs=4096 count=1` 命令的作用是什么？
# a
用于清空磁盘第一个分区（元数据分区）的前 4096 字节。通过向 `/dev/<disk>1` 写入 4KB 的零数据，擦除可能残留的旧 OSD 元数据或文件系统签名，确保后续 OSD 初始化不会因残留数据而失败。

# q
所给代码中，如何通过 `etcd_delete` 删除节点心跳信息？
# a
调用 `etcd_delete(key="/cluster/nodes/heartbeat/%s" % nid)`，其中 `nid` 为节点 ID，构造出具体的 etcd 键路径，并执行删除操作。若删除成功，会记录日志 `"delete nodes key success, result:%s" % result`。

