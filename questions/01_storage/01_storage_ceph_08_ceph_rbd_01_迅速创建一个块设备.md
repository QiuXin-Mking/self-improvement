# q
Ceph RBD是什么？
# a
Ceph RBD（RADOS Block Device）是 Ceph 集群提供的块存储服务，它允许用户像使用本地磁盘一样创建块设备，可在上面创建文件系统或用作虚拟机磁盘。RBD 块设备的数据条带化存储在 Ceph 集群的 OSD 中，具有高性能、高可靠和可扩展性。

# q
创建 RBD 镜像时启用 `--image-feature layering` 特性的作用是什么？
# a
它启用快照分层（layering）功能。该特性允许多个快照共享底层数据块，只保存与原始镜像的差异部分，从而节省存储空间，并支持快速回滚到镜像的某个状态。

# q
如何将 Ceph RBD 镜像映射为本地块设备？
# a
执行命令 ```sudo rbd map <pool>/<image> --id <client>```，成功后会在 `/dev/rbd/` 或 `/dev/` 下生成一个设备文件（如 `/dev/rbd0`），然后可像操作普通磁盘一样进行分区、格式化或挂载。

# q
对 RBD 块设备进行扩容的核心步骤是什么？
# a
首先使用 ```rbd resize --size <新大小MB> <pool>/<image>``` 调整镜像容量，然后对已挂载的文件系统执行在线扩容命令，例如：```resize2fs /dev/rbdX```（针对 ext 系列文件系统）或 ```xfs_growfs /mount/point```（针对 XFS），完成后即可用 `df -h` 看到扩容后的大小。

