# q
如何找到目标 PG 所在的 OSD？
# a
使用命令 `ceph pg map <PG_ID>` 可以查看指定 PG 映射到哪些 OSD，包括 acting set 等信息。

# q
如何列出 PG 中包含的所有对象？
# a
使用命令 `ceph pg ls-by-pg <PG_ID> --format plain` 可以列出指定 PG 内的所有对象（以明文格式显示）。

# q
如何查找某个对象在 Ceph 集群中的位置信息？
# a
使用 `rados -p <pool_name> stat <object_name>` 可以查看指定存储池中某个对象的元数据和状态，包括其所在的 PG 等信息。

