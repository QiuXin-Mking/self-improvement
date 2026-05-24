# q
如何理解 `lctl dl` 命令输出中 Lustre 设备层次，如 osd、mgs、mds、ost、osc 等服务角色？
# a
- **osd-ldiskfs**：底层对象存储设备，负责对本地 ldiskfs 文件系统进行 IO 操作。
- **mgs / mgc**：管理服务（MGS）和客户端，负责管理配置信息并在节点间同步。
- **mds / mdt**：元数据服务（MDS）和元数据目标（MDT），管理文件系统的目录结构和元数据。
- **ost / obdfilter**：对象存储目标（OST）和对象过滤器，负责实际的文件数据存储。
- **osc / osp**：对象存储客户端（osc 用于客户端，osp 用于 MDT），用于与 OST 通信并聚合 IO。
- **lwp**：轻量级进程（Light Weight Process），各服务间用来互相监控或通信。
- **lov / lmv**：逻辑对象卷（LOV）和逻辑元数据卷（LMV），在客户端聚合 OST 和 MDT 的视图。

# q
在 Lustre 扩容中，使用 `mkfs.lustre` 创建新 MDT 和 OST 时分别需要哪些关键参数？
# a
- 创建 **MDT0001**：
  ```bash
  mkfs.lustre --mdt --mgsnode=172.31.0.26@tcp --fsname=nas_test --index=1 --reformat /dev/vdb
  ```
  其中 `--mdt` 指明目标类型为元数据目标，`--index` 为新 MDT 的索引号，`--reformat` 表示重新格式化设备。

- 创建 **OST0002** 和 **OST0003**：
  ```bash
  mkfs.lustre --ost --mgsnode=172.31.0.26@tcp --fsname=nas_test --index=2 --reformat /dev/vdc
  ```
  与 MDT 类似，`--ost` 指明为对象存储目标，`--index` 指定唯一索引号，`--mgsnode` 指定 MGS 地址。

# q
`lfs check all` 在 Lustre 扩容后有什么作用？
# a
`lfs check all` 用于检查所有服务（包括 MDT 和所有 OST）的可用状态。扩容后运行该命令可验证新增的 OST（如 `nas_test-OST0002-osc-MDT0000`）和 MDT（如 `nas_test-MDT0001`）是否已成功注册并显示为 `active`。如果所有目标都返回 `active`，说明客户端和管理节点均已感知新目标，扩容步骤基本完成。

