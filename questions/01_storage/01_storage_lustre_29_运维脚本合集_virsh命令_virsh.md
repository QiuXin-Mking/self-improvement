# q
如何为虚拟机创建快照？
# a
使用 `virsh snapshot-create-as` 命令，格式如下：
```bash
virsh snapshot-create-as <vm_name> <snapshot_name> --description "<description>"
```
- `<vm_name>`：虚拟机的名称或 ID。
- `<snapshot_name>`：快照的名称。
- `--description`：可选，对快照的简要描述。

示例：
```bash
virsh snapshot-create-as my_vm backup1 --description "Backup on October 26, 2023"
```

# q
如何删除虚拟机的指定快照？
# a
使用 `virsh snapshot-delete` 命令，格式如下：
```bash
virsh snapshot-delete <vm_name> --snapshotname <snapshot_name>
```
示例：
```bash
virsh snapshot-delete my_vm --snapshotname backup1
```

# q
如何将虚拟机恢复到某个快照状态？
# a
使用 `virsh snapshot-revert` 命令，格式如下：
```bash
virsh snapshot-revert <vm_name> <snapshot_name>
```
示例：
```bash
virsh snapshot-revert my_vm backup1
```

