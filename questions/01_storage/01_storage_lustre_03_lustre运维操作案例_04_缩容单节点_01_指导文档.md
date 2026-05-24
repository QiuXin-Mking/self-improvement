# q
在Lustre缩容过程中，如何将特定MDT（如mdt0）上的元数据迁移到其他MDT？
# a
使用 `lfs migrate -m` 命令，指定目标MDT索引列表，例如将所有元数据迁移到 mdt1、mdt2、mdt3：
```bash
lfs migrate -m 1,2,3 /mnt/lustre/*
```

# q
如何在所有节点上安全停用某个MDT服务（如MDT0000）？
# a
1. 查找MDT对应的 mdc active 文件路径：
   ```bash
   find /sys/fs/lustre/ -name "active" | grep MDT0000 | grep mdc
   ```
2. 在所有节点执行禁用命令：
   ```bash
   lctl set_param mdc.<fsname>-MDT0000-mdc-<instance>.active=0
   ```

# q
在对OST进行缩容时，如何防止新文件继续在目标OST上创建？
# a
将目标OST对应的所有 `osp` 实例的 `max_create_count` 设置为 0，该操作需在每个节点执行：
```bash
lctl set_param osp.<fsname>-OST0000-osc-MDT*.max_create_count=0
```
实例名称可通过 `lctl dl` 查询获得。

# q
如何将待缩容OST（如OST0000）上的所有现有数据迁移出去？
# a
使用 `lfs find` 筛选目标OST上的对象并通过 `lfs_migrate` 强制迁移，由于已设置 `max_create_count=0`，数据不会迁回原OST：
```bash
lfs find /mnt/lustre/ -obd <fsname>-OST0000 | lfs_migrate -y
```

# q
停用OST服务时，如何正确执行 `active=0` 操作？
# a
必须在**所有MDT节点**上，针对每个MDT与目标OST组合的 `osp` 实例执行禁用命令。例如停用OST0002和OST0003时：
- 在 mdt0000 节点执行：
  ```bash
  lctl set_param osp.<fsname>-OST0002-osc-MDT0000.active=0
  lctl set_param osp.<fsname>-OST0003-osc-MDT0000.active=0
  ```
- 在 mdt0001 节点执行：
  ```bash
  lctl set_param osp.<fsname>-OST0002-osc-MDT0001.active=0
  lctl set_param osp.<fsname>-OST0003-osc-MDT0001.active=0
  ```

