# q
如何将Lustre OST从一个服务节点迁移到另一个服务节点的标准流程是什么？
# a
标准步骤：
1. 在旧节点卸载 OST：
```
umount /dev/rbd1
```
2. 取消 RBD 映射：
```
rbd unmap /dev/rbd1
```
3. 在新节点映射 RBD：
```
rbd map rbd/lustre_ost16
```
4. 修改文件系统参数，指定新服务节点和 MGS 地址：
```
tunefs.lustre --erase-params --servicenode=192.168.6.174@tcp --mgsnode=192.168.5.171@tcp /dev/rbd1
```
5. 挂载；
```
mount -t lustre /dev/rbd1 /data/lustre_ost16
```

# q
如何确认 RBD 设备在两节点间标识一致，确保 Lustre 迁移正确？
# a
在两个节点上分别运行 `blkid` 命令，检查对应设备的 LABEL 和 UUID 是否相同。例如，观察到的 RBD 设备：
- `/dev/rbd1`：LABEL="st_nas-OST0010"，UUID="5d10fec4-b4fc-4628-93ad-449fe6412194"
两个节点上此设备的 LABEL 和 UUID 必须完全一致。

# q
在 Lustre 迁移中，`tunefs.lustre --erase-params` 的作用是什么？
# a
`--erase-params` 用于清除目标设备上旧的 Lustre 参数（如原服务节点 NID），然后再通过 `--servicenode` 和 `--mgsnode` 重新设置，使文件系统能挂载到新的服务节点并正常注册到 MGS。

# q
进行 Lustre OST 迁移时，需要关注 ceph 集群的状态吗？
# a
迁移前应查看 `ceph -s` 确保集群健康，避免在迁移操作期间触发不必要的状态变更。案例中集群处于 HEALTH_WARN（scrub 滞后）但所有 OSD 均为 up/in，且设置了 `noout,nodeep-scrub` 标志，保障迁移时 OSD 不自动标记 out，减小风险。

